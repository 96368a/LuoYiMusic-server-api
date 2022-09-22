package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content;"` // 评论内容
	SongID  uint64 `json:"songId"`   // 歌曲id
	UserID  uint64 `json:"userId"`   // 用户id
}
