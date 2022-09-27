/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/leszkolukasz/crunchyroll-notifier/models"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes account",
	Long:  "No longer receive notification for this account",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
		userList := models.ExtractUserList()
		userList = models.RemoveUser(userList, "leszko.lucas@gmail.com")
		models.InjectUserList(userList)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
