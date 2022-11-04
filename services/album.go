package services

import (
	"errors"
	"fmt"
	"github.com/96368a/LuoYiMusic-server-api/model"
)

func AddAlbum(name string, artistId uint64) (*model.Album, error) {
	album := &model.Album{
		Name:     name,
		ArtistID: artistId,
	}
	return album, model.DB.Create(album).Error
}

func SearchAlbum(name string, pageSize int, page int) ([]model.Album, int64, error) {
	var albums []model.Album
	name = fmt.Sprintf("%%%s%%", name)
	var count int64
	db := model.DB.Model(&model.Album{}).Where("name Like ?", name).Count(&count)
	if page < 1 {
		page = 1
	}
	if count > (int64)((page-1)*pageSize) {
		db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&albums)
	} else {
		db.Find(&albums)
	}
	return albums, count, nil
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
