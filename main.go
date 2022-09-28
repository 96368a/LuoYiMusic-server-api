package main

import (
	"github.com/96368a/LuoYiMusic-server-api/router"
	"github.com/96368a/LuoYiMusic-server-api/test"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := router.InitRouter()
	r.SetTrustedProxies(nil)

	test.Test()
	panic(r.Run(":8888"))
}
