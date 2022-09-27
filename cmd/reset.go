package cmd

import (
	"github.com/leszkolukasz/crunchyroll-notifier/config"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Generates empty config file",
	Long:  `Application requires config file to exist. `,
	Run: func(cmd *cobra.Command, args []string) {
		config.GenerateConfiguration()
		panic("OK") // Hack
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
