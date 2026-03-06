#!/usr/bin/env python3
"""Generate BUILD files for pre-generated proto Go code."""

import os
import re
import hashlib
import glob

REPO_ROOT = os.path.dirname(os.path.abspath(__file__))

# Map Go import paths to Bazel labels
IMPORT_TO_LABEL = {
    # Internal proto packages -> our new pre-gen targets
    "github.com/buildbuddy-io/buildbuddy/proto/acl": "//proto/acl",
    "github.com/buildbuddy-io/buildbuddy/proto/action_cache": "//proto/action_cache",
    "github.com/buildbuddy-io/buildbuddy/proto/build_event_stream": "//proto/build_event_stream",
    "github.com/buildbuddy-io/buildbuddy/proto/cache": "//proto/cache",
    "github.com/buildbuddy-io/buildbuddy/proto/capability": "//proto/capability",
    "github.com/buildbuddy-io/buildbuddy/proto/command_line": "//proto/command_line",
    "github.com/buildbuddy-io/buildbuddy/proto/context": "//proto/context",
    "github.com/buildbuddy-io/buildbuddy/proto/failure_details": "//proto/failure_details",
    "github.com/buildbuddy-io/buildbuddy/proto/invocation": "//proto/invocation",
    "github.com/buildbuddy-io/buildbuddy/proto/invocation_policy": "//proto/invocation_policy",
    "github.com/buildbuddy-io/buildbuddy/proto/invocation_status": "//proto/invocation_status",
    "github.com/buildbuddy-io/buildbuddy/proto/option_filters": "//proto/option_filters",
    "github.com/buildbuddy-io/buildbuddy/proto/package_load_metrics": "//proto/package_load_metrics",
    "github.com/buildbuddy-io/buildbuddy/proto/remote_execution": "//proto/remote_execution",
    "github.com/buildbuddy-io/buildbuddy/proto/resource": "//proto/resource",
    "github.com/buildbuddy-io/buildbuddy/proto/scheduler": "//proto/scheduler",
    "github.com/buildbuddy-io/buildbuddy/proto/semver": "//proto/semver",
    "github.com/buildbuddy-io/buildbuddy/proto/stat_filter": "//proto/stat_filter",
    "github.com/buildbuddy-io/buildbuddy/proto/strategy_policy": "//proto/strategy_policy",
    "github.com/buildbuddy-io/buildbuddy/proto/target": "//proto/target",
    "github.com/buildbuddy-io/buildbuddy/proto/trace": "//proto/trace",
    "github.com/buildbuddy-io/buildbuddy/proto/user_id": "//proto/user_id",
    "github.com/buildbuddy-io/buildbuddy/proto/api/v1/common": "//proto/api/v1/common",
    # External deps
    "github.com/planetscale/vtprotobuf/protohelpers": "@com_github_planetscale_vtprotobuf//protohelpers",
    "github.com/planetscale/vtprotobuf/types/known/anypb": "@com_github_planetscale_vtprotobuf//types/known/anypb",
    "github.com/planetscale/vtprotobuf/types/known/durationpb": "@com_github_planetscale_vtprotobuf//types/known/durationpb",
    "github.com/planetscale/vtprotobuf/types/known/emptypb": "@com_github_planetscale_vtprotobuf//types/known/emptypb",
    "github.com/planetscale/vtprotobuf/types/known/fieldmaskpb": "@com_github_planetscale_vtprotobuf//types/known/fieldmaskpb",
    "github.com/planetscale/vtprotobuf/types/known/structpb": "@com_github_planetscale_vtprotobuf//types/known/structpb",
    "github.com/planetscale/vtprotobuf/types/known/timestamppb": "@com_github_planetscale_vtprotobuf//types/known/timestamppb",
    "github.com/planetscale/vtprotobuf/types/known/wrapperspb": "@com_github_planetscale_vtprotobuf//types/known/wrapperspb",
    "github.com/planetscale/vtprotobuf/vtproto": "@com_github_planetscale_vtprotobuf//include/github.com/planetscale/vtprotobuf/vtproto",
    "google.golang.org/grpc": "@org_golang_google_grpc//:grpc",
    "google.golang.org/grpc/codes": "@org_golang_google_grpc//codes",
    "google.golang.org/grpc/status": "@org_golang_google_grpc//status",
    "google.golang.org/protobuf/proto": "@org_golang_google_protobuf//proto",
    "google.golang.org/protobuf/reflect/protoreflect": "@org_golang_google_protobuf//reflect/protoreflect",
    "google.golang.org/protobuf/runtime/protoimpl": "@org_golang_google_protobuf//runtime/protoimpl",
    "google.golang.org/protobuf/types/known/anypb": "@org_golang_google_protobuf//types/known/anypb",
    "google.golang.org/protobuf/types/known/durationpb": "@org_golang_google_protobuf//types/known/durationpb",
    "google.golang.org/protobuf/types/known/emptypb": "@org_golang_google_protobuf//types/known/emptypb",
    "google.golang.org/protobuf/types/known/fieldmaskpb": "@org_golang_google_protobuf//types/known/fieldmaskpb",
    "google.golang.org/protobuf/types/known/structpb": "@org_golang_google_protobuf//types/known/structpb",
    "google.golang.org/protobuf/types/known/timestamppb": "@org_golang_google_protobuf//types/known/timestamppb",
    "google.golang.org/protobuf/types/known/wrapperspb": "@org_golang_google_protobuf//types/known/wrapperspb",
    "google.golang.org/protobuf/types/descriptorpb": "@org_golang_google_protobuf//types/descriptorpb",
    "google.golang.org/genproto/googleapis/rpc/status": "@org_golang_google_genproto_googleapis_rpc//status",
    "google.golang.org/genproto/googleapis/api/annotations": "@org_golang_google_genproto_googleapis_api//annotations",
    "cloud.google.com/go/longrunning/autogen/longrunningpb": "@com_google_cloud_go_longrunning//autogen/longrunningpb",
    "github.com/buildbuddy-io/buildbuddy/proto/options": "//proto/option_filters",
}

# Stdlib packages to ignore
STDLIB = {"fmt", "io", "math", "math/bits", "reflect", "sort", "sync", "unsafe",
          "encoding/binary", "strings", "bytes", "crypto/sha256", "os", "testing",
          "hash", "hash/crc32"}

def extract_imports(go_file):
    """Extract import paths from a Go file's import blocks only."""
    with open(go_file) as f:
        lines = f.readlines()
    imports = set()
    in_import = False
    for line in lines:
        stripped = line.strip()
        if stripped == 'import (' or stripped.startswith('import ('):
            in_import = True
            continue
        if in_import and stripped == ')':
            in_import = False
            continue
        if in_import:
            m = re.search(r'"([^"]+)"', stripped)
            if m:
                path = m.group(1)
                if path not in STDLIB and '.' in path.split('/')[0]:
                    imports.add(path)
        # Single-line import
        if stripped.startswith('import "'):
            m = re.search(r'"([^"]+)"', stripped)
            if m:
                path = m.group(1)
                if path not in STDLIB and '.' in path.split('/')[0]:
                    imports.add(path)
    return imports

def get_proto_source(pkg_dir):
    """Find the .proto source file for a package directory."""
    # proto/acl -> proto/acl.proto
    # proto/api/v1/common -> proto/api/v1/common.proto
    rel = os.path.relpath(pkg_dir, REPO_ROOT)
    name = os.path.basename(rel)
    proto_dir = os.path.dirname(rel)
    proto_file = os.path.join(proto_dir, f"{name}.proto")
    if os.path.exists(os.path.join(REPO_ROOT, proto_file)):
        return proto_file
    return None

def hash_file(path):
    with open(path, 'rb') as f:
        return hashlib.sha256(f.read()).hexdigest()

def generate_build(pkg_dir, importpath):
    """Generate BUILD file for a pre-generated proto package."""
    srcs = sorted(f for f in os.listdir(pkg_dir) if f.endswith('.pb.go'))
    
    # Collect all imports from all .pb.go files
    all_imports = set()
    for src in srcs:
        all_imports |= extract_imports(os.path.join(pkg_dir, src))
    
    # Map to Bazel labels
    deps = set()
    unmapped = set()
    for imp in all_imports:
        if imp in IMPORT_TO_LABEL:
            label = IMPORT_TO_LABEL[imp]
            # Don't depend on self
            rel = os.path.relpath(pkg_dir, REPO_ROOT)
            self_label = "//" + rel
            if label != self_label:
                deps.add(label)
        else:
            unmapped.add(imp)
    
    if unmapped:
        print(f"  WARNING: unmapped imports in {pkg_dir}: {unmapped}")
    
    deps = sorted(deps)
    name = os.path.basename(pkg_dir)
    
    # Determine proto source file path (relative from BUILD location)
    rel = os.path.relpath(pkg_dir, REPO_ROOT)
    proto_source = get_proto_source(pkg_dir)
    
    # Figure out the proto file label
    # proto/acl -> //proto:acl.proto
    # proto/api/v1/common -> //proto/api/v1:common.proto
    proto_dir = os.path.dirname(proto_source)
    proto_basename = os.path.basename(proto_source)
    proto_label = f"//{proto_dir}:{proto_basename}"
    
    # How deep is the package from repo root for rundir
    depth = len(rel.split('/'))
    
    has_grpc = any('_grpc.pb.go' in s for s in srcs)
    
    lines = []
    lines.append('load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")')
    lines.append('')
    lines.append('package(default_visibility = ["//visibility:public"])')
    lines.append('')
    lines.append('go_library(')
    lines.append(f'    name = "{name}",')
    lines.append(f'    srcs = [')
    for src in srcs:
        lines.append(f'        "{src}",')
    lines.append(f'    ],')
    lines.append(f'    importpath = "{importpath}",')
    if deps:
        lines.append(f'    deps = [')
        for dep in deps:
            lines.append(f'        "{dep}",')
        lines.append(f'    ],')
    lines.append(')')
    lines.append('')
    lines.append('go_test(')
    lines.append(f'    name = "{name}_drift_test",')
    lines.append(f'    srcs = ["drift_test.go"],')
    lines.append(f'    data = [')
    lines.append(f'        ".proto_hash",')
    lines.append(f'        "{proto_label}",')
    lines.append(f'    ],')
    lines.append(f'    rundir = ".",')
    lines.append(')')
    lines.append('')
    
    return '\n'.join(lines)

def generate_drift_test(pkg_dir, proto_source):
    """Generate drift_test.go for a package."""
    name = os.path.basename(pkg_dir)
    # Package name for Go - use the directory name with underscores
    # Actually, the package name in the .pb.go is the proto package name
    # Let's read it from the first .pb.go
    first_pb = sorted(f for f in os.listdir(pkg_dir) if f.endswith('.pb.go'))[0]
    with open(os.path.join(pkg_dir, first_pb)) as f:
        for line in f:
            m = re.match(r'package (\w+)', line)
            if m:
                go_pkg = m.group(1)
                break
    
    return f'''package {go_pkg}_test

import (
\t"crypto/sha256"
\t"fmt"
\t"os"
\t"strings"
\t"testing"
)

func TestProtoHashDrift(t *testing.T) {{
\tprotoBytes, err := os.ReadFile("{proto_source}")
\tif err != nil {{
\t\tt.Fatalf("reading proto source: %v", err)
\t}}
\twantHash := fmt.Sprintf("%x", sha256.Sum256(protoBytes))

\tgotHashBytes, err := os.ReadFile("{os.path.relpath(pkg_dir, REPO_ROOT)}/.proto_hash")
\tif err != nil {{
\t\tt.Fatalf("reading .proto_hash: %v", err)
\t}}
\tgotHash := strings.TrimSpace(string(gotHashBytes))

\tif gotHash != wantHash {{
\t\tt.Errorf("{proto_source} has changed since .pb.go files were generated\\n"+
\t\t\t"  stored hash: %s\\n"+
\t\t\t"  current hash: %s",
\t\t\tgotHash, wantHash)
\t}}
}}
'''

# All packages to process
PACKAGES = {
    "proto/acl": "github.com/buildbuddy-io/buildbuddy/proto/acl",
    "proto/action_cache": "github.com/buildbuddy-io/buildbuddy/proto/action_cache",
    "proto/build_event_stream": "github.com/buildbuddy-io/buildbuddy/proto/build_event_stream",
    "proto/cache": "github.com/buildbuddy-io/buildbuddy/proto/cache",
    "proto/capability": "github.com/buildbuddy-io/buildbuddy/proto/capability",
    "proto/command_line": "github.com/buildbuddy-io/buildbuddy/proto/command_line",
    "proto/context": "github.com/buildbuddy-io/buildbuddy/proto/context",
    "proto/failure_details": "github.com/buildbuddy-io/buildbuddy/proto/failure_details",
    "proto/invocation_policy": "github.com/buildbuddy-io/buildbuddy/proto/invocation_policy",
    "proto/invocation_status": "github.com/buildbuddy-io/buildbuddy/proto/invocation_status",
    "proto/option_filters": "github.com/buildbuddy-io/buildbuddy/proto/options",
    "proto/package_load_metrics": "github.com/buildbuddy-io/buildbuddy/proto/package_load_metrics",
    "proto/remote_execution": "github.com/buildbuddy-io/buildbuddy/proto/remote_execution",
    "proto/resource": "github.com/buildbuddy-io/buildbuddy/proto/resource",
    "proto/scheduler": "github.com/buildbuddy-io/buildbuddy/proto/scheduler",
    "proto/semver": "github.com/buildbuddy-io/buildbuddy/proto/semver",
    "proto/stat_filter": "github.com/buildbuddy-io/buildbuddy/proto/stat_filter",
    "proto/strategy_policy": "github.com/buildbuddy-io/buildbuddy/proto/strategy_policy",
    "proto/target": "github.com/buildbuddy-io/buildbuddy/proto/target",
    "proto/trace": "github.com/buildbuddy-io/buildbuddy/proto/trace",
    "proto/user_id": "github.com/buildbuddy-io/buildbuddy/proto/user_id",
    "proto/api/v1/common": "github.com/buildbuddy-io/buildbuddy/proto/api/v1/common",
}

def main():
    for pkg_rel, importpath in sorted(PACKAGES.items()):
        pkg_dir = os.path.join(REPO_ROOT, pkg_rel)
        if not os.path.exists(pkg_dir):
            print(f"SKIP (no dir): {pkg_rel}")
            continue
        
        proto_source = get_proto_source(pkg_dir)
        if not proto_source:
            print(f"SKIP (no .proto): {pkg_rel}")
            continue
        
        print(f"Processing {pkg_rel}...")
        
        # Skip invocation - already done manually
        if pkg_rel == "proto/invocation":
            continue
        
        # Generate BUILD
        build_content = generate_build(pkg_dir, importpath)
        with open(os.path.join(pkg_dir, 'BUILD'), 'w') as f:
            f.write(build_content)
        
        # Generate drift_test.go
        test_content = generate_drift_test(pkg_dir, proto_source)
        with open(os.path.join(pkg_dir, 'drift_test.go'), 'w') as f:
            f.write(test_content)
        
        # Generate .proto_hash
        proto_path = os.path.join(REPO_ROOT, proto_source)
        h = hash_file(proto_path)
        with open(os.path.join(pkg_dir, '.proto_hash'), 'w') as f:
            f.write(h + '\n')

if __name__ == '__main__':
    main()
