package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the contents of the debian package",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Access flags via: cmd.Flags().GetString("<flagname> (ex: work-dir)")
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	// Add flags for the generate command.
	generateCmd.Flags().StringP("work-dir", "w", "", "Working directory to prepare the package.")
	generateCmd.Flags().StringP("output", "o", "", "Output directory for deb package file.")
	generateCmd.Flags().StringP("deb-json", "j", "", "Path to the deb.json file (default: deb.json)")
	generateCmd.Flags().StringP("version", "v", "", "Version of the package")
	generateCmd.Flags().StringP("arch", "a", "", "Architecture of the package (ex: amd64)")
}
