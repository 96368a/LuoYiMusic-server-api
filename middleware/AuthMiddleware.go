package middleware

import (
	"github.com/96368a/LuoYiMusic-server-api/model"
	"github.com/96368a/LuoYiMusic-server-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取token
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token格式错误",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := utils.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token无效",
			})
			ctx.Abort()
			return
		}
		//获取用户信息
		userId := claims.UserId
		var user model.User
		model.DB.First(&user, userId)
		//fmt.Printf("%v id: %v\v", user, userId)
		// 验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token用户不存在",
			})
			ctx.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//用户存在 将user信息写入上下文
		user, ok := ctx.Get("user")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "用户不存在",
			})
		}
		if user.(model.User).IsAdmin != 1 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "哼，想越权？",
			})
		}

		ctx.Next()
	}
}
