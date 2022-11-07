package api

import (
	"github.com/96368a/LuoYiMusic-server-api/dto"
	"github.com/96368a/LuoYiMusic-server-api/model"
	"github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/96368a/LuoYiMusic-server-api/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPlaylist(c *gin.Context) {
	var playlist dto.PlaylistDto
	c.ShouldBind(&playlist)
	if len(playlist.Name) <= 0 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	sessionUser, _ := c.Get("user")
	newPlaylist, err := services.AddPlaylist(playlist.Name, playlist.Status, sessionUser.(model.User).ID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.Success(c, gin.H{"playlist": vo.ToPlaylistVo(newPlaylist)}, "歌单添加成功")
}

func SearchPlaylist(c *gin.Context) {
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
	playlists, count, err := services.SearchPlaylist(name, page.PageSize, page.Page, nil)
	if err != nil {
		utils.Fail(c, 500, "内部错误", nil)
		return
	}
	utils.Success(c, gin.H{
		"data":        vo.ToPlaylistVos(playlists),
		"currentPage": page.Page,
		"pageSize":    page.PageSize,
		"total":       count,
	}, "获取成功")
}

func DelPlaylist(c *gin.Context) {
	var playlistDto dto.PlaylistDto
	c.ShouldBind(&playlistDto)
	if playlistDto.ID <= 0 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	err := services.DelPlaylist(playlistDto.ID, nil)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.Success(c, nil, "删除歌单成功")
}
