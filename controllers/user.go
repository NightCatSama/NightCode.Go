package controllers

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"nightcode/model"
	"nightcode/proxy"
)

type RegisterRule struct {
	Account  string `valid:"alphanum~账号只能包含字母和数字,length(6|20)~账号不能少于6个字符和不多于20个字符"`
	Password string `valid:"length(6|20)~密码能不少于6个字符和不多于20个字符"`
	Email    string `valid:"email~邮箱格式不正确,optional"`
}

// 获取用户列表
func HandleGetUsers(c echo.Context) error {
	users, err := proxy.GetUsers()
	if err != nil {
		return Error(c, "查询失败", nil, nil)
	} else {
		return Success(c, "查询成功", users)
	}
}

// 添加用户
func HandleAddUser(c echo.Context) error {
	params := RegisterRule{
		Account:  c.FormValue("account"),
		Password: c.FormValue("password"),
		Email:    c.FormValue("email"),
	}

	if _, err := valid.ValidateStruct(params); err != nil {
		return Error(c, err.Error(), nil, err.Error())
	}

	if _, err := proxy.GetUserByAccount(params.Account); err == nil {
		return Error(c, "用户已存在", nil, nil)
	}

	err := proxy.AddUser(&model.User{
		Account:  params.Account,
		Password: params.Password,
		Email:    params.Email,
	})
	if err != nil {
		return Error(c, "添加失败", nil, err.Error())
	} else {
		return Success(c, "添加成功", nil)
	}
}
