package router

import (
	"github.com/96368a/LuoYiMusic-server-api/controller"
	"github.com/96368a/LuoYiMusic-server-api/controller/api"
	"github.com/96368a/LuoYiMusic-server-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//跨域处理
	r.Use(middleware.CORSMiddleware())

	r.GET("/song", controller.SongById)

	recommendGroup := r.Group("/recommend")
	recommendGroup.GET("/songs", controller.RecommendSongs)
	recommendGroup.GET("/playlists", controller.RecommendPlaylists)

	songGroup := r.Group("/song")
	songGroup.GET("/search", controller.SearchSongs)
	songGroup.GET("/news", controller.SongNews)

	albumGroup := r.Group("/album")
	albumGroup.GET("/search", controller.SearchAlbums)
	albumGroup.GET("/news", controller.AlbumNews)

	artistGroup := r.Group("/artist")
	artistGroup.GET("/search", controller.SearchArtist)

	playListGroup := r.Group("/playlist")
	playListGroup.GET("/search", controller.SearchPlaylist)
	playListGroup.GET("/self", middleware.AuthMiddleware(), controller.SelfPlaylist)
	playListGroup.POST("/add", middleware.AuthMiddleware(), controller.AddPlaylist)
	playListGroup.POST("/del", middleware.AuthMiddleware(), controller.DelPlaylist)
	playListGroup.POST("/addSong", middleware.AuthMiddleware(), controller.AddSongPlaylist)

	userGroup := r.Group("/user")
	userGroup.POST("/register", controller.Register)
	userGroup.POST("/login", controller.Login)
	userGroup.POST("/update", middleware.AuthMiddleware(), controller.UpdateUser)
	userGroup.GET("/info", middleware.AuthMiddleware(), controller.UserInfo)
	userGroup.POST("/changePassword", middleware.AuthMiddleware(), controller.ChangePassword)

	apiGroup := r.Group("/api", middleware.AuthMiddleware(), middleware.AdminAuthMiddleware())
	apiGroup.GET("/users", api.GetAllUsers)
	apiGroup.GET("/user/search", api.SearchUsers)
	apiGroup.POST("/user/add", api.AddUser)
	apiGroup.POST("/user/del", api.DelUser)
	apiGroup.POST("/user/update", api.UpdateUser)
	apiGroup.POST("/user/changePassword", api.ChangePassword)
	apiGroup.POST("/user/setAdmin", api.SetAdmin)
	apiGroup.POST("/user/removeAdmin", api.RemoveAdmin)

	apiGroup.POST("/song/add", api.AddArtist)
	apiGroup.POST("/song/uploads", api.SongUploads)
	apiGroup.POST("/song/del", api.DelSong)

	apiGroup.POST("/artist/del", api.DelArtist)
	apiGroup.POST("/album/del", api.DelAlbum)

	apiGroup.POST("/playlist/add", api.AddPlaylist)
	apiGroup.POST("/playlist/del", api.DelPlaylist)
	apiGroup.GET("/playlist/search", api.SearchPlaylist)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "页面不存在",
		})
	})

	return r
}
