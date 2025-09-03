package utils

import "github.com/google/uuid"

func GenerateUUID() string {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return uuid.NewString()
	}
	return uuidV7.String()
}
