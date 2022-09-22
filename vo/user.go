package vo

import "github.com/96368a/LuoYiMusic-server-api/model"

type UserVo struct {
	Nickname string `json:"nickname"`                    // 昵称
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 用户密码
}

type UserInfoVo struct {
	ID        uint64 `json:"id"`                           //用户id
	Nickname  string `json:"nickname"`                     // 昵称
	Username  string `json:"username" binding:"required"`  // 用户名
	Signature string `json:"signature" binding:"required"` //用户签名
}

func ToUserVO(user model.User) UserInfoVo {
	return UserInfoVo{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Username:  user.Username,
		Signature: user.Signature,
	}
}
