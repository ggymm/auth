package auth

func (auth *Auth) check(t string) (bool, error) {
	var (
		err   error
		token *Token
	)

	// 获取 token 信息
	token, err = auth.readToken(t)
	if err != nil {
		return false, err
	}

	// 更新 token 的活跃时间
	err = token.update()
	if err != nil {
		return false, err
	}

	// 更新 token 的过期时间（续签）
	if auth.Renew && token.Timeout > 0 {
		err = token.updateTimeout()
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
