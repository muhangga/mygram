package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SECREY_KEY = "F1N4LPR0J3CT"

type JwtService interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct{}

func NewJwtService() *jwtService {
	return &jwtService{}
}

type jwtCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claims := &jwtCustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECREY_KEY))
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(SECREY_KEY), nil
	})
}
