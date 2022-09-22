package main

import (
	"github.com/96368a/LuoYiMusic-server-api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := router.InitRouter()
	r.SetTrustedProxies(nil)
	//alas := datatypes.JSON([]byte(`["233","dddd"]`))
	//model.DB.Create(&song.Artist{
	//	Alias:       alas,
	//	Description: "23333",
	//	Name:        "test",
	//	PicID:       0,
	//	PicURL:      "dddd",
	//})
	//var artist []song.Artist
	//model.DB.Find(&artist, datatypes.JSONQuery("Alias").Equals("", `dddd`))
	//model.DB.Raw("SELECT * FROM artists,json_each(artists.alias) where json_each.value = ?", "dddd").Scan(&artist)
	//fmt.Printf("%v\n", artist)
	panic(r.Run(":8888"))
}
