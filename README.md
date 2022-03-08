# Go Image Library Benchmarking

This repository contains a test harness and wrapper functions for benchmarking Go image library functionality.

The current list of libraries:
- [`containers/image`](https://github.com/containers/image), called "skopeo" here
 since [`skopeo`](https://github.com/containers/skopeo) is the most well-known entrypoint for this library.
- [`google/go-containerregistry`](https://github.com/google/go-containerregistry),
specifically the [`crane` package](https://github.com/google/go-containerregistry/tree/main/pkg/crane).

## Benchmarks

You can run benchmarks on your machine by running `make`.

Hardware info:

```
goos: linux
goarch: amd64
pkg: gitlab.com/estroz/go-image-lib-benchmark
cpu: 11th Gen Intel(R) Core(TM) i7-1195G7 @ 2.90GHz
```

Current stats:

```
benchmark               iter     time/iter    bytes alloc           allocs
---------               ----     ---------    -----------           ------
BenchmarkCopySkopeo-8   1000   10.05 ms/op   1119918 B/op   6660 allocs/op
BenchmarkCopyCrane-8    1000    1.33 ms/op    417651 B/op   2985 allocs/op
```

## CLI

You can invoke supported library copy-like methods via CLI:

```sh
docker pull nginx:latest
docker pull registry:2
docker run -d -p 127.0.0.1:5000:5000 --restart always --name registry-crane registry:2
make cli
./bin/cli crane copy nginx:latest localhost:5000/nginx:latest
```
