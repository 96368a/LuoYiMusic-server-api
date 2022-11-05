package services

import (
	"errors"
	"fmt"
	"github.com/96368a/LuoYiMusic-server-api/model"
)

func AddPlaylist(name string, status uint64, userId uint64) (model.Playlist, error) {
	playlist := model.Playlist{
		Name:   name,
		Status: status,
		UserID: userId,
	}
	return playlist, model.DB.Create(&playlist).Error
}

func SearchPlaylist(name string, pageSize int, page int) ([]model.Playlist, int64, error) {
	var playlists []model.Playlist
	name = fmt.Sprintf("%%%s%%", name)
	var count int64
	db := model.DB.Model(&model.Playlist{}).Where("name Like ?", name).Or("description Like ?", name).Count(&count)
	if page < 1 {
		page = 1
	}
	if count > (int64)((page-1)*pageSize) {
		db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&playlists)
	} else {
		db.Find(&playlists)
	}
	return playlists, count, nil
}

func DelPlaylist(id uint64) error {
	playlist := model.Playlist{}
	db := model.DB.First(&playlist, id)
	if db.Error != nil {
		return errors.New("歌单不存在")
	}

	return db.Delete(&playlist).Error
}
