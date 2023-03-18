package jwtauth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var DefaultJwtAuth = New()
var jwtKey = []byte("apple")

var defaultOptions = &options{
	signKey: jwtKey,
	expired: 24 * time.Hour,
	keyFunc: func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	},
}

func New(opts ...Option) *JwtAuth {
	o := defaultOptions
	for _, opt := range opts {
		opt(o)
	}
	return &JwtAuth{
		options: o,
	}
}

type options struct {
	expired time.Duration
	signKey interface{}
	keyFunc jwt.Keyfunc
}

type Option func(*options)

func WithSignKey(key interface{}) Option {
	return func(o *options) {
		o.signKey = key
	}
}

func WithExpired(expired time.Duration) Option {
	return func(o *options) {
		o.expired = expired
	}
}

func WithKeyFunc(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyFunc = keyFunc
	}
}
