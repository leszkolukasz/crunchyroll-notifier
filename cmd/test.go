package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func foo() {
	fmt.Println(viper.Get("arg"))
	viper.Set("arg", "co tam")
	viper.SetEnvPrefix("spf") // will be uppercased automatically
	viper.BindEnv("id")

	os.Setenv("SPF_ID", "13") // typically done outside of the app
	viper.WriteConfig()
}
