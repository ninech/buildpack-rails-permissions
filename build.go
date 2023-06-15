package permissions

import (
	"os"
	"path/filepath"
	"time"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

const (
	tmpDirName   = "tmp"
	pidsDirName  = "pids"
	cacheDirName = "cache"
)

func Build(logger scribe.Emitter) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)
		before := time.Now()

		dirs := []string{
			filepath.Join(context.WorkingDir, tmpDirName),
			filepath.Join(context.WorkingDir, tmpDirName, pidsDirName),
			filepath.Join(context.WorkingDir, tmpDirName, cacheDirName),
		}

		logger.Process("changing permission for directories:")
		for _, dir := range dirs {
			logger.Process(dir)

			if err := os.MkdirAll(dir, 0770); err != nil {
				return packit.BuildResult{}, err
			}

			if err := os.Chmod(dir, 0770); err != nil {
				return packit.BuildResult{}, err
			}
		}

		logger.Action("Completed in %s", time.Since(before))

		return packit.BuildResult{}, nil
	}
}
