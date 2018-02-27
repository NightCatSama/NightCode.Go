package test

import (
	"math/rand"
	"nightcode/model"
	"nightcode/proxy"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// 获取随机字符串
func randStringBytes(n int) string {
	b := make([]byte, n)
	rand.Seed(42)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestUserActions(t *testing.T) {
	Convey("链接数据库用户表", t, func() {
		So(model.UserCollection, ShouldNotBeNil)

		Convey("获取用户列表", func() {
			_, err := proxy.GetUsers()
			So(err, ShouldBeNil)
		})

		account := randStringBytes(8)
		password := randStringBytes(8)
		Convey("添加用户", func() {
			err := proxy.AddUser(&model.User{
				Account:  account,
				Password: password,
			})
			So(err, ShouldBeNil)
		})

		Convey("通过用户名查询用户", func() {
			_, err := proxy.GetUserByAccount(account)
			So(err, ShouldBeNil)
		})

		Convey("用过用户名删除用户", func() {
			err := proxy.RmoveUserByAccount(account)
			So(err, ShouldBeNil)
		})
	})
}
