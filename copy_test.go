package main

import (
	"context"
	"testing"
)

func BenchmarkCopySkopeo(b *testing.B) {
	benchmarkCopy(b, CopySkopeo, "docker://")
}

func BenchmarkCopyCrane(b *testing.B) {
	benchmarkCopy(b, CopyCrane, "")
}

func benchmarkCopy(b *testing.B, f copyFunc, prefix string) {
	b.Helper()

	b.StopTimer()
	srcRegSrv, dstRegSrv, srcImgs, dstImgs := getRegistries(b, prefix)
	defer srcRegSrv.Close()
	defer dstRegSrv.Close()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		checkErrB(b, f(ctx, srcImgs[i], dstImgs[i]))
	}
}
