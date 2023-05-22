package techcrunch

import (
	"github.com/spf13/cobra"
)

var (
	category string
)

// techCrunchCmd represents the techCrunch command
var TechcrunchCmd = &cobra.Command{
	Use:   "tc [flags]",
	Short: "Get latest blog posts from TechCrunch",
	Run: func(cmd *cobra.Command, args []string) {
		techcrunch(category)
	},
}

func init() {
	// Flag to specify category of blog posts to search for
	TechcrunchCmd.Flags().StringVarP(&category, "category", "c", "", "Category of blog posts to search for")
}
