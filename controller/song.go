package controller

import (
	"github.com/96368a/LuoYiMusic-server-api/dto"
	"github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/96368a/LuoYiMusic-server-api/vo"
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

func SearchSongs(c *gin.Context) {
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
	songs, count, err := services.SearchSong(name, page.PageSize, page.Page)
	if err != nil {
		utils.Fail(c, 500, "内部错误", nil)
		return
	}
	utils.Success(c, gin.H{
		"users":       vo.ToSongInfoVos(songs),
		"currentPage": page.Page,
		"pageSize":    page.PageSize,
		"total":       count,
	}, "获取成功")
}
