package main

import (
	"context"

	"github.com/google/go-containerregistry/pkg/crane"
)

// CopyCrane copies srcRef to dstRef using go-containerregistry libs.
func CopyCrane(ctx context.Context, srcRef, dstRef string) error {
	opts := []crane.Option{
		crane.WithContext(ctx),
		crane.WithTransport(roundTripper),
		crane.Insecure,
	}
	return crane.Copy(srcRef, dstRef, opts...)
}
