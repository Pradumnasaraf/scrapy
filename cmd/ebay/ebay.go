package ebay

import (
	"github.com/spf13/cobra"
)

var (
	pages int = 1
)

// ebayCmd represents the ebay command
var EbayCmd = &cobra.Command{
	Use:   "ebay [flags]",
	Short: "Get Product and Price from Ebay",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		ebay(args, pages)

	},
}

func init() {
	EbayCmd.Flags().IntVarP(&pages, "pages", "p", 1, "Number of pages to scrape")
}
