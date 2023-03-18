package jwtauth

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JwtAuth struct {
	*options
}

type CustomClaims struct {
	UserID int64
	jwt.RegisteredClaims
}

func (j *JwtAuth) GenToken(ctx *gin.Context, userID int64) (string, error) {
	claims := &CustomClaims{
		UserID: userID,
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
	//token, err := rdb.Get(ctx, util.GetLoginKey(claims.UserID))
	//if err != nil {
	//	return err
	//}
	//if token != j.GetToken(ctx) {
	//	return fmt.Errorf("登录信息不匹配")
	//}
	return nil
}

func (j *JwtAuth) ParseUserID(ctx *gin.Context) (int64, error) {
	claims, err := j.ParseClaims(ctx)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
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
