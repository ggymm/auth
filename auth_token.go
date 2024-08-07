package auth

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (auth *Auth) newToken() *Token {
	// TODO: 生成不同风格的 token
	t := uuid.New().String()
	id := fmt.Sprintf("%s:token:%s", auth.TokenName, t)
	return &Token{
		Id:    id,
		Value: t,

		CreateTime: time.Now().UnixMilli(),
		UpdateTime: time.Now().UnixMilli(),
	}
}

func (auth *Auth) readToken(token string) (*Token, error) {
	key := fmt.Sprintf("%s:token:%s", auth.TokenName, token)
	data, err := auth.Store.Get(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	t := new(Token)
	err = decode[Token](data, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (auth *Auth) createToken(id int64, config LoginConfig) (*Token, error) {
	// 判断是否允许重复登录
	if auth.Concurrent {
		// 在允许重复登录的情况下
		// 需要判断是否允许共享 token
		if auth.Shared {
			// 查询是否有可用的 token
			s, err := auth.readSession(id)
			if err != nil {
				return nil, err
			}

			// 如果查询到可用的 token。直接返回
			if s != nil && len(s.Tokens) > 0 {
				token := new(Token)
				if len(config.Device) == 0 {
					token = s.Tokens[0]
				} else {
					for _, t := range s.Tokens {
						if t.Device == config.Device {
							token = t
							break
						}
					}
				}
				if token != nil {
					// 续签 token
					if auth.Renew && config.Timeout > 0 {
						err = token.updateTimeout()
						if err != nil {
							return nil, err
						}
					}
					return token, nil
				}
			}
		}
	} else {
		// 如果不允许重复登陆，需要踢出其他 token 的登陆状态（同device）
		// 需要将其他 token 的登陆状态设置为无效
		s, err := auth.readSession(id)
		if err != nil {
			return nil, err
		}

		if s != nil {
			for i, t := range s.Tokens {
				if t.Device == config.Device {
					err = t.delete()
					if err != nil {
						return nil, err
					}
					s.Tokens = append(s.Tokens[:i], s.Tokens[i+1:]...)
				}
			}
			err = s.update()
			if err != nil {
				return nil, err
			}
		}
	}

	// 生成 token
	t := auth.newToken()
	t.UserId = id
	t.Device = config.Device
	t.Timeout = config.Timeout
	return t, t.update()
}
