package model

import "gorm.io/gorm"

type Comment struct {
	Content string `json:"content;"`                    // 评论内容
	SongID  uint64 `json:"songId" gorm:"column:songId"` // 歌曲id
	UserID  uint64 `json:"userId" gorm:"column:userId"` // 用户id
	gorm.Model
}
