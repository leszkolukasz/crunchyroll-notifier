package cmd

import (
	"github.com/leszkolukasz/crunchyroll-notifier/config"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets application to default settings",
	Long:  "Resets application to default settings. Use this if application does not work. WARNING: This will remove all added users and configs",

	Run: func(cmd *cobra.Command, args []string) {
		config.GenerateConfiguration()
		panic("OK") // Hack
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
