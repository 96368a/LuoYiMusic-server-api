package vo

import (
	"github.com/96368a/LuoYiMusic-server-api/model"
	"github.com/96368a/LuoYiMusic-server-api/services"
	"gorm.io/datatypes"
)

type PlaylistVo struct {
	ID              uint64         `json:"id"`           // 歌单id
	Name            string         `json:"name"`         // 歌单名
	Description     string         `json:"description"`  // 描述
	CoverImgID      uint64         `json:"coverImgId"`   // 歌单封面
	CoverImgURL     string         `json:"coverImgUrl"`  // 封面url
	CreateTime      uint64         `json:"createTime"`   // 创建时间
	Tags            datatypes.JSON `json:"tags"`         // 标签
	CommentCount    uint64         `json:"commentCount"` // 评论数
	PlayCount       uint64         `json:"playCount"`    // 播放次数
	Status          uint64         `json:"status"`
	SubscribedCount uint64         `json:"subscribedCount"` // 订阅次数
	UpdateTime      uint64         `json:"updateTime"`      // 最后更新时间
	UserID          uint64         `json:"userId"`
	Songs           []SongInfoVo   `json:"songs"`
}

func ToPlaylistVo(playlist model.Playlist) PlaylistVo {
	songs, err := services.PlaylistSongs(playlist.ID)
	if err != nil {
		return PlaylistVo{}
	}
	return PlaylistVo{
		ID:          playlist.ID,
		Name:        playlist.Name,
		Description: playlist.Description,
		Status:      playlist.Status,
		UserID:      playlist.UserID,
		Songs:       ToSongInfoVos(songs),
	}
}
func ToPlaylistVos(playlists []model.Playlist) []PlaylistVo {
	playlistVos := make([]PlaylistVo, len(playlists))
	for i, playlist := range playlists {
		playlistVos[i] = ToPlaylistVo(playlist)
	}
	return playlistVos
}
