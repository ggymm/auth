package auth

func (auth *Auth) logout(id int64, device ...string) error {
	var (
		err     error
		tokens  []*Token
		session *Session
	)

	// 获取 session
	session, err = auth.readSession(id)
	if err != nil {
		return err
	}

	// 获取 token 列表
	if len(device) == 0 {
		tokens = session.Tokens
		session.Tokens = make([]*Token, 0)
	} else {
		for _, t := range session.Tokens {
			for i, d := range device {
				if t.Device == d {
					tokens = append(tokens, t)
					session.Tokens = append(session.Tokens[:i], session.Tokens[i+1:]...)
				}
			}
		}
	}

	// 删除 token
	// 更新 session
	err = session.update()
	if err != nil {
		return err
	}
	for _, token := range tokens {
		err = token.delete()
		if err != nil {
			return err
		}
	}

	// 检查 session
	if len(session.Tokens) == 0 {
		err = session.delete()
		if err != nil {
			return err
		}
	}
	return nil
}

func (auth *Auth) logoutToken(t string) error {
	var (
		err     error
		token   *Token
		session *Session
	)

	// 获取 token
	// 获取 session
	token, err = auth.readToken(t)
	if err != nil {
		return err
	}
	session, err = auth.readSession(token.UserId)
	if err != nil {
		return err
	}

	// 删除 token
	// 更新 session
	err = token.delete()
	if err != nil {
		return err
	}
	for i, v := range session.Tokens {
		if v == token {
			session.Tokens = append(session.Tokens[:i], session.Tokens[i+1:]...)
		}
	}
	err = session.update()
	if err != nil {
		return err
	}

	// 检查 session
	if len(session.Tokens) == 0 {
		err = session.delete()
		if err != nil {
			return err
		}
	}
	return nil
}
