package auth

import (
	"github.com/pkg/errors"
)

const (
	ErrAuthNotInit = "auth not init"
	InvalidLoginId = "invalid login id"
)

func NotInit() bool {
	return auth == nil
}

// Login 登录
func Login(id int64, config ...LoginConfig) (*Token, error) {
	if NotInit() {
		return nil, errors.New(ErrAuthNotInit)
	}

	// 校验参数
	if id <= 0 {
		return nil, errors.New(InvalidLoginId)
	}

	cfg := defaultLoginConfig
	if len(config) > 0 {
		cfg = config[0]
	}
	return auth.login(id, cfg)
}

// Check 检查 token 是否有效
func Check(token string) (bool, error) {
	if NotInit() {
		return false, errors.New(ErrAuthNotInit)
	}
	return auth.check(token)
}

func Logout(id int64, device ...string) error {
	if NotInit() {
		return errors.New(ErrAuthNotInit)
	}
	return auth.logout(id, device...)
}

func GetSession(token string) (any, error) {
	if NotInit() {
		return nil, errors.New(ErrAuthNotInit)
	}
	return auth.readSessionData(token)
}

func SaveSession(id int64, data any) error {
	if NotInit() {
		return errors.New(ErrAuthNotInit)
	}
	return auth.updateSessionData(id, data)
}
