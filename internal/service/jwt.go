package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"vk/internal/repository/models"
)

func DecodeToken(jwt string) {

}
func EncodeToken(jwt string) {

}

var secretKey = []byte("secret-key")

func createToken(customer models.Customer) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       customer.Id,
			"username": customer.Username,
			"admin":    customer.Admin,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// func (s service) VerifyCustomer()
type JwtCustomClaim struct {
	models.Customer
	jwt.RegisteredClaims
}

func verifyToken(tokenString string) (*JwtCustomClaim, error) {
	// pass your custom claims to the parser function
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return &JwtCustomClaim{}, err
	}
	// type-assert `Claims` into a variable of the appropriate type
	myClaims := token.Claims.(*JwtCustomClaim)
	return myClaims, nil
}
