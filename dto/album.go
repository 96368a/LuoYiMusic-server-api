package dto

type AlbumDto struct {
	ID          uint64 `json:"id"`          // 专辑id
	Name        string `json:"name"`        // 专辑名
	Description string `json:"description"` // 专辑描述
	ArtistID    uint64 `json:"artistId"`    // 所属歌手
	PicID       uint64 `json:"picId"`       // 专辑id
	PicURL      string `json:"picUrl"`      // 专辑url
}
