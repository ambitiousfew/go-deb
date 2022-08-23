package cmd

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ambitiousfew/go-deb/debian"
	"github.com/spf13/cobra"
)

func buildPackage(wd string, output string) error {
	oCmd := exec.Command("fakeroot", "dpkg-deb", "--build", "debian", output)
	oCmd.Dir = wd
	oCmd.Stdout = os.Stdout
	oCmd.Stderr = os.Stderr
	return oCmd.Run()
}

func lintPackage(wd string, output string) error {
	oCmd := exec.Command("lintian", output)
	oCmd.Dir = wd
	oCmd.Stdout = os.Stdout
	oCmd.Stderr = os.Stderr
	return oCmd.Run()
}

func buildContents(cmd *cobra.Command, args []string) {
	// Access flags via: cmd.Flags().GetString("<flagname> (ex: work-dir)")
	output, _ := cmd.Flags().GetString("output")

	workdir, _ := cmd.Flags().GetString("work-dir")
	debfile, _ := cmd.Flags().GetString("deb-json")
	version, _ := cmd.Flags().GetString("version")
	arch, _ := cmd.Flags().GetString("arch")

	pkgDir := filepath.Join(workdir)
	if o, err := filepath.Abs(output); err != nil {
		log.Fatal(err)
	} else {
		output = o
	}

	debJSON := debian.Package{}

	// load the deb.json file
	if err := debJSON.Load(debfile); err != nil {
		log.Fatal(err)
	}
	log.Println("deb.json loaded")

	// normalize data
	debJSON.Normalize(pkgDir, version, arch)
	log.Println("pkg data normalized")

	log.Printf("Generating files in %s", pkgDir)
	if err := debJSON.GenerateFiles(filepath.Dir(debfile), pkgDir); err != nil {
		log.Fatal(err)
	}

	log.Printf("Building package in %s to %s", workdir, output)
	if err := buildPackage(pkgDir, output); err != nil {
		log.Fatal(err)
	}

	log.Printf("linting package in %s to %s", workdir, output)
	lintPackage(pkgDir, output) // it does not need to fail.
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the contents of the debian package",
	Long:  ``,
	Run:   buildContents,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	// Add flags for the generate command.
	generateCmd.Flags().StringP("work-dir", "w", "pkg-build", "Working directory to prepare the package. (default: pkg-build")
	generateCmd.Flags().StringP("output", "o", "", "Output directory for deb package file.")
	generateCmd.Flags().StringP("deb-json", "j", "deb.json", "Path to the deb.json file (default: deb.json)")
	generateCmd.Flags().StringP("version", "v", "", "Version of the package")
	generateCmd.Flags().StringP("arch", "a", "amd64", "Architecture of the package (ex: amd64)")
}
