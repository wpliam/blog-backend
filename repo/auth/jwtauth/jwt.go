package jwtauth

import (
	"blog-backend/constant"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("apple")

var DefaultJwtAuth = &JwtAuth{
	signKey: jwtKey,
	expired: constant.LoginValidTime,
	keyFunc: func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	},
}
