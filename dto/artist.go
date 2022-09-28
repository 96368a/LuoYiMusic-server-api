package dto

type ArtistDto struct {
	ID          uint64   `json:"id"`          // id
	Name        string   `json:"name"`        // 歌手名字
	Description string   `json:"description"` // 歌手描述
	Alias       []string `json:"alias"`       // 歌手别名
	PicID       uint64   `json:"picId"`       // 歌手图片id
	PicURL      string   `json:"picUrl"`      // 歌手图片url
}
