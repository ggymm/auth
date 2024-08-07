package auth

import (
	"time"
)

var auth *Auth

type Auth struct {
	Log   Log
	Store Store

	Renew      bool // 是否自动更新 token 的过期时间（续签）
	Shared     bool // 是否共享 token
	Concurrent bool // 是否允许并发登录

	TokenName      string        // token 名称（like: company-token）
	LoginLimit     int           // 最大登录数（允许并发登陆，非共享 token）
	DefaultTimeout time.Duration // token 过期时间（秒）
}

func NewAuth() *Auth {
	return &Auth{}
}

func SetDefault(a *Auth) {
	auth = a
}
