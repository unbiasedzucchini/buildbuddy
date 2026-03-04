// Uploads all built kernels (for all architectures) to GCS.
// See README.md in this dir for instructions.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/buildbuddy-io/buildbuddy/server/util/flag"
	"github.com/buildbuddy-io/buildbuddy/server/util/log"
)

var (
	destinationPrefix = flag.String("destination_prefix", "gs://buildbuddy-tools/binaries/linux", "GCS path prefix used when uploading built guest kernels")
)

// Set by x_defs in BUILD.
var (
	x86V5_15KernelRlocationpath   string
	x86V6_1KernelRlocationpath    string
	arm64V5_10KernelRlocationpath string
)

func main() {
	flag.Parse()
	if err := run(*destinationPrefix); err != nil {
		fmt.Fprintf(os.Stderr, "upload: %s\n", err)
		os.Exit(1)
	}
}

func run(destinationPrefix string) error {
	uploads := []struct {
		arch          string
		version       string
		rlocationpath string
	}{
		{arch: "x86_64", version: "v5.15", rlocationpath: x86V5_15KernelRlocationpath},
		{arch: "x86_64", version: "v6.1", rlocationpath: x86V6_1KernelRlocationpath},
		{arch: "aarch64", version: "v5.10", rlocationpath: arm64V5_10KernelRlocationpath},
	}
	type depsUpdate struct {
		repoName string
		sha256   string
		url      string
	}

	destPrefix := strings.TrimRight(destinationPrefix, "/")
	updates := make([]depsUpdate, 0, len(uploads))

	for _, upload := range uploads {
		if upload.rlocationpath == "" {
			return fmt.Errorf("%s kernel runfile path was not configured", upload.arch)
		}

		src, err := runfiles.Rlocation(upload.rlocationpath)
		if err != nil {
			return fmt.Errorf("rlocation %q: %w", upload.rlocationpath, err)
		}
		sha256Hex, err := fileSHA256(src)
		if err != nil {
			return fmt.Errorf("hash %q: %w", src, err)
		}
		name := fmt.Sprintf("vmlinux-%s-%s-%s", upload.arch, upload.version, sha256Hex)
		dest := destPrefix + "/" + name
		repoName, err := depsRepoName(upload.arch, upload.version)
		if err != nil {
			return err
		}

		exists, err := gcsObjectExists(dest)
		if err != nil {
			return fmt.Errorf("check whether %q already exists: %w", dest, err)
		}
		if exists {
			log.Infof("Skipping upload for %s guest kernel; object already exists at %s", upload.arch, dest)
		} else {
			log.Infof("Uploading %s guest kernel to %s", upload.arch, dest)
			cmd := exec.Command("gsutil", "cp", src, dest)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("upload %s guest kernel to %q: %w", upload.arch, dest, err)
			}
		}
		updates = append(updates, depsUpdate{
			repoName: repoName,
			sha256:   sha256Hex,
			url:      depsURL(dest),
		})
	}

	fmt.Println()
	fmt.Println("deps.bzl update is required to use new kernels:")
	fmt.Println()
	for _, update := range updates {
		fmt.Println("    http_file(")
		fmt.Printf("        name = %q,\n", update.repoName)
		fmt.Printf("        sha256 = %q,\n", update.sha256)
		fmt.Printf("        urls = [%q],\n", update.url)
		fmt.Println("        executable = True,")
		fmt.Println("    )")
	}

	return nil
}

func depsRepoName(arch, version string) (string, error) {
	switch {
	case arch == "x86_64" && version == "v5.15":
		return "org_kernel_git_linux_kernel-vmlinux", nil
	case arch == "x86_64" && version == "v6.1":
		return "org_kernel_git_linux_kernel-vmlinux-6.1", nil
	case arch == "aarch64" && version == "v5.10":
		return "org_kernel_git_linux_kernel-vmlinux-arm64", nil
	default:
		return "", fmt.Errorf("unsupported kernel target arch=%q version=%q", arch, version)
	}
}

func depsURL(dest string) string {
	if strings.HasPrefix(dest, "gs://") {
		return "https://storage.googleapis.com/" + strings.TrimPrefix(dest, "gs://")
	}
	return dest
}

func gcsObjectExists(dest string) (bool, error) {
	cmd := exec.Command("gsutil", "ls", dest)
	output, err := cmd.CombinedOutput()
	if err == nil {
		return true, nil
	}
	msg := string(output)
	if strings.Contains(msg, "matched no objects") ||
		strings.Contains(msg, "No URLs matched") ||
		strings.Contains(msg, "NotFoundException") {
		return false, nil
	}
	return false, fmt.Errorf("gsutil ls: %w: %s", err, strings.TrimSpace(msg))
}

func fileSHA256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
