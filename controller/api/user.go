package api

import (
	"github.com/96368a/LuoYiMusic-server-api/dto"
	"github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(c *gin.Context) {
	var user dto.UserDto
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
	utils.Success(c, gin.H{"user": newUser}, "用户添加成功")
}

func UpdateUser(c *gin.Context) {

}
