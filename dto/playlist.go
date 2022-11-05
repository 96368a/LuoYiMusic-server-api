package dto

type PlaylistDto struct {
	ID          uint64 `json:"id"`          // 歌单id
	Name        string `json:"name"`        // 歌单名
	Description string `json:"description"` // 描述
	Status      uint64 `json:"status"`
}
