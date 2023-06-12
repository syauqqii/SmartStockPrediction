package Utils

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("ytta1234ytta123ytta12ytta1ytta")

type JWTClaim struct {
	Username string
	Role     string
	jwt.RegisteredClaims
}