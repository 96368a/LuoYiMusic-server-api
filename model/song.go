package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	ID      uint64         `json:"id" gorm:"primary_key;"` // 歌曲id
	Album   uint64         `json:"album"`                  // 所属专辑id
	Alias   datatypes.JSON `json:"alias" gorm:"json"`      // 歌曲别名
	Artists datatypes.JSON `json:"artists" gorm:"json"`    // 歌手列表id
	Name    string         `json:"name"`                   // 歌曲名
}
