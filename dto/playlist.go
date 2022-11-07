package dto

type PlaylistDto struct {
	ID          uint64 `json:"id"`          // 歌单id
	Name        string `json:"name"`        // 歌单名
	Description string `json:"description"` // 描述
	Status      uint64 `json:"status"`
}

type PlaylistItemDto struct {
	PlaylistID uint64   `json:"playlistId"` // 歌单id
	SongID     uint64   `json:"songId"`     // 歌曲id
	SongIDs    []uint64 `json:"songIds"`    // 歌曲ids
	UserID     uint64   `json:"userId"`     // 添加用户id
	Index      uint64   `json:"index"`
}
