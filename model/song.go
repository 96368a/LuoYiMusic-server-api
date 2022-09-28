package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Song struct {
	ID      uint64         `json:"id" gorm:"primary_key;"` // 歌曲id
	Name    string         `json:"name"`                   // 歌曲名
	Album   uint64         `json:"album"`                  // 所属专辑id
	Artists datatypes.JSON `json:"artists" gorm:"json"`    // 歌手列表id
	Alias   datatypes.JSON `json:"alias" gorm:"json"`      // 歌曲别名
	Hash    string
	gorm.Model
}
