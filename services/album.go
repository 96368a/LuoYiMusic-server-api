package services

import (
	"errors"
	"github.com/96368a/LuoYiMusic-server-api/model"
)

func AddAlbum(name string, artistId uint64) (*model.Album, error) {
	album := &model.Album{
		Name:     name,
		ArtistID: artistId,
	}
	return album, model.DB.Create(album).Error
}

func CheckAlbum(name string) (*model.Album, bool) {
	var album *model.Album
	return album, model.DB.First(&album, "name = ?", name).Error == nil
}

func DelAlbum(id uint64) error {
	album := model.Album{}
	db := model.DB.First(&album, id)
	if db.Error != nil {
		return errors.New("专辑不存在")
	}
	err := model.DB.Delete(&model.Song{}, "album = ?", album.ID).Error
	if err != nil {
		return err
	}
	return db.Delete(&album).Error
}
