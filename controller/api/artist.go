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
