package vo

import (
	"github.com/96368a/LuoYiMusic-server-api/model"
	"github.com/96368a/LuoYiMusic-server-api/services"
	"gorm.io/datatypes"
)

type SongInfoVo struct {
	ID      uint64         `json:"id" gorm:"primary_key;"` // 歌曲id
	Name    string         `json:"name"`                   // 歌曲名
	Album   AlbumVo        `json:"album"`                  // 所属专辑id
	Artists []ArtistVo     `json:"artists" gorm:"json"`    // 歌手列表id
	Alias   datatypes.JSON `json:"alias" gorm:"json"`      // 歌曲别名
	Hash    string
}

func ToSongInfoVo(song model.Song) SongInfoVo {
	s := services.GetSongInfo(song)
	return SongInfoVo{
		ID:      s.ID,
		Name:    s.Name,
		Album:   ToAlbumVo(s.Album),
		Artists: ToArtistVos(s.Artists),
		Alias:   s.Alias,
	}
}
func ToSongInfoVos(songs []model.Song) []SongInfoVo {
	songInfos := make([]SongInfoVo, len(songs))
	for i, song := range songs {
		songInfos[i] = ToSongInfoVo(song)
	}
	return songInfos
}
