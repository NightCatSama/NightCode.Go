package model

import (
	"fmt"
	"gopkg.in/mgo.v2"

	"github.com/spf13/viper"
)

var S *mgo.Session
var (
	UserCollection,
	TestCollection *mgo.Collection
)

// 链接数据库
func LinkDb() {
	config := viper.GetStringMap("database")
	uri := config["host"].(string) + ":" + config["port"].(string)
	S, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}

	S.SetMode(mgo.Monotonic, false)
	S.SetSafe(&mgo.Safe{})

	fmt.Printf("[MongoDB is connected]: %s \n", uri)

	db := S.DB(config["db"].(string))

	UserCollection = db.C("user")
	TestCollection = db.C("test")
}

func CloseDb() {
	UserCollection = nil
	TestCollection = nil
	S.Close()
}
