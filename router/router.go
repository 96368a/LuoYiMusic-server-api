package router

import (
	"github.com/96368a/LuoYiMusic-server-api/controller"
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

	return r
}
