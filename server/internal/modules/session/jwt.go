package session

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func EncodeToken(key []byte, issuer, uuid string) (string, error) {
	atClaims := JWTClaim{
		UUID: uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    issuer,
		},
	}

	jwtInst := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return jwtInst.SignedString(key)
}

func DecodeToken(jwtToken string, key []byte) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaim); ok {
		if claims.ExpiresAt < time.Now().Local().Unix() {
			return nil, errors.New("Looks like that's an expired token")
		}

		return claims, nil
	}

	return nil, errors.New("Unfortunately it wasn't possible create the json web token of the session")
}
