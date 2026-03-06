package action_cache_test

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestProtoHashDrift(t *testing.T) {
	protoBytes, err := os.ReadFile("proto/action_cache.proto")
	if err != nil {
		t.Fatalf("reading proto source: %v", err)
	}
	wantHash := fmt.Sprintf("%x", sha256.Sum256(protoBytes))

	gotHashBytes, err := os.ReadFile("proto/action_cache/.proto_hash")
	if err != nil {
		t.Fatalf("reading .proto_hash: %v", err)
	}
	gotHash := strings.TrimSpace(string(gotHashBytes))

	if gotHash != wantHash {
		t.Errorf("proto/action_cache.proto has changed since .pb.go files were generated\n"+
			"  stored hash: %s\n"+
			"  current hash: %s",
			gotHash, wantHash)
	}
}
