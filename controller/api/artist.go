package api

import (
	"github.com/96368a/LuoYiMusic-server-api/dto"
	"github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddArtist(c *gin.Context) {
	var artist dto.ArtistDto
	c.ShouldBind(&artist)
	err, _ := services.AddArtist(artist.Name)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "内部错误", nil)
		return
	}
	utils.Success(c, nil, "添加成功")
}

func DelArtist(c *gin.Context) {
	var artist dto.ArtistDto
	c.ShouldBind(&artist)
	if artist.ID <= 0 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	err := services.DelArtist(artist.ID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.Success(c, nil, "删除成功")
}
