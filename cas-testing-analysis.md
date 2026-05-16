# CAS Exhaustive Testing Plan

Starting with the ContentAddressableStorage service is a good pilot. CAS is small enough to cover deeply, but it still exercises the important BuildBuddy gRPC concerns: auth and tenant prefixing, digest validation, cache storage, compression, streaming `GetTree`, chunked `SplitBlob` / `SpliceBlob`, and ByteStream interoperability.

## Recommended Approach

1. Start with typed CAS tests, not reflection.

   The existing CAS test harness spins up an in-process CAS + ByteStream server in `server/remote_cache/content_addressable_storage_server/content_addressable_storage_server_test.go`. Typed tests give precise assertions and are easier to debug than a reflection-driven sweep. Reflection is still useful later as a broad API smoke test.

2. Cover the core CAS state machine.

   Exercise `FindMissingBlobs -> BatchUpdateBlobs -> FindMissingBlobs -> BatchReadBlobs` across empty digests, valid missing digests, valid uploaded digests, duplicates, mixed good and bad batch entries, invalid hashes and sizes, digest functions, authenticated tenant prefixing, instance-name behavior, and read-only credentials.

3. Treat compression as a first-class matrix.

   `BatchUpdateBlobs` and `BatchReadBlobs` branch on zstd support, cache compressor support, accepted client compressors, and digest validation after decompression. Tests should cover identity-to-identity, identity-to-zstd, zstd-to-identity, disabled zstd, corrupt zstd bytes, and cache implementations with and without native compressed storage.

4. Test `GetTree` as graph traversal.

   Existing tests cover normal directory trees. Additional coverage should include malformed directory blobs, missing child directories, duplicate child names, empty root handling, bad page tokens, page-size boundaries, and subtree-cache behavior.

5. Include BuildBuddy CAS extensions.

   `SplitBlob` and `SpliceBlob` are likely places for boundary bugs. Coverage should include reordered chunks, duplicate chunks, omitted chunks, wrong digest function, unsupported chunking function, single chunk, zero chunks, missing chunks, and chunk manifests whose declared blob digest does not match reconstructed bytes.

## What To Defer

A reflection-driven negative sweep is worth doing later, but it should not be the first CAS test layer. It gives weak assertions such as "does not panic and returns a gRPC status." That is useful for the whole API surface, especially when new RPCs are added, but CAS needs semantic assertions first.

I would also avoid starting with a new ByteStream resource-name fuzzer. This repo already has fuzz targets for resource-name parsing in `server/remote_cache/digest/digest_test.go`. The next higher-value fuzz/property work is around chunk boundaries, manifest validation, and generated CAS operation sequences.

## First Implementation Slice

The first deterministic slice should live in `server/remote_cache/content_addressable_storage_server/content_addressable_storage_server_test.go` and add:

- invalid digest rejection for `BatchUpdateBlobs` and `BatchReadBlobs`
- authenticated tenant prefix isolation for `FindMissingBlobs`, `BatchUpdateBlobs`, and `BatchReadBlobs`
- `SpliceBlob` validation for reordered chunks
- corrupt zstd upload handling
- `GetTree` handling for missing root digests, malformed root directory blobs, and missing child directories
- duplicate and mixed-missing `BatchReadBlobs` response semantics
- unsupported chunking-function rejection for `SplitBlob` and `SpliceBlob`
- ByteStream write/read interoperability with CAS read/write APIs, including authenticated tenant prefixing

Use Bazel for verification:

```bash
bazel test --remote_download_minimal //server/remote_cache/content_addressable_storage_server:content_addressable_storage_server_test
```
