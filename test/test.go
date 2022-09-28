package test

import (
	"encoding/json"
	"fmt"
	"github.com/96368a/LuoYiMusic-server-api/dto"
	"github.com/96368a/LuoYiMusic-server-api/model"
	"github.com/dhowden/tag"
	"gorm.io/datatypes"
	"log"
	"os"
)

func Test() {
	fmt.Printf("test\n")
	//testSongId3()
	//services.AddArtist("test")
	//ok, _ := services.CheckArtist("ttest")
	//fmt.Printf("%v\n", ok)

}

func testSongId3() {
	f, _ := os.Open("resources/musics/1.mp3")
	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	picture := m.Picture()
	os.WriteFile("resources/musics/1.png", picture.Data, 0644)
	log.Print(picture.String())
	log.Print(m.Format()) // The detected format.
	log.Print(m.Title())  // The title of the track (see Metadata interface for more details).
}

func TestDBJson() {
	alas := datatypes.JSON([]byte(`["233","dddd"]`))
	model.DB.Create(&model.Artist{
		Alias:       alas,
		Description: "23333",
		Name:        "test",
		PicID:       0,
		PicURL:      "dddd",
	})
	var artist []model.Artist
	model.DB.Find(&artist, datatypes.JSONQuery("Alias").Equals("", `dddd`))
	model.DB.Raw("SELECT * FROM artists,json_each(artists.alias) where json_each.value = ?", "dddd").Scan(&artist)
	fmt.Printf("%v\n", artist)
}

func TestArtist() {
	artistDto := dto.ArtistDto{
		ID:          0,
		Name:        "",
		Description: "",
		Alias:       nil,
		PicID:       0,
		PicURL:      "",
	}
	artist, _ := json.Marshal(artistDto)
	fmt.Printf("%v\n", string(artist))
}