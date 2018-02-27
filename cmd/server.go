package cmd

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"nightcode/controllers"
	"nightcode/model"
	"nightcode/routes"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Long:  `Use echo library to run a server.`,
	Run: func(cmd *cobra.Command, args []string) {
		OpenServer()
	},
}

func OpenServer() {
	// 链接数据库
	go model.LinkDb()
	defer model.CloseDb()

	e := echo.New()

	e.Debug = viper.GetBool("debug")

	// 自定义错误处理
	e.HTTPErrorHandler = controllers.HttpErrorHandler

	// 处理请求
	routes.AddUserRoutes(e)

	// 开启服务器
	port := viper.GetString("server.port")
	sPort := ":" + port
	fmt.Printf("[Server is opened]: localhost%s \n", sPort)
	e.Logger.Fatal(e.Start(sPort))
}
