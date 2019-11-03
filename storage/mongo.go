package storage

import (
	"github.com/globalsign/mgo"
	"github.com/tacnoman/mustard-api/core"
)

var db *mgo.Database

func Db() *mgo.Database {
	if db == nil {
		s, err := mgo.Dial(core.GetEnv("MONGODB_URI", "mongodb://localhost:27017"))
		if err != nil {
			panic(err)
		}

		db = s.DB(core.GetEnv("MONGODB_DATABASE", "mustard_dev"))
	}

	return db
}
