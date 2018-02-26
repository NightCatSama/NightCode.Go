package cmd

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"nightcode/controllers"
	"nightcode/mongo"
)

func OpenServer() {
	// 链接数据库 "mongo/mongo.go"
	go mongo.LinkDb()
	defer mongo.CloseDb()

	e := echo.New()

	// 处理请求
	e.POST("/addUser", controllers.HandleAddUser)
	e.GET("/users", controllers.HandleGetUsers)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the credentials from HTTP request header and perform a security
			// check

			// For invalid credentials
			// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			return next(c)
		}
	})

	// 开启服务器
	port := viper.GetString("server.port")
	sPort := ":" + port
	fmt.Printf("[Server is opened]: localhost%s \n", sPort)
	e.Logger.Fatal(e.Start(sPort))
}
