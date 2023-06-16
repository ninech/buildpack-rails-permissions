package permissions

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		layersDir  string
		workingDir string
		cnbDir     string

		build packit.BuildFunc
	)

	it.Before(func() {
		var err error
		layersDir, err = os.MkdirTemp("", "layers")
		Expect(err).NotTo(HaveOccurred())

		cnbDir, err = os.MkdirTemp("", "cnb")
		Expect(err).NotTo(HaveOccurred())

		workingDir, err = os.MkdirTemp("", "working-dir")
		Expect(err).NotTo(HaveOccurred())

		build = Build(scribe.NewEmitter(bytes.NewBuffer(nil)))
	})

	it.After(func() {
		Expect(os.RemoveAll(layersDir)).To(Succeed())
		Expect(os.RemoveAll(cnbDir)).To(Succeed())
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	it("creates rails dirs and sets permissions", func() {
		_, err := build(packit.BuildContext{
			WorkingDir: workingDir,
			CNBPath:    cnbDir,
			Stack:      "some-stack",
			BuildpackInfo: packit.BuildpackInfo{
				Name:    "Some Buildpack",
				Version: "some-version",
			},
			Plan: packit.BuildpackPlan{
				Entries: []packit.BuildpackPlanEntry{},
			},
			Layers: packit.Layers{Path: layersDir},
		})
		Expect(err).NotTo(HaveOccurred())

		tmp, err := os.Stat(filepath.Join(workingDir, tmpDirName))
		Expect(err).NotTo(HaveOccurred())
		Expect(tmp.Mode()).To(Equal(os.ModeDir | 0770))

		pids, err := os.Stat(filepath.Join(workingDir, tmpDirName, pidsDirName))
		Expect(err).NotTo(HaveOccurred())
		Expect(pids.Mode()).To(Equal(os.ModeDir | 0770))

		cache, err := os.Stat(filepath.Join(workingDir, tmpDirName, cacheDirName))
		Expect(err).NotTo(HaveOccurred())
		Expect(cache.Mode()).To(Equal(os.ModeDir | 0770))
	})
}
