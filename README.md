# Go Image Library Benchmarking

This repository contains a test harness and wrapper functions for benchmarking Go image library functionality.

The current list of libraries:
- [`containers/image`](https://github.com/containers/image), called "skopeo" here
 since [`skopeo`](https://github.com/containers/skopeo) is the most well-known entrypoint for this library.
- [`google/go-containerregistry`](https://github.com/google/go-containerregistry),
specifically the [`crane` package](https://github.com/google/go-containerregistry/tree/master/pkg/crane).

## Benchmarks

You can run benchmarks on your machine by running `make`.

Current stats:

```console
$ make V=1
CGO_ENABLED=0 go test -bench=. -benchmem -benchtime=10s -tags containers_image_openpgp -v
goos: linux
goarch: amd64
pkg: gitlab.com/estroz/go-image-lib-benchmark
cpu: 11th Gen Intel(R) Core(TM) i7-1195G7 @ 2.90GHz
BenchmarkCopySkopeo
BenchmarkCopySkopeo-8               1114          10202165 ns/op         1092486 B/op       6592 allocs/op
BenchmarkCopyCrane
BenchmarkCopyCrane-8                7525           1416258 ns/op          413058 B/op       2981 allocs/op
PASS
ok      gitlab.com/estroz/go-image-lib-benchmark        42.700s
```

## CLI

You can invoke supported library copy-like methods via CLI:

```sh
docker pull nginx:latest
docker pull registry:2
make cli
docker run -d -p 127.0.0.1:5000:5000 --restart always --name registry-crane registry:2
time ./bin/cli crane copy nginx:latest localhost:5000/nginx:latest
docker run -d -p 127.0.0.1:5001:5000 --restart always --name registry-skopeo registry:2
time ./bin/cli skopeo copy docker://nginx:latest docker://localhost:5001/nginx:latest
```
