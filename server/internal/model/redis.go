package model

import "errors"

var ErrRedisNotFound = errors.New("redis not found")

const (
	UserPrefix = "user:"
	//SET user:code:<user_id>:<email>
	KeyEmailCode = UserPrefix + "code:%d:%s"
	//SET user:profile:<user_id>
	KeyProfile = UserPrefix + "profile:%d"
	//SET user:token:<user_id>
	KeyToken = UserPrefix + "token:%d"
	//SET user:online:<user_id>
	KeyOnline = UserPrefix + "online:%d"
	//SET user:last_login:<user_id>
	KeyLastLogin = UserPrefix + "last_login:%d"
)
