package services

import (
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
