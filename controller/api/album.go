package api

import (
	"github.com/96368a/LuoYiMusic-server-api/dto"
	"github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DelAlbum(c *gin.Context) {
	var album dto.AlbumDto
	c.ShouldBind(&album)
	if album.ID <= 0 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	err := services.DelAlbum(album.ID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.Success(c, nil, "删除成功")
}
