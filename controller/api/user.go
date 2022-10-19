package api

import (
	"errors"
	"github.com/96368a/LuoYiMusic-server-api/dto"
	"github.com/96368a/LuoYiMusic-server-api/model"
	"github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/96368a/LuoYiMusic-server-api/vo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func DelUser(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	if user.ID <= 0 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	sessionUser, _ := c.Get("user")
	if sessionUser.(model.User).ID == user.ID {
		utils.Fail(c, http.StatusBadRequest, "不能删除自己！", nil)
		return
	}

	err := services.DelUser(user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Fail(c, http.StatusBadRequest, "用户不存在", nil)
		} else {
			utils.Fail(c, http.StatusInternalServerError, "内部错误", nil)
		}
		return
	}
	utils.Success(c, nil, "删除成功")
}

func UpdateUser(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	if user.ID <= 0 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	if user.Nickname != "" && len(user.Nickname) < 4 {
		utils.Fail(c, http.StatusBadRequest, "昵称长度至少为4", nil)
		return
	}

	err := services.UpdateUser(user.ID, user.Nickname, user.Signature)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "更新失败", nil)
		return
	}
	utils.Success(c, nil, "更新成功")
}

func ChangePassword(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	if user.ID < 0 || len(user.Password) < 6 || len(user.Password) > 20 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	err := services.UpdatePassword(user.ID, user.Password)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "密码更改失败", nil)
		return
	}
	utils.Success(c, nil, "密码更新成功")
}

func GetAllUsers(c *gin.Context) {
	var page dto.Page
	err := c.ShouldBind(&page)

	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	if page.PageSize == 0 {
		page.PageSize = 5
	}
	users, count, err := services.GetAllUsers(page.PageSize, page.Page)
	//fmt.Printf("%v \n", users)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "内部错误", nil)
		return
	}
	utils.Success(c, gin.H{
		"users":       vo.ToUserVOs(users),
		"currentPage": page.Page,
		"pageSize":    page.PageSize,
		"total":       count,
	}, "获取成功")
}

func SearchUsers(c *gin.Context) {
	var page dto.Page
	err := c.ShouldBind(&page)

	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	if page.PageSize == 0 {
		page.PageSize = 5
	}
	name := c.Query("name")
	//if name == "" {
	//	utils.Fail(c, 400, "参数错误", nil)
	//	return
	//}
	users, count, err := services.SearchUsers(name, page.PageSize, page.Page)
	if err != nil {
		utils.Fail(c, 500, "内部错误", nil)
		return
	}
	utils.Success(c, gin.H{
		"users":       vo.ToUserVOs(users),
		"currentPage": page.Page,
		"pageSize":    page.PageSize,
		"total":       count,
	}, "获取成功")
}

func SetAdmin(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	if user.ID <= 0 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	err := services.SetAdminUser(user.ID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "内部错误", nil)
		return
	}
	utils.Success(c, nil, "设置成功")
}

func RemoveAdmin(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)

	if user.ID <= 0 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	sessionUser, _ := c.Get("user")
	if sessionUser.(model.User).ID == user.ID {
		utils.Fail(c, http.StatusBadRequest, "不能移除自身的管理员权限", nil)
		return
	}

	err := services.RemoveAdmin(user.ID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "内部错误", nil)
		return
	}
	utils.Success(c, nil, "设置成功")
}
