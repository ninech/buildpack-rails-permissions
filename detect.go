package permissions

import (
	"fmt"
	"path/filepath"

	"github.com/paketo-buildpacks/packit/v2"
	railsassets "github.com/paketo-buildpacks/rails-assets"
)

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		hasRails, err := railsassets.NewGemfileParser().Parse(filepath.Join(context.WorkingDir, "Gemfile"))
		if err != nil {
			return packit.DetectResult{}, fmt.Errorf("failed to parse Gemfile: %w", err)
		}

		if !hasRails {
			return packit.DetectResult{}, packit.Fail.WithMessage("failed to find rails gem in Gemfile")
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{},
				Requires: []packit.BuildPlanRequirement{},
			},
		}, nil
	}
}
