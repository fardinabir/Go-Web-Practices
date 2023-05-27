package controllers

import (
	"fmt"
	"github.com/fardinabir/Go_CRUD_API/model"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

var signSecretKey = []byte("fardinabir")

func newToken(username string, expiry int, tokenType string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = model.TokenDetails{
		Authorized: true,
		TokenType:  tokenType,
		UserName:   username,
		Expiry:     time.Now().Add(time.Minute * time.Duration(expiry)).Unix(),
	}

	signedToken, err := token.SignedString(signSecretKey)
	if err != nil {
		fmt.Errorf("error while token signing", err.Error())
	}
	return signedToken
}

func validateToken(token string) (jwt.MapClaims, error) {
	parsedToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnauthorizedReq
		}
		return signSecretKey, nil
	})
	claims := parsedToken.Claims.(jwt.MapClaims)

	if ok := parsedToken.Valid; !ok {
		log.Println("Invalid Token")
		return nil, ErrInvalidToken
	}
	exp := int64(claims["expiry"].(float64))
	if time.Now().After(time.Unix(exp, 0)) {
		log.Println("Token Expired")
		return nil, ErrTokenExpired
	}
	return claims, nil
}
