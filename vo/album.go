package vo

import (
	"github.com/96368a/LuoYiMusic-server-api/model"
	"github.com/96368a/LuoYiMusic-server-api/services"
)

type AlbumVo struct {
	ID          uint64 `json:"id" gorm:"primary_key;"` // 专辑id
	Name        string `json:"name"`                   // 专辑名
	Description string `json:"description"`            // 专辑描述
	//ArtistID    uint64   `json:"artistId"`               // 所属歌手
	Artist ArtistVo `json:"artist"` // 所属歌手
	PicID  uint64   `json:"picId"`  // 专辑id
	PicURL string   `json:"picUrl"` // 专辑url
}

func ToAlbumVo(album model.Album) AlbumVo {
	artist, err := services.ArtistByID(album.ArtistID)
	if err != nil {
		return AlbumVo{}
	}
	return AlbumVo{
		ID:          album.ID,
		Name:        album.Name,
		Description: album.Description,
		//ArtistID:    album.ArtistID,
		Artist: ToArtistVo(*artist),
		PicID:  album.PicID,
		PicURL: album.PicURL,
	}
}
func ToAlbumVos(albums []model.Album) []AlbumVo {
	albumVos := make([]AlbumVo, len(albums))
	for i, album := range albums {
		albumVos[i] = ToAlbumVo(album)
	}
	return albumVos
}
