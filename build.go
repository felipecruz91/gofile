package gofile

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/containerd/platforms"
	"github.com/felipecruz91/gofile/spec"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/exporter/containerimage/exptypes"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	keyFilename         = "filename"
	defaultGofileName   = "Gofile.yaml"
	localNameDockerfile = "dockerfile"
)

func Build(ctx context.Context, c client.Client) (*client.Result, error) {
	cfg, err := GetGofile(ctx, c)
	if err != nil {
		return nil, errors.Wrap(err, "got error getting Gofile")
	}

	state, img, err := Gofile2LLB(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert Gofile to LLB")
	}

	def, err := state.Marshal(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal local source")
	}

	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve Gofile")
	}

	ref, err := res.SingleRef()
	if err != nil {
		return nil, err
	}

	config, err := json.Marshal(img)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal image config")
	}
	k := platforms.Format(platforms.DefaultSpec())

	res.AddMeta(fmt.Sprintf("%s/%s", exptypes.ExporterImageConfigKey, k), config)

	res.SetRef(ref)

	return res, nil
}

func GetGofile(ctx context.Context, c client.Client) (*spec.Gofile, error) {
	opts := c.BuildOpts().Opts

	filename := opts[keyFilename]
	if filename == "" {
		filename = defaultGofileName
	}

	name := "load Gofile"
	if filename != defaultGofileName {
		name += " from " + filename
	}

	src := llb.Local(localNameDockerfile,
		llb.IncludePatterns([]string{filename}),
		llb.SessionID(c.BuildOpts().SessionID),
		llb.SharedKeyHint(defaultGofileName),
		llb.WithCustomName("[internal] "+name))

	def, err := src.Marshal(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal local source")
	}

	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve Gofile")
	}

	ref, err := res.SingleRef()
	if err != nil {
		return nil, err
	}

	b, err := ref.ReadFile(ctx, client.ReadRequest{
		Filename: filename,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to read Gofile")
	}

	var goFile spec.Gofile
	if err := yaml.Unmarshal(b, &goFile); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal Gofile")
	}

	return &goFile, nil
}
