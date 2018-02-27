package proxy

import (
	"gopkg.in/mgo.v2/bson"
	"time"

	"nightcode/model"
)

// 添加用户
func AddUser(u *model.User) error {
	err := model.UserCollection.Insert(&model.User{
		Account:    u.Account,
		Password:   u.Password,
		Email:      u.Email,
		CreateTime: time.Now(),
	})
	return err
}

// 获取用户列表
func GetUsers() ([]model.User, error) {
	users := []model.User{}
	err := model.UserCollection.Find(bson.M{}).All(&users)
	return users, err
}

// 根据账号搜索用户
func GetUserByAccount(account string) (model.User, error) {
	user := model.User{}
	err := model.UserCollection.Find(bson.M{"account": account}).One(&user)
	return user, err
}

// 删除用户
func RmoveUserByAccount(account string) error {
	err := model.UserCollection.Remove(bson.M{"account": account})
	return err
}
