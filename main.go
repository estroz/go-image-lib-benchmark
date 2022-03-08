package main

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	roundTripper *http.Transport
)

type copyFunc func(context.Context, string, string) error

func init() {
	roundTripper = http.DefaultTransport.(*http.Transport).Clone()
	roundTripper.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatalf("expected at least 2 args, got: %d", len(args))
	}

	libName, methodName := args[0], args[1]

	ctx := context.Background()

	switch libName {
	case "skopeo":
		runSkopeo(ctx, methodName, args[2:]...)
	case "crane":
		runCrane(ctx, methodName, args[2:]...)
	default:
		log.Fatalf("unknown lib name %q, must be one of: skopeo, crane", libName)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runCrane(ctx context.Context, methodName string, args ...string) {
	switch methodName {
	case "copy":
		srcRef, dstRef := args[0], args[1]
		checkErr(CopyCrane(ctx, srcRef, dstRef))
	default:
		log.Fatalf("crane: unknown method name %q, must be one of: copy", methodName)
	}
}

func runSkopeo(ctx context.Context, methodName string, args ...string) {
	switch methodName {
	case "copy":
		srcRef, dstRef := args[0], args[1]
		checkErr(CopySkopeo(ctx, srcRef, dstRef))
	default:
		log.Fatalf("skopeo: unknown method name %q, must be one of: copy", methodName)
	}
}
