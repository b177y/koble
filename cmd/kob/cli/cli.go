package cli

import (
	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var Commands []*cobra.Command

var NK *koble.Koble

// var Config koble.Config

// func SetupConfig() error {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath("$HOME/.config/koble")
// 	viper.SetDefault("driver", koble.DriverConfig{Name: "podman"})
// 	viper.SetDefault("terminal", "gnome")
// 	viper.SetDefault("launch_terms", true)
// 	viper.SetDefault("launch_shell", false)
// 	viper.SetDefault("noninteractive", false)
// 	viper.SetDefault("nocolor", false)
// 	viper.SetDefault("default_namespace", "GLOBAL")
// 	viper.SetDefault("machine_memory", 128)
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		return err
// 	}
// 	err = viper.Unmarshal(&Config)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
