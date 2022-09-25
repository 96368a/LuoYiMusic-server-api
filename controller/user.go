package controller

import (
	"github.com/96368a/LuoYiMusic-server-api/dto"
	"github.com/96368a/LuoYiMusic-server-api/model"
	"github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/96368a/LuoYiMusic-server-api/vo"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	if len(user.Nickname) <= 0 || len(user.Username) < 4 || len(user.Password) < 6 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	newUser, err := services.AddUser(user.Nickname, user.Username, user.Password)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	token, err := utils.ReleaseToken(*newUser)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "token生成失败", nil)
	}
	c.Header("Authorization", token)
	utils.Success(c, gin.H{"token": token}, "注册成功")
}

func Login(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	if len(user.Username) < 4 || len(user.Password) < 6 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	loginUser, err := services.Login(user.Username, user.Password)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	token, err := utils.ReleaseToken(*loginUser)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "token生成失败", nil)
	}
	c.Header("Authorization", token)
	utils.Success(c, gin.H{
		"token": token,
	}, "登录成功")
}

func UpdateUser(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	//拿到当前登录用户，并在数据库中查找
	sessionUser, _ := c.Get("user")
	var dbUser model.User
	model.DB.First(&dbUser, sessionUser.(model.User).ID)

	if user.Nickname != "" && len(user.Nickname) < 4 {
		utils.Fail(c, http.StatusBadRequest, "昵称长度至少为4", nil)
		return
	} else {
		dbUser.Nickname = user.Nickname
	}

	if user.Signature != "" {
		dbUser.Signature = user.Signature
	}

	err := services.UpdateUser(&dbUser)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "更新失败", nil)
		return
	}
	utils.Success(c, gin.H{
		"user": vo.ToUserVO(dbUser),
	}, "更新成功")
}

func ChangePassword(c *gin.Context) {
	var user dto.UserPasswordDto
	err := c.ShouldBind(&user)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	//拿到当前登录用户，并在数据库中查找
	sessionUser, _ := c.Get("user")
	var dbUser model.User
	model.DB.First(&dbUser, sessionUser.(model.User).ID)

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.OldPassword))
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "原密码不正确", nil)
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "内部错误", nil)
		return
	}
	dbUser.Password = string(hashPassword)
	err = services.UpdateUser(&dbUser)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "密码更改失败", nil)
		return
	}
	utils.Success(c, nil, "密码更新成功")
}

func UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	utils.Success(c, gin.H{
		"user": vo.ToUserVO(user.(model.User)),
	}, "获取成功")
}
