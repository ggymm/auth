package auth

import (
	"time"
)

var defaultLoginConfig = LoginConfig{
	Device:  "web",
	Timeout: 30 * time.Minute,
}

type LoginConfig struct {
	Device  string        // 设备信息
	Timeout time.Duration // 超时时间（单位s）
}

func (auth *Auth) login(id int64, config LoginConfig) (*Token, error) {
	var (
		err     error
		token   *Token
		session *Session
	)

	// 创建 token
	token, err = auth.createToken(id, config)
	if err != nil {
		return nil, err
	}

	// 创建 session
	session, err = auth.createSession(id, token)
	if err != nil {
		return nil, err
	}

	// 判断是否超过了最大登陆数
	if auth.LoginLimit > 0 && len(session.Tokens) > auth.LoginLimit {
		// TODO
		// 如果超过，选择以下策略删除 token
		// 1）删除最先登陆
		// 2) 删除最先过期
		// 3）删除最不活跃
	}
	return token, nil
}
