package services

import (
	"errors"
	"fmt"
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

func SearchArtist(name string, pageSize int, page int) ([]model.Artist, int64, error) {
	var artists []model.Artist
	name = fmt.Sprintf("%%%s%%", name)
	var count int64
	db := model.DB.Model(&model.Artist{}).Where("name Like ?", name).Count(&count)
	if page < 1 {
		page = 1
	}
	if count > (int64)((page-1)*pageSize) {
		db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&artists)
	} else {
		db.Find(&artists)
	}
	return artists, count, nil
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
