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

func SearchPlaylist(name string, pageSize int, page int, user *model.User) ([]model.Playlist, int64, error) {
	var playlists []model.Playlist
	name = fmt.Sprintf("%%%s%%", name)
	var count int64
	db := model.DB
	if user == nil {
		db = model.DB.Model(&model.Playlist{}).Where("name Like ?", name).Or("description Like ?", name).Count(&count)
	} else if user.ID == 0 {
		db = model.DB.Model(&model.Playlist{}).Where("( name Like ? OR description Like ? ) AND status = ?", name, name, 1).Count(&count)

	} else {
		db = model.DB.Debug().Model(&model.Playlist{}).Where("( name Like ? or description Like ? ) and userId = ?", name, name, user.ID).Count(&count)
	}
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

func DelPlaylist(id uint64, user *model.User) error {
	playlist := model.Playlist{}
	db := model.DB.First(&playlist, id)
	if db.Error != nil {
		return errors.New("歌单不存在")
	}

	if user != nil && playlist.UserID != user.ID {
		return errors.New("越权！")
	}

	return db.Delete(&playlist).Error
}
