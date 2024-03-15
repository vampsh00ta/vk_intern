package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"vk/internal/repository/models"
)

type JwtCustomClaim struct {
	models.Customer
	jwt.RegisteredClaims
}

func (s service) CreateAccessToken(customer models.Customer, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &JwtCustomClaim{
		customer,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.cfg.Secret))
	if err != nil {
		return "", err
	}
	return t, err
}
func (s service) IsAuthorized(requestToken string) (bool, error) {
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
func (s service) ExtractCustomerFromToken(requestToken string) (JwtCustomClaim, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.cfg.Secret), nil
	})

	if err != nil {
		return JwtCustomClaim{}, err
	}

	claims, ok := token.Claims.(JwtCustomClaim)

	if !ok && !token.Valid {
		return JwtCustomClaim{}, fmt.Errorf("Invalid Token")
	}

	return claims, nil
}
