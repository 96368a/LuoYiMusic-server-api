package services

import (
	"errors"
	"github.com/96368a/LuoYiMusic-server-api/model"
)

func AddArtist(name string) (*model.Artist, error) {
	artist := &model.Artist{
		Name: name,
	}
	return artist, model.DB.Create(artist).Error
}

func CheckArtist(name string) (*model.Artist, bool) {
	var artist model.Artist
	return &artist, model.DB.First(&artist, "name = ?", name).Error == nil
}

func DelArtist(id uint64) error {
	artist := model.Artist{}
	db := model.DB.First(&artist, id)
	if db.Error != nil {
		return errors.New("歌手不存在")
	}
	//db.Delete(&model.Song{}, "", artist.ID)
	return db.Delete(&artist).Error
}
