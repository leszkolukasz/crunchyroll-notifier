/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/leszkolukasz/crunchyroll-notifier/models"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds new account",
	Long:  "Adds new Crunchyroll account to receive notifications for",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		userList := models.ExtractUserList()
		userList = models.AddUser(userList, "leszko.lucas@gmail.com", "^&52rbr9W$SLgBk")
		models.InjectUserList(userList)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
