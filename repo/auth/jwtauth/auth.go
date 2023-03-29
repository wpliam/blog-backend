package jwtauth

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// JwtAuth jwt鉴权
type JwtAuth struct {
	*options
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

func (j *JwtAuth) ParseToken(ctx *gin.Context) error {
	claims, err := j.ParseClaims(ctx)
	if err != nil {
		return err
	}
	if !claims.VerifyExpiresAt(time.Now(), false) {
		return fmt.Errorf("access token expired")
	}
	return nil
}

func (j *JwtAuth) ParseUid(ctx *gin.Context) (int64, error) {
	claims, err := j.ParseClaims(ctx)
	if err != nil {
		return 0, err
	}
	return claims.Uid, nil
}

func (j *JwtAuth) ParseClaims(ctx *gin.Context) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(j.GetToken(ctx), &CustomClaims{}, j.keyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("access token invalid")
}

func (j *JwtAuth) GetToken(ctx *gin.Context) string {
	return strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
}
