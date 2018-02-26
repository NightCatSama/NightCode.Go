package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"

	"github.com/spf13/viper"
)

var S *mgo.Session
var UserCollection *mgo.Collection
var TestCollection *mgo.Collection

// 链接数据库
func LinkDb() {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	db := viper.GetString("database.db")
	uri := host + ":" + port
	S, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	S.SetMode(mgo.Monotonic, true)

	fmt.Printf("[MongoDB is connected]: %s/%s \n", uri, db)

	UserCollection = S.DB(db).C("user")
	TestCollection = S.DB(db).C("test")
}

func CloseDb() {
	S.Close()
}
