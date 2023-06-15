package permissions_test

import (
	"os"
	"path/filepath"
	"testing"

	permissions "github.com/ninech/buildpack-rails-permissions"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		workingDir string
		detect     packit.DetectFunc
	)

	it.Before(func() {
		var err error
		workingDir, err = os.MkdirTemp("", "working-dir")
		Expect(err).NotTo(HaveOccurred())

		const gemfile = `
source "https://rubygems.org"

gem "rails", "~> 7.0.0"
		`
		err = os.WriteFile(filepath.Join(workingDir, "Gemfile"), []byte(gemfile), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())

		detect = permissions.Detect()
	})

	it.After(func() {
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	context("when conditions for detect true are met", func() {
		it("detects", func() {
			_, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
		})
	})
}
