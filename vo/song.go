package vo

import (
	"github.com/96368a/LuoYiMusic-server-api/model"
	"gorm.io/datatypes"
)

type SongInfoVo struct {
	ID      uint64         `json:"id" gorm:"primary_key;"` // 歌曲id
	Name    string         `json:"name"`                   // 歌曲名
	Album   uint64         `json:"album"`                  // 所属专辑id
	Artists datatypes.JSON `json:"artists" gorm:"json"`    // 歌手列表id
	Alias   datatypes.JSON `json:"alias" gorm:"json"`      // 歌曲别名
	Hash    string
}

func ToSongInfoVo(song model.Song) SongInfoVo {
	return SongInfoVo{
		ID:      song.ID,
		Name:    song.Name,
		Album:   song.Album,
		Artists: song.Artists,
		Alias:   song.Alias,
	}
}
func ToSongInfoVos(songs []model.Song) []SongInfoVo {
	songInfos := make([]SongInfoVo, len(songs))
	for i, song := range songs {
		songInfos[i] = ToSongInfoVo(song)
	}
	return songInfos
}
