package api

import (
	"fmt"
	"github.com/96368a/LuoYiMusic-server-api/dto"
	"github.com/96368a/LuoYiMusic-server-api/model"
	services "github.com/96368a/LuoYiMusic-server-api/services"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/dhowden/tag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func SongUploads(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		log.Println(file.Filename)
		//取md5作为文件名
		fileHash := utils.GetFileMd5(file)
		dst := fmt.Sprintf("./resources/musics/%s", fileHash)
		// 上传文件至指定目录
		f, _ := file.Open()
		songInfo, _ := tag.ReadFrom(f)
		//检查是否为正确的歌曲文件
		if songInfo != nil {

			//依次检查歌手、专辑、歌曲是否存在，不存在则新建
			artists := strings.Split(songInfo.Artist(), "/")
			artistIds := make([]uint64, len(artists))
			var artist *model.Artist
			//检查歌手列表
			for i, name := range artists {
				art, ok := services.CheckArtist(name)
				if !ok {
					art, _ = services.AddArtist(name)
				}
				artistIds[i] = art.ID
				if i == 0 {
					artist = art
				}
			}
			//检查专辑
			album, ok := services.CheckAlbum(songInfo.Album())
			if !ok {
				album, _ = services.AddAlbum(songInfo.Album(), artist.ID)
			}

			song, ok := services.CheckSong(songInfo.Title(), album.ID)
			if !ok {
				song, _ = services.AddSong(songInfo.Title(), album.ID, artistIds, fileHash)
			}

			fmt.Printf("%v\n", song)
			err := c.SaveUploadedFile(file, dst)

			if err != nil {
				utils.Fail(c, http.StatusInternalServerError, "内部错误", nil)
				return
			}
		}
	}
	utils.Success(c, gin.H{
		"data": len(files),
	}, "上传成功")
}

func DelSong(c *gin.Context) {
	var song dto.SongDto
	c.ShouldBind(&song)
	if song.ID <= 0 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	err := services.DelSong(song.ID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.Success(c, nil, "删除成功")
}
