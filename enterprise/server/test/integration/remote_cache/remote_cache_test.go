package remote_cache_test

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/buildbuddy-io/buildbuddy/enterprise/server/testutil/buildbuddy_enterprise"
	"github.com/buildbuddy-io/buildbuddy/server/testutil/testbazel"
	"github.com/buildbuddy-io/buildbuddy/server/util/authutil"
	"github.com/buildbuddy-io/buildbuddy/server/util/retry"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/prototext"

	akpb "github.com/buildbuddy-io/buildbuddy/proto/api_key"
	bbspb "github.com/buildbuddy-io/buildbuddy/proto/buildbuddy_service"
	capb "github.com/buildbuddy-io/buildbuddy/proto/cache"
	cappb "github.com/buildbuddy-io/buildbuddy/proto/capability"
	inpb "github.com/buildbuddy-io/buildbuddy/proto/invocation"
	bspb "google.golang.org/genproto/googleapis/bytestream"
)

var (
	workspaceContents = map[string]string{
		"BUILD": `genrule(name = "hello_txt", outs = ["hello.txt"], cmd_bash = "echo 'Hello world' > $@")`,
	}
)

func TestBuild_RemoteCacheFlags_Anonymous_SecondBuildIsCached(t *testing.T) {
	app := buildbuddy_enterprise.Run(t)
	ctx := context.Background()
	ws := testbazel.MakeTempModule(t, workspaceContents)
	buildFlags := []string{"//:hello.txt"}
	buildFlags = append(buildFlags, app.BESBazelFlags()...)
	buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)

	result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	require.NotContains(
		t, result.Stderr, "1 remote cache hit",
		"sanity check: initial build shouldn't be cached",
	)

	// Clear the local cache so we can try for a remote cache hit.
	testbazel.Clean(ctx, t, ws)
	result = testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	assert.Contains(
		t, result.Stderr, "1 remote cache hit",
		"second build should be cached since anonymous users have cache write capabilities by default",
	)
}

func TestBuild_RemoteCacheFlags_ReadWriteApiKey_SecondBuildIsCached(t *testing.T) {
	ws := testbazel.MakeTempModule(t, workspaceContents)
	// Run the app with an API key we control so that we can authorize using it.
	app := buildbuddy_enterprise.Run(t)
	webClient := buildbuddy_enterprise.LoginAsDefaultSelfAuthUser(t, app)
	// Create a new read-write key
	rsp := &akpb.CreateApiKeyResponse{}
	err := webClient.RPC("CreateApiKey", &akpb.CreateApiKeyRequest{
		RequestContext: webClient.RequestContext,
		Capability:     []cappb.Capability{cappb.Capability_CACHE_WRITE},
	}, rsp)
	require.NoError(t, err)
	readWriteKey := rsp.ApiKey.Value
	buildFlags := []string{"//:hello.txt", fmt.Sprintf("--remote_header=%s=%s", authutil.APIKeyHeader, readWriteKey)}
	buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)

	ctx := context.Background()
	result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	assert.NotContains(
		t, result.Stderr, "1 remote cache hit",
		"sanity check: initial build shouldn't be cached",
	)

	// Clear the local cache so we can try for a remote cache hit.
	testbazel.Clean(ctx, t, ws)

	result = testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	assert.Contains(
		t, result.Stderr, "1 remote cache hit",
		"second build should be cached since the API key used in the first build has cache write capabilities",
	)
}

func TestBuild_RemoteCacheFlags_ReadOnlyApiKey_SecondBuildIsNotCached(t *testing.T) {
	ws := testbazel.MakeTempModule(t, workspaceContents)
	app := buildbuddy_enterprise.Run(t)
	webClient := buildbuddy_enterprise.LoginAsDefaultSelfAuthUser(t, app)
	rsp := &akpb.CreateApiKeyResponse{}
	// Create a new read-only key
	err := webClient.RPC("CreateApiKey", &akpb.CreateApiKeyRequest{
		RequestContext: webClient.RequestContext,
		Capability:     []cappb.Capability{},
	}, rsp)
	require.NoError(t, err)
	readOnlyKey := rsp.ApiKey.Value
	buildFlags := []string{"//:hello.txt", fmt.Sprintf("--remote_header=%s=%s", authutil.APIKeyHeader, readOnlyKey)}
	buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)

	ctx := context.Background()
	result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	assert.NotContains(
		t, result.Stderr, "1 remote cache hit",
		"sanity check: initial build shouldn't be cached",
	)

	// Clear the local cache so the remote cache will be queried.
	testbazel.Clean(ctx, t, ws)

	result = testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	assert.NotContains(
		t, result.Stderr, "1 remote cache hit",
		"second build should not be cached since the first build was done with a read-only key",
	)
}

func TestBuild_RemoteCacheFlags_CasOnlyApiKey_SecondBuildIsNotCached(t *testing.T) {
	ws := testbazel.MakeTempModule(t, workspaceContents)
	app := buildbuddy_enterprise.Run(t)
	webClient := buildbuddy_enterprise.LoginAsDefaultSelfAuthUser(t, app)
	rsp := &akpb.CreateApiKeyResponse{}
	// Create a new CAS-only key
	err := webClient.RPC("CreateApiKey", &akpb.CreateApiKeyRequest{
		RequestContext: webClient.RequestContext,
		Capability:     []cappb.Capability{cappb.Capability_CAS_WRITE},
	}, rsp)
	require.NoError(t, err)
	readOnlyKey := rsp.ApiKey.Value
	buildFlags := []string{"//:hello.txt", fmt.Sprintf("--remote_header=%s=%s", authutil.APIKeyHeader, readOnlyKey)}
	buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)

	ctx := context.Background()
	result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	assert.NotContains(
		t, result.Stderr, "1 remote cache hit",
		"sanity check: initial build shouldn't be cached",
	)

	// Clear the local cache so the remote cache will be queried.
	testbazel.Clean(ctx, t, ws)

	result = testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	assert.NotContains(
		t, result.Stderr, "1 remote cache hit",
		"second build should not be cached since the first build was done with a cas-only key",
	)
}

func TestBuild_RemoteCacheFlags_NoAuthConfigured_SecondBuildIsCached(t *testing.T) {
	app := buildbuddy_enterprise.RunWithConfig(t, buildbuddy_enterprise.DefaultAppConfig(t), buildbuddy_enterprise.NoAuthConfig)
	ctx := context.Background()
	ws := testbazel.MakeTempModule(t, workspaceContents)
	buildFlags := []string{"//:hello.txt"}
	buildFlags = append(buildFlags, app.BESBazelFlags()...)
	buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)

	result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	require.NotContains(
		t, result.Stderr, "1 remote cache hit",
		"sanity check: initial build shouldn't be cached",
	)

	// Clear the local cache so we can try for a remote cache hit.
	testbazel.Clean(ctx, t, ws)

	result = testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	assert.Contains(
		t, result.Stderr, "1 remote cache hit",
		"second build should be cached since anonymous users have cache write capabilities by default",
	)
}

func TestBuild_RemoteCacheFlags_Compression_SecondBuildIsCached(t *testing.T) {
	app := buildbuddy_enterprise.RunWithConfig(
		t, buildbuddy_enterprise.DefaultAppConfig(t), buildbuddy_enterprise.NoAuthConfig, "--cache.zstd_transcoding_enabled=true")
	ctx := context.Background()
	ws := testbazel.MakeTempModule(t, workspaceContents)
	buildFlags := []string{"//:hello.txt", "--experimental_remote_cache_compression"}
	buildFlags = append(buildFlags, app.BESBazelFlags()...)
	buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)

	result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	require.NotContains(
		t, result.Stderr, "1 remote cache hit",
		"sanity check: initial build shouldn't be cached",
	)

	// Clear the local cache so we can try for a remote cache hit.
	testbazel.Clean(ctx, t, ws)

	result = testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	assert.Contains(
		t, result.Stderr, "1 remote cache hit",
		"second build should be cached since anonymous users have cache write capabilities by default",
	)
}

func TestBuild_RemoteCache_RemoteBuildEventUploadAllUploadsLocalOutput(t *testing.T) {
	app := buildbuddy_enterprise.RunWithConfig(t, buildbuddy_enterprise.DefaultAppConfig(t), buildbuddy_enterprise.NoAuthConfig)
	ctx := context.Background()
	ws := testbazel.MakeTempModule(t, map[string]string{
		"BUILD": `genrule(
    name = "foo",
    outs = ["foo.txt"],
    cmd_bash = "echo 'foo bar' > $@",
    tags = ["no-remote"],
)`,
	})
	bepJSONPath := filepath.Join(ws, "bep.json")
	buildFlags := []string{
		"//:foo",
		"--remote_build_event_upload=all",
		"--build_event_json_file=" + bepJSONPath,
	}
	buildFlags = append(buildFlags, app.BESBazelFlags()...)
	buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)

	result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	require.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	bytestreamURI := findBEPFileURI(t, bepJSONPath, "foo.txt")
	assert.True(t, strings.HasPrefix(bytestreamURI, "bytestream://"), "BEP output URI should be a bytestream URI")
	blob := readBytestreamURI(t, ctx, app.ByteStreamClient(t), bytestreamURI)
	assert.Equal(t, "foo bar\n", string(blob))
}

func TestBuild_RemoteCache_RemoteDownloadMinimalBEPReferencesBytestream(t *testing.T) {
	app := buildbuddy_enterprise.RunWithConfig(t, buildbuddy_enterprise.DefaultAppConfig(t), buildbuddy_enterprise.NoAuthConfig)
	ctx := context.Background()
	ws := testbazel.MakeTempModule(t, map[string]string{
		"BUILD": `genrule(
    name = "foo",
    outs = ["foo.txt"],
    cmd_bash = "echo 'foo bar' > $@",
)`,
	})
	buildFlags := []string{"//:foo"}
	buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)

	result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)
	require.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")

	testbazel.Clean(ctx, t, ws)
	bepJSONPath := filepath.Join(ws, "bep.json")
	buildFlags = append(buildFlags, "--remote_download_minimal", "--build_event_json_file="+bepJSONPath)
	result = testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

	require.NoError(t, result.Error)
	assert.Contains(t, result.Stderr, "Build completed successfully")
	assert.Contains(t, result.Stderr, "1 remote cache hit")
	bepJSON := testfsReadFile(t, bepJSONPath)
	assert.NotContains(t, bepJSON, "file://")
	bytestreamURI := findBEPFileURI(t, bepJSONPath, "foo.txt")
	assert.True(t, strings.HasPrefix(bytestreamURI, "bytestream://"), "BEP output URI should be a bytestream URI")
	blob := readBytestreamURI(t, ctx, app.ByteStreamClient(t), bytestreamURI)
	assert.Equal(t, "foo bar\n", string(blob))
}

func TestBuild_RemoteCache_ScoreCard(t *testing.T) {
	app := buildbuddy_enterprise.RunWithConfig(
		t, buildbuddy_enterprise.DefaultAppConfig(t), buildbuddy_enterprise.NoAuthConfig,
		"--redis_command_buffer_flush_period=0",
		"--cache_stats_finalization_delay=0")
	bbService := app.BuildBuddyServiceClient(t)
	ctx := context.Background()
	ws := testbazel.MakeTempModule(t, workspaceContents)
	iid := newUUID(t)
	buildFlags := []string{"//:hello.txt", "--invocation_id=" + iid}
	buildFlags = append(buildFlags, app.BESBazelFlags()...)
	buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)

	{
		result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

		assert.NoError(t, result.Error)
		inv := getInvocationWithStats(t, ctx, bbService, iid)
		sc := inv.GetScoreCard()
		clearActionIDs(sc)
		expectedSC := &capb.ScoreCard{
			Misses: []*capb.ScoreCard_Result{
				{
					ActionMnemonic: "Genrule",
					TargetId:       "//:hello_txt",
				},
			},
		}
		expectedText, err := prototext.Marshal(expectedSC)
		require.NoError(t, err)
		scText, err := prototext.Marshal(sc)
		require.NoError(t, err)
		require.Equal(t, string(expectedText), string(scText))
	}

	// Clear the local cache so we can try for a remote cache hit.
	testbazel.Clean(ctx, t, ws)

	{
		iid = newUUID(t)
		buildFlags = append(buildFlags, "--invocation_id="+iid)

		result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)

		assert.NoError(t, result.Error)
		inv := getInvocationWithStats(t, ctx, bbService, iid)
		sc := inv.GetScoreCard()
		clearActionIDs(sc)
		expected := &capb.ScoreCard{
			Misses: []*capb.ScoreCard_Result{},
		}
		expectedText, err := prototext.Marshal(expected)
		require.NoError(t, err)
		scText, err := prototext.Marshal(sc)
		require.NoError(t, err)
		require.Equal(t, string(expectedText), string(scText))
	}
}

func TestBuild_RemoteCache_RejectsInvalidAPIKeys(t *testing.T) {
	app := buildbuddy_enterprise.Run(t, "--auth.enable_anonymous_usage=true")
	ctx := context.Background()
	ws := testbazel.MakeTempModule(t, workspaceContents)

	{
		// explicit empty API key; should fail
		testbazel.Clean(ctx, t, ws)
		iid := newUUID(t)
		buildFlags := []string{
			"//:hello.txt", "--invocation_id=" + iid,
			"--remote_header=x-buildbuddy-api-key=",
			"--remote_upload_local_results"}
		buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)
		result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)
		require.Contains(t, result.Stderr, "UNAUTHENTICATED")
		require.Contains(t, result.Stderr, "missing API key")
	}
	{
		// invalid API key; should fail
		testbazel.Clean(ctx, t, ws)
		iid := newUUID(t)
		buildFlags := []string{
			"//:hello.txt", "--invocation_id=" + iid,
			"--remote_header=x-buildbuddy-api-key=INVALID",
			"--remote_upload_local_results"}
		buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)
		result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)
		require.Contains(t, result.Stderr, "UNAUTHENTICATED")
		require.Contains(t, result.Stderr, "Invalid API key")
	}
	{
		// no explicit API key; should succeed
		testbazel.Clean(ctx, t, ws)
		iid := newUUID(t)
		buildFlags := []string{
			"//:hello.txt", "--invocation_id=" + iid,
			"--remote_upload_local_results"}
		buildFlags = append(buildFlags, app.RemoteCacheBazelFlags()...)
		result := testbazel.Invoke(ctx, t, ws, "build", buildFlags...)
		require.NotContains(t, result.Stderr, "UNAUTHENTICATED")
	}
}

func clearActionIDs(sc *capb.ScoreCard) {
	for _, m := range sc.GetMisses() {
		m.ActionId = ""
	}
}

func getInvocationWithStats(t *testing.T, ctx context.Context, bbService bbspb.BuildBuddyServiceClient, iid string) *inpb.Invocation {
	// Even though we've disabled redis buffering by setting
	// --redis_command_buffer_flush_period=0, we still need to poll for stats to
	// be written, since the stats recorder runs as a background job after the
	// invocation is completed. We use a short timeout though, since the job
	// should start immediately.
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	r := retry.DefaultWithContext(ctx)
	for r.Next() {
		res, err := bbService.GetInvocation(ctx, &inpb.GetInvocationRequest{
			Lookup: &inpb.InvocationLookup{InvocationId: iid},
		})
		require.NoError(t, err)
		require.Len(t, res.Invocation, 1, "expected exactly one invocation in GetInvocationResponse")
		inv := res.Invocation[0]
		cs := inv.GetCacheStats()
		if cs.GetCasCacheHits() > 0 || cs.GetCasCacheMisses() > 0 || cs.GetCasCacheUploads() > 0 {
			return inv
		}
	}
	require.FailNow(t, "Timed out waiting for invocation with cache stats")
	return nil
}

func newUUID(t *testing.T) string {
	id, err := uuid.NewRandom()
	if err != nil {
		t.Fatal(err)
	}
	return id.String()
}

func findBEPFileURI(t *testing.T, bepJSONPath, name string) string {
	t.Helper()
	bepJSON := testfsReadFile(t, bepJSONPath)
	re := regexp.MustCompile(`(?s)"name":"` + regexp.QuoteMeta(name) + `"[^{}]*"uri":"(bytestream://[^"]+)"`)
	matches := re.FindStringSubmatch(bepJSON)
	require.Len(t, matches, 2, "BEP JSON should contain a bytestream URI for %q:\n%s", name, bepJSON)
	return matches[1]
}

func testfsReadFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	require.NoError(t, err)
	return string(b)
}

func readBytestreamURI(t *testing.T, ctx context.Context, client bspb.ByteStreamClient, bytestreamURI string) []byte {
	t.Helper()
	u, err := url.Parse(bytestreamURI)
	require.NoError(t, err)
	resourceName := strings.TrimPrefix(u.Path, "/")
	require.NotEmpty(t, resourceName, "bytestream URI should include a resource name")
	stream, err := client.Read(ctx, &bspb.ReadRequest{ResourceName: resourceName})
	require.NoError(t, err)
	var out []byte
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return out
		}
		require.NoError(t, err)
		out = append(out, res.GetData()...)
	}
}
