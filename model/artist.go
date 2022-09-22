package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	ID          uint64         `json:"id" gorm:"primary_key;"`  // id
	Alias       datatypes.JSON `json:"alias" gorm:"type:json;"` // 歌手别名
	Description string         `json:"description"`             // 歌手描述
	Name        string         `json:"name"`                    // 歌手名字
	PicID       uint64         `json:"picId"`                   // 歌手图片id
	PicURL      string         `json:"picUrl"`                  // 歌手图片url
}
