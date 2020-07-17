package crypto

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 10
)

type hash struct {
}

func Bcrypt() Hashing {
	return &hash{}
}

func (e *hash) Hash(rawString string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(rawString), cost)
	if err != nil {
		return "", errors.New(fmt.Sprintf("error when trying to hash:\n %s ", err))
	}
	hash := string(hashByte)
	return hash, nil
}

func (e *hash) CheckHash(rawString string, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(rawString)); err != nil {
		return err
	}
	return nil
}
