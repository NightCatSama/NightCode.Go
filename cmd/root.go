package cmd

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const cfgFile = "$GOPATH/src/nightcode/"

var rootCmd = &cobra.Command{
	Use:   "nightcode",
	Short: "NightCode is a website that can record code",
	Long: `A convenient and beautiful code record website
						and learn golang.`,
	Run: func(cmd *cobra.Command, args []string) {
		OpenServer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("using config file", viper.ConfigFileUsed())

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config changed")
	})
}

func init() {
	cobra.OnInitialize(initConfig)

	viper.BindPFlags(rootCmd.PersistentFlags())
	viper.BindPFlags(rootCmd.Flags())
}
