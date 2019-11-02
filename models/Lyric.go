package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

type Lyric struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Lyric     string    `json:"lyric" db:"lyric"`
	UserID    uuid.UUID `json:"userId" db:"user_id"`
	User      *User     `json:"user,omitempty" belongs_to:"user"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (l Lyric) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
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
