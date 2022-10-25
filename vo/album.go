package vo

import "github.com/96368a/LuoYiMusic-server-api/model"

type AlbumVo struct {
	ID          uint64 `json:"id" gorm:"primary_key;"`          // 专辑id
	Name        string `json:"name"`                            // 专辑名
	Description string `json:"description"`                     // 专辑描述
	ArtistID    uint64 `json:"artistId" gorm:"column:artistId"` // 所属歌手
	PicID       uint64 `json:"picId" gorm:"column:picId"`       // 专辑id
	PicURL      string `json:"picUrl" gorm:"column:picUrl"`     // 专辑url
}

func ToAlbumVo(album model.Album) AlbumVo {
	return AlbumVo{
		ID:          album.ID,
		Name:        album.Name,
		Description: album.Description,
		ArtistID:    album.ArtistID,
		PicID:       album.PicID,
		PicURL:      album.PicURL,
	}
}
func ToAlbumVos(albums []model.Album) []AlbumVo {
	albumVos := make([]AlbumVo, len(albums))
	for i, album := range albums {
		albumVos[i] = ToAlbumVo(album)
	}
	return albumVos
}
