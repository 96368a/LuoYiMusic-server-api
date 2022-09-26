package dto

type UserDto struct {
	ID        uint64 `json:"id"`        //用户id
	Nickname  string `json:"nickname"`  // 昵称
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 用户密码
	Signature string `json:"signature"` //用户签名
}

type UserPasswordDto struct {
	OldPassword string `json:"oldPassword" binding:"required,min=6,max=20"` // 用户密码
	Password    string `json:"password" binding:"required,min=6,max=20"`    // 用户密码
}

type Page struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"pageSize"`
}
