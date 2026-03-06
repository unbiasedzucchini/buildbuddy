package invocation_test

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"testing"
)

// TestProtoHashDrift ensures the pre-generated .pb.go files stay in sync with
// the source .proto file. If this test fails, regenerate the .pb.go files:
//
//	bazel build //proto:invocation_go_proto
//	cp bazel-bin/proto/invocation_go_proto_/.../invocation/*.pb.go proto/invocation/
//	sha256sum proto/invocation.proto | awk '{print $1}' > proto/invocation/.proto_hash
func TestProtoHashDrift(t *testing.T) {
	protoBytes, err := os.ReadFile("proto/invocation.proto")
	if err != nil {
		t.Fatalf("reading proto source: %v", err)
	}
	wantHash := fmt.Sprintf("%x", sha256.Sum256(protoBytes))

	gotHashBytes, err := os.ReadFile("proto/invocation/.proto_hash")
	if err != nil {
		t.Fatalf("reading .proto_hash: %v", err)
	}
	gotHash := strings.TrimSpace(string(gotHashBytes))

	if gotHash != wantHash {
		t.Errorf("proto/invocation.proto has changed since .pb.go files were generated\n"+
			"  stored hash: %s\n"+
			"  current hash: %s\n"+
			"Regenerate with:\n"+
			"  bazel build //proto:invocation_go_proto\n"+
			"  cp bazel-bin/proto/invocation_go_proto_/.../invocation/*.pb.go proto/invocation/\n"+
			"  sha256sum proto/invocation.proto | awk '{print $1}' > proto/invocation/.proto_hash",
			gotHash, wantHash)
	}
}
