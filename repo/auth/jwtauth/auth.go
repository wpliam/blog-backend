package jwtauth

import (
	"blog-backend/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// JwtAuth jwt鉴权
type JwtAuth struct {
	expired time.Duration
	signKey interface{}
	keyFunc jwt.Keyfunc
}

type CustomClaims struct {
	Uid int64
	jwt.RegisteredClaims
}

func (j *JwtAuth) GenToken(ctx *gin.Context, uid int64) (string, error) {
	claims := &CustomClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(j.expired),
			},
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.signKey)
}

func (j *JwtAuth) ParseClaims(ctx *gin.Context) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(util.GetToken(ctx), &CustomClaims{}, j.keyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("access token invalid")
}
