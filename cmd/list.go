/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/leszkolukasz/crunchyroll-notifier/models"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all accounts",
	Long:  "Lists all Crunchyroll account to receive notifications for",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		userList := models.ExtractUserList()

		if len(userList) == 0 {
			fmt.Println("No users.")
			return
		}

		fmt.Println("Users:")
		for idx, user := range userList {
			fmt.Println(idx+1, user.Email, user.Password)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
