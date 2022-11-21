package controller

import (
	"github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/96368a/LuoYiMusic-server-api/vo"
	"github.com/gin-gonic/gin"
)

func RecommendSongs(c *gin.Context) {
	songs, err := services.RecommendSongs()
	if err != nil {
		utils.Fail(c, 500, "内部错误", nil)
	}
	utils.Success(c, gin.H{
		"data": vo.ToSongInfoVos(songs),
	}, "获取成功")
}
