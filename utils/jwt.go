package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
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

func loadECDSAPrivateKey(path string) (*ecdsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	return x509.ParseECPrivateKey(block.Bytes)
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
	privateKeyPath := os.Getenv("ECDSA_PRIVATE_KEY_PATH")
	privateKey, err := loadECDSAPrivateKey(privateKeyPath)
	if err != nil {
		panic("failed to load ECDSA private key")
	}

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", errors.New("error signing the token")
	}

	fmt.Println("JWT Token created and signed:", signedToken)
	return signedToken, nil
}
