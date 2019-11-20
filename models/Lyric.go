package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/tacnoman/mustard-api/core"
	"github.com/tacnoman/mustard-api/storage"
)

type Lyric struct {
	ID        string    `json:"id" bson:"_id"`
	Title     string    `json:"title"`
	Lyric     string    `json:"lyric"`
	UserID    string    `json:"userId" bson:"userId"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

// String is not required by pop and may be deleted
func (l Lyric) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

func getLyricCollection() *mgo.Collection {
	return storage.Db().C("lyrics")
}

func (l Lyric) GetLyrics(user User) []Lyric {
	collection := getLyricCollection()

	var lyrics []Lyric
	collection.Find(bson.M{"userId": user.ID}).All(&lyrics)

	return lyrics
}

func (l *Lyric) Create(user User) {
	l.ID = core.GenUUIDv4()
	l.CreatedAt = time.Now()
	l.UpdatedAt = time.Now()
	l.UserID = user.ID

	err := getLyricCollection().Insert(&l)

	if err != nil {
		fmt.Println(err)
	}
}

func (l *Lyric) Update(id string, user User) {
	l.UpdatedAt = time.Now()
	l.ID = id
	l.UserID = user.ID

	getLyricCollection().Update(bson.M{
		"userId": user.ID,
		"_id":    id,
	}, &l)
}

func (l *Lyric) Delete(id string, user User) {
	getLyricCollection().Remove(bson.M{
		"userId": user.ID,
		"_id":    id,
	})
}
