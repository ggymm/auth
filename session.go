package auth

import (
	"sync"
	"time"
)

type Token struct {
	sync.Mutex

	Id    string `json:"id"`    // token id
	Value string `json:"value"` // token 值

	UserId  int64         `json:"userId"`  // 用户 id
	Device  string        `json:"device"`  // 设备信息
	Timeout time.Duration `json:"timeout"` // 过期时间

	CreateTime int64 `json:"createTime"` // 创建时间
	UpdateTime int64 `json:"updateTime"` // 最后更新时间
}

func (t *Token) delete() error {
	t.Lock()
	defer t.Unlock()

	return auth.Store.Delete([]byte(t.Id))
}

func (t *Token) update() error {
	t.Lock()
	defer t.Unlock()

	t.UpdateTime = time.Now().UnixMilli()

	// 序列化
	data, err := encode[Token](t)
	if err != nil {
		return err
	}
	return auth.Store.Put([]byte(t.Id), data, t.Timeout)
}

func (t *Token) updateTimeout() error {
	t.Lock()
	defer t.Unlock()

	return auth.Store.UpdateTimeout([]byte(t.Id), t.Timeout)
}

type Session struct {
	sync.Mutex

	Id string `json:"id"` // session id

	UserId   int64 `json:"userId"`   // 用户 id
	UserData any   `json:"userData"` // 用户自定义数据

	Tokens     []*Token `json:"tokens"`     // token 列表
	CreateTime int64    `json:"createTime"` // 创建时间
	UpdateTime int64    `json:"updateTime"` // 最后更新时间
}

func (s *Session) delete() error {
	s.Lock()
	defer s.Unlock()

	return auth.Store.Delete([]byte(s.Id))
}

func (s *Session) update() error {
	s.Lock()
	defer s.Unlock()

	s.UpdateTime = time.Now().UnixMilli()

	// 序列化
	data, err := encode[Session](s)
	if err != nil {
		return err
	}
	return auth.Store.Put([]byte(s.Id), data, NeverExpire)
}

func (s *Session) cleanToken() error {
	s.Lock()
	defer s.Unlock()

	for i, token := range s.Tokens {
		data, err := auth.Store.Get([]byte(token.Id))
		if err != nil || data == nil {
			s.Tokens = append(s.Tokens[:i], s.Tokens[i+1:]...)
		}
	}
	return s.update()
}

func (s *Session) updateToken(token *Token) error {
	s.Lock()
	defer s.Unlock()

	s.Tokens = append(s.Tokens, token)
	return s.update()
}
