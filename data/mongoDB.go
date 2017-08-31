package data

import (
	"gopkg.in/mgo.v2"
)

func InitSession()(session *mgo.Session,err error){
	session, err = mgo.Dial("127.0.0.1:27017")
	session.SetMode(mgo.Monotonic, true)
	return
}
