package gofile

import (
	"runtime"

	"github.com/moby/buildkit/util/system"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

func NewImageConfig() *specs.Image {
	return &specs.Image{
		Platform: specs.Platform{
			OS:           "linux",
			Architecture: runtime.GOARCH,
		},
		RootFS: specs.RootFS{
			Type: "layers",
		},
		Config: specs.ImageConfig{
			WorkingDir: "/",
			Env:        []string{"PATH=" + system.DefaultPathEnvUnix},
			Cmd:        []string{"/bin/server"},
		},
	}
}
