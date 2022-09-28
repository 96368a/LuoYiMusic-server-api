package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Artist struct {
	ID          uint64         `json:"id" gorm:"primary_key;"`      // id
	Name        string         `json:"name"`                        // 歌手名字
	Description string         `json:"description"`                 // 歌手描述
	Alias       datatypes.JSON `json:"alias" gorm:"type:json;"`     // 歌手别名
	PicID       uint64         `json:"picId" gorm:"column:picId"`   // 歌手图片id
	PicURL      string         `json:"picUrl" gorm:"column:picUrl"` // 歌手图片url
	gorm.Model
}
