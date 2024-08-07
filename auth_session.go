package auth

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

func (auth *Auth) readSession(id int64) (*Session, error) {
	var (
		err       error
		data      []byte
		session   *Session
		sessionId = fmt.Sprintf("%s:session:%d", auth.TokenName, id)
	)

	data, err = auth.Store.Get([]byte(sessionId))
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("session not found")
	}

	err = decode[Session](data, session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (auth *Auth) createSession(id int64, token *Token) (*Session, error) {
	var (
		err       error
		data      []byte
		session   *Session
		sessionId = fmt.Sprintf("%s:session:%d", auth.TokenName, id)
	)

	data, err = auth.Store.Get([]byte(sessionId))
	if err != nil {
		return nil, err
	}

	if data == nil {
		session = &Session{
			Id:         sessionId,
			UserId:     id,
			Tokens:     []*Token{token},
			CreateTime: time.Now().UnixMilli(),
		}

		// 保存 token
		return session, session.update()
	} else {
		err = decode[Session](data, session)
		if err != nil {
			return nil, err
		}

		// 添加 token
		return session, session.updateToken(token)
	}
}

func (auth *Auth) readSessionData(t string) (any, error) {
	var (
		err     error
		token   *Token
		session *Session
	)

	// 获取 token
	// 获取 session
	token, err = auth.readToken(t)
	if err != nil {
		return nil, err
	}
	session, err = auth.readSession(token.UserId)
	if err != nil {
		return nil, err
	}
	return session.UserData, nil
}

func (auth *Auth) updateSessionData(id int64, data any) error {
	// 获取 session
	session, err := auth.readSession(id)
	if err != nil {
		return err
	}

	// 更新 session data
	session.UserData = data
	return session.update()
}
