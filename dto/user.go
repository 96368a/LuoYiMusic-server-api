package dto

type UserDto struct {
	Nickname  string `json:"nickname"`  // 昵称
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 用户密码
	Signature string `json:"signature"` //用户签名
}
