/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/leszkolukasz/crunchyroll-notifier/models"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Removes all accounts",
	Long:  "No longer receive notification for any account",

	Run: func(cmd *cobra.Command, args []string) {
		models.InjectUserList([]models.User{})
		fmt.Println("Clearing finished.")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
