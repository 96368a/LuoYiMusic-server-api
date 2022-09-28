package services

import (
	"encoding/json"
	"github.com/96368a/LuoYiMusic-server-api/model"
)

func AddSong(name string, albumId uint64, artistIds []uint64, hash string) (*model.Song, error) {
	ids, _ := json.Marshal(artistIds)
	song := &model.Song{
		Name:    name,
		Album:   albumId,
		Artists: ids,
		Hash:    hash,
	}
	return song, model.DB.Create(song).Error
}

// CheckSong 用歌曲名、歌曲专辑名检测歌曲唯一性
func CheckSong(name string, albumId uint64) (*model.Song, bool) {
	var song *model.Song
	//ids, _ := json.Marshal(artistIds)
	return song, model.DB.Where(&model.Song{
		Name:  name,
		Album: albumId,
		//Artists: ids,
	}).First(&song).Error == nil
}

func SongById(id uint64) (*model.Song, error) {
	var song *model.Song
	return song, model.DB.First(&song, id).Error
}

func SearchSong(name string) {

}
