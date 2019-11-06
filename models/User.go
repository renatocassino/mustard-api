package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/tacnoman/mustard-api/core"
	"github.com/tacnoman/mustard-api/dtos"
	"github.com/tacnoman/mustard-api/storage"
)

type User struct {
	ID            string    `json:"id" bson:"_id"`
	GoogleID      string    `json:"googleId" bson:"googleID"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"emailVerified" bson:"emailVerified"`
	Name          string    `json:"name"`
	GivenName     string    `json:"givenName" bson:"givenName"`
	FamilyName    string    `json:"familyName" bson:"familyName"`
	Picture       string    `json:"picture"`
	Locale        string    `json:"locale"`
	Lyrics        []Lyric   `json:"lyrics,omitempty" bson:"-"`
	CreatedAt     time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updatedAt"`
	jwt.StandardClaims
}

func getUserCollection() *mgo.Collection {
	return storage.Db().C("users")
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

func (u *User) SetDataByGoogleDTO(auth *dtos.GoogleAuthDTO) {
	u.GoogleID = auth.ID
	u.Email = auth.Email
	u.EmailVerified = auth.VerifiedEmail
	u.Name = auth.Name
	u.GivenName = auth.Name
	u.FamilyName = auth.FamilyName
	u.Picture = auth.Picture
	u.Locale = auth.Locale
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u User) FindByGoogleID(id string) *User {
	getUserCollection().Find(bson.M{"googleID": id}).One(&u)
	return &u
}

func (u *User) InsertOrUpdate(auth *dtos.GoogleAuthDTO) {
	existUser := User{}.FindByGoogleID(auth.ID)
	u.SetDataByGoogleDTO(auth)

	if existUser.ID != "" {
		u.ID = existUser.ID
		getUserCollection().Update(bson.M{"googleID": u.GoogleID}, u)
		return
	}

	u.ID = core.GenUUIDv4()
	getUserCollection().Insert(&u)
}

func (u User) LoggedUser(authorization string) (*User, error) {
	tokenstring := strings.Replace(authorization, "Bearer ", "", 1)

	user := User{}
	_, err := jwt.ParseWithClaims(tokenstring, &user, func(token *jwt.Token) (interface{}, error) {
		return []byte(core.GetEnv("JWT_TOKEN", "secret")), nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &user, nil
}
