package connector

import (
	mgo "gopkg.in/mgo.v2"
)

func Connect(host []string) *mgo.Session {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: host,
	})
	if err != nil {
		panic(err)
	}

	return session
}
