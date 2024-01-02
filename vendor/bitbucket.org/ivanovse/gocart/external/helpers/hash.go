package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func PasswordHashVerify(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateHashSession() string {
	t := time.Now()
	hashBytes := md5.Sum([]byte(t.Format("20060102150405")))
	//expiration := time.Now().Add(365 * 24 * time.Hour)
	return hex.EncodeToString(hashBytes[:])
}
