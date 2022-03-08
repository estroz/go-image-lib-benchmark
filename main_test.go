package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/registry"
	"github.com/google/go-containerregistry/pkg/v1/random"
)

func getRegistries(b *testing.B, prefix string) (srcRegSrv, dstRegSrv *httptest.Server, srcImgs, dstImgs []string) {
	srcImgs = make([]string, b.N)
	dstImgs = make([]string, b.N)

	logOut := bytes.Buffer{}

	srcRegHandler := registry.New(newLoggerOpt(&logOut, "[src]"))
	srcRegSrv = httptest.NewServer(srcRegHandler)
	srcAddr := srcRegSrv.Listener.Addr()
	dstRegHandler := registry.New(newLoggerOpt(&logOut, "[dst]"))
	dstRegSrv = httptest.NewServer(dstRegHandler)
	dstAddr := dstRegSrv.Listener.Addr()

	pushOpts := []crane.Option{
		crane.WithTransport(roundTripper),
		crane.Insecure,
	}

	for i := 0; i < b.N; i++ {
		srcImgs[i] = fmt.Sprintf("%s/sai-bench/image:%d", srcAddr, i)
		dstImgs[i] = fmt.Sprintf("%s/sai-bench/image:%d", dstAddr, i)

		rimg, err := random.Image(100, 2)
		checkErrB(b, err)

		if err = crane.Push(rimg, srcImgs[i], pushOpts...); err != nil {
			fmt.Fprintln(os.Stderr, lastNLines(&logOut, 10))
			b.Fatal(err)
		}

		// Set the prefix after doing the request since ref prefixes aren't valid transport protos.
		if prefix != "" {
			srcImgs[i] = fmt.Sprintf("%s%s", prefix, srcImgs[i])
			dstImgs[i] = fmt.Sprintf("%s%s", prefix, dstImgs[i])
		}
	}

	return srcRegSrv, dstRegSrv, srcImgs, dstImgs
}

func newLoggerOpt(logOut io.Writer, prefix string) registry.Option {
	return registry.Logger(log.New(logOut, prefix, log.LstdFlags))
}

func lastNLines(buf *bytes.Buffer, n int) string {
	split := strings.Split(buf.String(), "\n")
	if len(split) < n {
		n = len(split)
	}
	return strings.Join(split[len(split)-n:], "\n")
}

func checkErrB(b *testing.B, err error) {
	b.Helper()
	if err != nil {
		b.Fatal(err)
	}
}
