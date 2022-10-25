package vo

import (
	"github.com/96368a/LuoYiMusic-server-api/model"
	"gorm.io/datatypes"
)

type ArtistVo struct {
	ID          uint64         `json:"id" gorm:"primary_key;"`      // id
	Name        string         `json:"name"`                        // 歌手名字
	Description string         `json:"description"`                 // 歌手描述
	Alias       datatypes.JSON `json:"alias" gorm:"type:json;"`     // 歌手别名
	PicID       uint64         `json:"picId" gorm:"column:picId"`   // 歌手图片id
	PicURL      string         `json:"picUrl" gorm:"column:picUrl"` // 歌手图片url
}

func ToArtistVo(artist model.Artist) ArtistVo {
	return ArtistVo{
		ID:          artist.ID,
		Name:        artist.Name,
		Description: artist.Description,
		Alias:       artist.Alias,
		PicID:       artist.PicID,
		PicURL:      artist.PicURL,
	}
}
func ToArtistVos(artists []model.Artist) []ArtistVo {
	artistVos := make([]ArtistVo, len(artists))
	for i, artist := range artists {
		artistVos[i] = ToArtistVo(artist)
	}
	return artistVos
}
