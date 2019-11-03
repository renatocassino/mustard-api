package core

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func GenUUIDv4() string {
	id, _ := uuid.NewV4()
	return fmt.Sprintf("%s", id)
}
