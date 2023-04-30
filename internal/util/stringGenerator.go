package util

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateString(length int) (string, error) {
	randomName := make([]byte, length)

	_, err := rand.Read(randomName)
	if err != nil {
		return "", err
	}
	fileName := hex.EncodeToString(randomName)

	return fileName, nil
}
