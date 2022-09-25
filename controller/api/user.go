package api

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

func AddUser(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	if len(user.Nickname) <= 0 && len(user.Username) < 4 && len(user.Password) < 6 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	newUser, err := services.AddUser(user.Nickname, user.Username, user.Password)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.Success(c, gin.H{"user": newUser}, "用户添加成功")
}

func UpdateUser(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	if user.ID <= 0 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	var dbUser model.User
	err := model.DB.First(&dbUser, user.ID).Error
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "用户不存在", nil)
		return
	}
	if user.Nickname != "" && len(user.Nickname) < 4 {
		utils.Fail(c, http.StatusBadRequest, "昵称长度至少为4", nil)
		return
	} else {
		dbUser.Nickname = user.Nickname
	}

	if user.Signature != "" {
		dbUser.Signature = user.Signature
	}

	err = services.UpdateUser(&dbUser)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "更新失败", nil)
		return
	}
	utils.Success(c, gin.H{
		"user": vo.ToUserVO(dbUser),
	}, "更新成功")
}

func ChangePassword(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	if user.ID < 0 || len(user.Password) < 6 || len(user.Password) > 20 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	//根据id查找用户
	var dbUser model.User
	err := model.DB.First(&dbUser, user.ID).Error
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "用户不存在", nil)
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
