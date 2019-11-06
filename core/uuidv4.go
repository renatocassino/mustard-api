package core

import (
	"fmt"

	uuid "github.com/gofrs/uuid"
)

func GenUUIDv4() string {
	id, _ := uuid.NewV4()
	return fmt.Sprintf("%s", id)
}
