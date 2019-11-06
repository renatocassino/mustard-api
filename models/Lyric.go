package models

import (
	"encoding/json"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/tacnoman/mustard-api/storage"
)

type Lyric struct {
	ID        string    `json:"id" bson:"_id"`
	Title     string    `json:"title"`
	Lyric     string    `json:"lyric"`
	UserID    string    `json:"userId" bson:"userId"`
	CreatedAt time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt time.Time `json:"updated_at" bson:"updatedAt"`
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

	lyrics := []Lyric{}
	collection.Find(bson.M{"userId": user.ID}).All(&lyrics)

	return lyrics
}

// Lyrics is not required by pop and may be deleted
type Lyrics []Lyric

// String is not required by pop and may be deleted
func (l Lyrics) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *Lyric) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *Lyric) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *Lyric) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
