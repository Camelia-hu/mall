package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

func GenerateSalt() string {
	rand.Seed(int64(uint64(time.Now().UnixNano())))
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	salt := make([]byte, 16)
	for i := range salt {
		salt[i] = letters[rand.Intn(len(letters))]
	}
	return string(salt)
}

func HashPassword(password string, salt string) string {
	h := sha256.New()
	h.Write([]byte(password + salt))
	return hex.EncodeToString(h.Sum(nil))
}
