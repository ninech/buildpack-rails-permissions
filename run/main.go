package main

import (
	"os"

	permissions "github.com/ninech/buildpack-rails-permissions"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func main() {
	logger := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))
	packit.Run(permissions.Detect(), permissions.Build(logger))
}
