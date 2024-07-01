package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaim struct {
	Id    string `json:"id"`
	User  string `json:"user"`
	Admin bool   `json:"role"`
	jwt.RegisteredClaims
}

func CreateNewAuthToken(id string, email string, isAdmin bool) (string, error) {
	claim := AuthClaim{
		Id:    id,
		User:  email,
		Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "searchengine.com",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)

	secretKey, exists := os.LookupEnv("SECRET_KEY")

	if !exists {
		panic("secret key cannot be found in .env")
	}

	signedToke, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", errors.New("error signing the token")
	}

	return signedToke, nil
}
