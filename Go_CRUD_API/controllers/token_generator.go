package controllers

import (
	"fmt"
	"github.com/fardinabir/Go_CRUD_API/model"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"log"
	"time"
)

type TokenAuth struct {
	SignSecretKey []byte
	ExpiryAccess  time.Duration
	ExpiryRefresh time.Duration
}

func newTokenAuth() *TokenAuth {
	return &TokenAuth{
		SignSecretKey: []byte(viper.GetString("auth.secret_key")),
		ExpiryAccess:  viper.GetDuration("auth.expiry.access_token"),
		ExpiryRefresh: viper.GetDuration("auth.expiry.refresh_token"),
	}
}

func (t *TokenAuth) generateTokens(userName string) model.Token {
	accToken := t.newToken(userName, t.ExpiryAccess, "access")
	refToken := t.newToken(userName, t.ExpiryRefresh, "refresh")
	return model.Token{accToken, refToken}
}

func (t *TokenAuth) newToken(username string, expiry time.Duration, tokenType string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = model.TokenDetails{
		Authorized: true,
		TokenType:  tokenType,
		UserName:   username,
		Expiry:     time.Now().Add(time.Minute * expiry).Unix(),
	}

	signedToken, err := token.SignedString(t.SignSecretKey)
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
		signSecret := []byte(viper.GetString("auth.secret_key"))
		return signSecret, nil
	})
	claims := parsedToken.Claims.(jwt.MapClaims)

	if ok := parsedToken.Valid; !ok {
		log.Println("Invalid Tokens")
		return nil, ErrInvalidToken
	}
	exp := int64(claims["expiry"].(float64))
	if time.Now().After(time.Unix(exp, 0)) {
		log.Println("Tokens Expired")
		return nil, ErrTokenExpired
	}
	return claims, nil
}
