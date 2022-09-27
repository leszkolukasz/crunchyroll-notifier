package cmd

import (
	"fmt"
	"github.com/leszkolukasz/crunchyroll-notifier/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func Execute() error {
	defer func() {
		if err := recover(); err != nil {
			if err == "OK" {
				return
			}
			panic(err)
		}
	}()

	var tmp *os.File
	defer (func() {
		os.Remove(tmp.Name())
	})()

	setupViper(&tmp)
	err := rootCmd.Execute()

	if err != nil {
		return err
	}

	teardownViper(tmp)

	return err
}

func init() {
}

func setupViper(tmp **os.File) {
	config.Import(tmp)

	viper.SetConfigFile((*tmp).Name())
	if err := viper.ReadInConfig(); err != nil {
		// Try to regenerate config file
		fmt.Println("Could not read config file.", err, ". Trying to rebuild the config file...")
		os.Remove((*tmp).Name())
		config.GenerateConfiguration()
		config.Import(tmp)
		viper.SetConfigFile((*tmp).Name())

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Rebuild failed")
			panic(err)
		}

		fmt.Println("Rebuild successful")
	}
}

func teardownViper(tmp *os.File) {
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
	}

	config.Export(tmp)
}
