package permissions

import (
	"os"
	"path/filepath"
	"testing"

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

	context("when a rails gemfile is present", func() {
		it.Before(func() {
			var err error
			workingDir, err = os.MkdirTemp("", "working-dir-*")
			Expect(err).NotTo(HaveOccurred())

			const gemfile = `
source "https://rubygems.org"

gem "rails", "~> 7.0.0"
		`
			err = os.WriteFile(filepath.Join(workingDir, "Gemfile"), []byte(gemfile), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())

			detect = Detect()
		})

		it.After(func() {
			Expect(os.RemoveAll(workingDir)).To(Succeed())
		})

		it("detects", func() {
			_, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
		})
	})

	context("when a gemfile is present without rails", func() {
		it.Before(func() {
			var err error
			workingDir, err = os.MkdirTemp("", "working-dir-*")
			Expect(err).NotTo(HaveOccurred())

			const gemfile = `
source "https://rubygems.org"

gem "something-else", "~> 1.0.0"
		`
			err = os.WriteFile(filepath.Join(workingDir, "Gemfile"), []byte(gemfile), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())

			detect = Detect()
		})

		it.After(func() {
			Expect(os.RemoveAll(workingDir)).To(Succeed())
		})

		it("fails detection", func() {
			_, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).To(HaveOccurred())
		})
	})

	context("when no gemfile is present", func() {
		it.Before(func() {
			detect = Detect()
		})

		it("fails detection", func() {
			_, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).To(HaveOccurred())
		})
	})
}
