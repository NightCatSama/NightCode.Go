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
		fmt.Println("[Error]: ", err)
		os.Exit(1)
	}

	fmt.Println("[Using config file]: ", viper.ConfigFileUsed())

	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config changed")
	})
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().Bool("debug", true, "enable/disable debug mode")

	viper.BindPFlags(rootCmd.PersistentFlags())
	viper.BindPFlags(rootCmd.Flags())

	rootCmd.AddCommand(serverCmd)
}
