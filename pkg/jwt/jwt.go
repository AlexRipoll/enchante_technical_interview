package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Service interface {
	GenerateToken(id string, role string) (string, error)
	ValidateToken(tokenString string) error
	Claims() map[string]interface{}
}

type service struct {
	key string
	ttl int64
}

func NewService(key string, ttl int64) Service {
	return &service{key,ttl}
}

var tokenClaims map[string]interface{}

func (s *service) GenerateToken(id string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Second * time.Duration(s.ttl)).Unix()

	return token.SignedString([]byte(s.key))
}

func (s *service) ValidateToken(tokenString string) error {
	token, err := s.verifyToken(tokenString)
	if err != nil {
		return err
	}
	var ok bool
	if tokenClaims, ok = token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return err
	}
	return nil
}

func (s *service) verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.key), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *service) Claims() map[string]interface{} {
	return tokenClaims
}




