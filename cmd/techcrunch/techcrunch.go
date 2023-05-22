package techcrunch

import (
	"github.com/spf13/cobra"
)

// techCrunchCmd represents the techCrunch command
var TechcrunchCmd = &cobra.Command{
	Use:   "tc",
	Short: "Get latest blog posts from TechCrunch",
	Run: func(cmd *cobra.Command, args []string) {
		techcrunch()
	},
}
