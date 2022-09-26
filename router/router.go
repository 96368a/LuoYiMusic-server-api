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

	userGroup := r.Group("/user")
	userGroup.POST("/register", controller.Register)
	userGroup.POST("/login", controller.Login)
	userGroup.POST("/update", middleware.AuthMiddleware(), controller.UpdateUser)
	userGroup.GET("/info", middleware.AuthMiddleware(), controller.UserInfo)
	userGroup.POST("/changePassword", middleware.AuthMiddleware(), controller.ChangePassword)

	apiGroup := r.Group("/api", middleware.AuthMiddleware(), middleware.AdminAuthMiddleware())
	apiGroup.GET("/users", api.GetAllUsers)
	apiGroup.POST("/user/add", api.AddUser)
	apiGroup.POST("/user/update", api.UpdateUser)
	apiGroup.POST("/user/changePassword", api.ChangePassword)
	apiGroup.POST("/user/setAdmin", api.SetAdmin)
	apiGroup.POST("/user/removeAdmin", api.RemoveAdmin)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "页面不存在",
		})
	})

	return r
}
