package services

import "github.com/96368a/LuoYiMusic-server-api/model"

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
