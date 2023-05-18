package ebay

import (
	"github.com/spf13/cobra"
)

// ebayCmd represents the ebay command
var EbayCmd = &cobra.Command{
	Use:   "ebay [flags]",
	Short: "Get Product and Price from Ebay",
	Run: func(cmd *cobra.Command, args []string) {
		ebay()
	},
}
