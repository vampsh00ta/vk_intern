package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
	"vk/internal/repository/models"
)

type JwtCustomClaim struct {
	jwt.RegisteredClaims
	models.Customer
}

func (s service) CreateAccessToken(customer models.Customer, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &JwtCustomClaim{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		customer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.cfg.Secret))
	if err != nil {
		return "", err
	}
	return t, err
}
func (s service) IsAuthorized(requestToken string) (bool, error) {
	splited := strings.Split(requestToken, " ")
	if len(splited) != 2 || splited[0] != "Bearer" {
		return false, fmt.Errorf("Invalid Token")
	}
	requestToken = splited[1]
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.cfg.Secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
func (s service) ExtractCustomerFromToken(requestToken string) (*JwtCustomClaim, error) {
	splited := strings.Split(requestToken, " ")
	if len(splited) != 2 || splited[0] != "Bearer" {
		return nil, fmt.Errorf("Invalid Token")
	}
	requestToken = splited[1]
	token, err := jwt.ParseWithClaims(requestToken, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.cfg.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaim)
	if !ok && !token.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}
	return claims, nil
}
