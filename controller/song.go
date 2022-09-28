package controller

import (
	"github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SongById(c *gin.Context) {
	songId := c.Query("id")
	id, err := strconv.ParseInt(songId, 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "内部错误", nil)
		return
	}
	song, err := services.SongById(uint64(id))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "内部错误1", nil)
		return
	}
	c.Header("Content-Type", "audio/mpeg")
	c.File("resources/musics/" + song.Hash)
}
