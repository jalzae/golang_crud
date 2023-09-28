package helper

import (
	"github.com/google/uuid"
)

func GetUUID()  uuid.UUID {
	newUUID := uuid.New()
	return newUUID
}
