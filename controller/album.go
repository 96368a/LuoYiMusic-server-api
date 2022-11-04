package controller

import (
	"github.com/96368a/LuoYiMusic-server-api/dto"
	"github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/96368a/LuoYiMusic-server-api/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchAlbums(c *gin.Context) {
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
	albums, count, err := services.SearchAlbum(name, page.PageSize, page.Page)
	if err != nil {
		utils.Fail(c, 500, "内部错误", nil)
		return
	}
	utils.Success(c, gin.H{
		"users":       vo.ToAlbumVos(albums),
		"currentPage": page.Page,
		"pageSize":    page.PageSize,
		"total":       count,
	}, "获取成功")
}
