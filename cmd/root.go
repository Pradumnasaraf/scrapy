package cmd

import (
	"os"

	"github.com/Pradumnasaraf/scrapy/cmd/ebay"
	"github.com/Pradumnasaraf/scrapy/cmd/techcrunch"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scrapy [command] [flags]",
	Short: "Scape the web from the command line",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	// Removes the default completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(ebay.EbayCmd)
	rootCmd.AddCommand(techcrunch.TechcrunchCmd)
}
