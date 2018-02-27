package routes

import (
	"github.com/labstack/echo"
	"nightcode/controllers"
)

func AddUserRoutes(e *echo.Echo) {
	e.POST("/register", controllers.HandleAddUser)
	e.GET("/users", controllers.HandleGetUsers)
}
