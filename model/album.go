package model

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	ID          uint64 `json:"id" gorm:"primary_key;"` // 专辑id
	ArtistID    uint64 `json:"artistId"`               // 所属歌手
	Description string `json:"description"`            // 专辑描述
	Name        string `json:"name"`                   // 专辑名
	PicID       uint64 `json:"picId"`                  // 专辑id
	PicURL      string `json:"picUrl"`                 // 专辑url
}
