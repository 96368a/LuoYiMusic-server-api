package controller

import (
	"github.com/96368a/LuoYiMusic-server-api/model"
	"github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/96368a/LuoYiMusic-server-api/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user vo.UserVo
	c.ShouldBind(&user)
	if len(user.Nickname) <= 0 && len(user.Username) < 6 && len(user.Password) < 6 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	newUser, err := services.Register(user.Nickname, user.Username, user.Password)
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
	var user vo.UserVo
	c.ShouldBind(&user)
	if len(user.Username) < 6 && len(user.Password) < 6 {
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

func UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	utils.Success(c, gin.H{
		"user": vo.ToUserVO(user.(model.User)),
	}, "获取成功")
}
