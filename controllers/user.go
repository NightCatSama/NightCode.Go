package controllers

import (
	"github.com/labstack/echo"
	"nightcode/proxy"
)

// 获取用户列表
func HandleGetUsers(c echo.Context) error {
	users, err := proxy.GetUSers()
	if err != nil {
		return Error(c, "失败", nil, nil)
	} else {
		return Success(c, "成功", users)
	}
}

// 添加用户
func HandleAddUser(c echo.Context) error {
	account := c.FormValue("account")
	password := c.FormValue("password")

	err := proxy.AddUser(account, password)
	if err != nil {
		return Error(c, "失败", nil, err)
	} else {
		return Success(c, "成功", nil)
	}
}
