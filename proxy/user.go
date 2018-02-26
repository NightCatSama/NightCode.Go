package proxy

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"

	"nightcode/mongo"
)

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// 添加用户
func AddUser(account string, password string) error {
	if mongo.UserCollection == nil {
		return fmt.Errorf("未连接数据库")
	}

	if _, err := GetUserByAccount(account); err == nil {
		return fmt.Errorf("用户已存在")
	}

	err := mongo.UserCollection.Insert(&User{account, password})
	return err
}

// 获取用户列表
func GetUSers() ([]User, error) {
	users := []User{}
	if mongo.UserCollection == nil {
		return users, fmt.Errorf("未连接数据库")
	}
	err := mongo.UserCollection.Find(bson.M{}).All(&users)
	return users, err
}

// 根据账号搜索用户
func GetUserByAccount(account string) (User, error) {
	user := User{}
	if mongo.UserCollection == nil {
		return user, fmt.Errorf("未连接数据库")
	}
	err := mongo.UserCollection.Find(bson.M{"account": account}).One(&user)
	return user, err
}

// 删除用户
func RmoveUserByAccount(account string) error {
	err := mongo.UserCollection.Remove(bson.M{"account": account})
	return err
}
