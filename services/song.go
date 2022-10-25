package services

import (
	"encoding/json"
	"fmt"
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

func SearchSong(name string, pageSize int, page int) ([]model.Song, int64, error) {
	var songs []model.Song
	name = fmt.Sprintf("%%%s%%", name)
	if page < 1 {
		page = 1
	}
	var count int64
	db := model.DB.Raw("SELECT songs.* from songs WHERE songs.name like ? union SELECT songs.* from songs,json_each(songs.artists) v1 WHERE v1.value = (SELECT id FROM artists WHERE name like ?)", name, name).Scan(&songs)
	if db.Error != nil {
		return nil, 0, db.Error
	}
	count = int64(len(songs))
	if count > (int64)(pageSize) {
		songs = songs[(page-1)*pageSize : pageSize]
	}
	return songs, count, nil
}

func GetSongInfo(song model.Song) model.SongInfo {
	songInfo := model.SongInfo{
		ID:    song.ID,
		Name:  song.Name,
		Alias: song.Alias,
		Hash:  song.Hash,
	}
	var artistIds []uint64
	json.Unmarshal(song.Artists, &artistIds)
	model.DB.Where("id = ?", song.Album).Find(&songInfo.Album)
	model.DB.Find(&songInfo.Artists, artistIds)
	return songInfo
}
func GetSongInfos(songs []model.Song) []model.SongInfo {
	songInfos := make([]model.SongInfo, len(songs))
	for i, song := range songs {
		songInfos[i] = GetSongInfo(song)
	}
	return songInfos
}
