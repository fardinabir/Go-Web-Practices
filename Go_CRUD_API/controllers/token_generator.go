package controllers

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var signSecretKey = []byte("fardinabir")

func newToken(username string, expiry int, tokenType string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["type"] = tokenType
	claims["username"] = username
	claims["expiry"] = time.Now().Add(time.Minute * time.Duration(expiry)).Unix()

	signedToken, err := token.SignedString(signSecretKey)
	if err != nil {
		fmt.Errorf("Error while token signing", err.Error())
	}

	return signedToken, err
}
