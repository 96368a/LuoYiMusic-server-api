package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Playlist struct {
	ID              uint64         `json:"id" gorm:"primary_key;"`                  // 歌单id
	Name            string         `json:"name"`                                    // 歌单名
	Description     string         `json:"description"`                             // 描述
	CoverImgID      uint64         `json:"coverImgId" gorm:"column:coverImgId"`     // 歌单封面
	CoverImgURL     string         `json:"coverImgUrl" gorm:"column:coverImgUrl"`   // 封面url
	CreateTime      uint64         `json:"createTime" gorm:"column:createTime"`     // 创建时间
	Tags            datatypes.JSON `json:"tags" gorm:"type:json"`                   // 标签
	CommentCount    uint64         `json:"commentCount" gorm:"column:commentCount"` // 评论数
	PlayCount       uint64         `json:"playCount" gorm:"column:playCount"`       // 播放次数
	Status          uint64         `json:"status"`
	SubscribedCount uint64         `json:"subscribedCount" gorm:"column:subscribedCount"` // 订阅次数
	UpdateTime      uint64         `json:"updateTime" gorm:"column:updateTime"`           // 最后更新时间
	UserID          uint64         `json:"userId" gorm:"column:userId"`                   // 创建用户id
	gorm.Model
}

type PlaylistItems struct {
	PlaylistID uint64 `json:"playlistId" gorm:"column:playlistId;primaryKey;autoIncrement:false";` // 歌单id
	SongID     uint64 `json:"songId" gorm:"column:songId;primaryKey;autoIncrement:false"`          // 歌曲id
	UserID     uint64 `json:"userId" gorm:"column:userId;primaryKey;autoIncrement:false"`          // 添加用户id
	Index      uint64 `json:"index"`
	gorm.Model
}
