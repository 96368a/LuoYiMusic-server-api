package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	ID              uint64         `json:"id" gorm:"primary_key;"` // 歌单id
	CommentCount    uint64         `json:"commentCount"`           // 评论数
	CoverImgID      uint64         `json:"coverImgId"`             // 歌单封面
	CoverImgURL     string         `json:"coverImgUrl"`            // 封面url
	CreateTime      uint64         `json:"createTime"`             // 创建时间
	Description     string         `json:"description"`            // 描述
	Name            string         `json:"name"`                   // 歌单名
	PlayCount       uint64         `json:"playCount"`              // 播放次数
	Status          uint64         `json:"status"`
	SubscribedCount uint64         `json:"subscribedCount"`       // 订阅次数
	Tags            datatypes.JSON `json:"tags" gorm:"type:json"` // 标签
	UpdateTime      uint64         `json:"updateTime"`            // 最后更新时间
	UserID          uint64         `json:"userId"`                // 创建用户id
}

type PlaylistItems struct {
	gorm.Model
	PlaylistID uint64 `json:"playlistId"` // 歌单id
	SongID     uint64 `json:"songId"`     // 歌曲id
	UserID     uint64 `json:"userId"`     // 添加用户id
}
