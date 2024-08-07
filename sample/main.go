package main

import (
	"time"

	"github.com/ggymm/auth"
)

func init() {
	// 存储
	s, err := NewStore()
	if err != nil {
		panic(err)
	}

	// 日志

	// 初始化
	a := auth.NewAuth()
	a.Store = s

	a.Renew = true
	a.Shared = false
	a.Concurrent = true

	a.TokenName = "company-token"
	a.LoginLimit = 10
	a.DefaultTimeout = 30 * time.Minute

	auth.SetDefault(a)
}

func main() {

}
