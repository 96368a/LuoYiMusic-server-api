package dto

type SongDto struct {
	ID    uint64 `json:"id"`    // 歌曲id
	Name  string `json:"name"`  // 歌曲名
	Album uint64 `json:"album"` // 所属专辑id
}
