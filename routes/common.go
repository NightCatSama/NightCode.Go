package routes

import (
	"github.com/labstack/echo"
)

func AddRoutes(e *echo.Echo) {
	AddUserRoutes(e)
}
