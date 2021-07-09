package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

func Encode(key []byte, issuer, email string) (string, error) {
	atClaims := JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    issuer,
		},
	}

	jwtInst := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return jwtInst.SignedString(key)
}

func IsExpiredAndDecode(jwtToken string, key []byte) (*JwtClaim, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaim); ok {
		if claims.ExpiresAt < time.Now().Local().Unix() {
			return nil, errors.New("expired token")
		}
		return claims, nil
	}

	return nil, errors.New("couldnt parse claims")
}
