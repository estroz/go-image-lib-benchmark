package main

import (
	"context"
	"io/ioutil"

	"github.com/containers/image/v5/copy"
	"github.com/containers/image/v5/signature"
	"github.com/containers/image/v5/transports/alltransports"
	"github.com/containers/image/v5/types"
)

// CopySkopeo copies srcRef to dstRef using containers/image libs.
func CopySkopeo(ctx context.Context, srcRef, dstRef string) error {
	srcIRef, err := alltransports.ParseImageName(srcRef)
	if err != nil {
		return err
	}
	dstIRef, err := alltransports.ParseImageName(dstRef)
	if err != nil {
		return err
	}

	sc := types.SystemContext{
		OCIInsecureSkipTLSVerify:          true,
		DockerDaemonInsecureSkipTLSVerify: true,
		DockerInsecureSkipTLSVerify:       types.NewOptionalBool(true),
		DockerAuthConfig:                  &types.DockerAuthConfig{},
	}
	srcCtx := new(types.SystemContext)
	*srcCtx = sc
	dstCtx := new(types.SystemContext)
	*dstCtx = sc

	policy := &signature.Policy{
		Default: []signature.PolicyRequirement{signature.NewPRInsecureAcceptAnything()},
	}
	policyCtx, err := signature.NewPolicyContext(policy)
	if err != nil {
		return err
	}

	opts := &copy.Options{
		SourceCtx:          srcCtx,
		DestinationCtx:     dstCtx,
		ImageListSelection: copy.CopyAllImages,
		ReportWriter:       ioutil.Discard,
		// PreserveDigests:    true,
	}
	_, err = copy.Image(ctx, policyCtx, dstIRef, srcIRef, opts)
	return err
}
