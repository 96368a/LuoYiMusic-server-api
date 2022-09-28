package model

import "gorm.io/gorm"

type Album struct {
	ID          uint64 `json:"id" gorm:"primary_key;"`          // 专辑id
	Name        string `json:"name"`                            // 专辑名
	Description string `json:"description"`                     // 专辑描述
	ArtistID    uint64 `json:"artistId" gorm:"column:artistId"` // 所属歌手
	PicID       uint64 `json:"picId" gorm:"column:picId"`       // 专辑id
	PicURL      string `json:"picUrl" gorm:"column:picUrl"`     // 专辑url
	gorm.Model
}
