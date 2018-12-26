package database

import (
	"gopkg.in/mgo.v2"
)

// Session .
var Session *mgo.Session
var err error

// InitDB 初始化数据库连接.
func InitDB() {
	Session,err = mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	Session.SetPoolLimit(100)

	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)

}



