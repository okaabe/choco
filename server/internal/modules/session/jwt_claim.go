package session

import (
	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	UUID string
	jwt.StandardClaims
}
