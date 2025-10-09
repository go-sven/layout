package middleware

import (
	"github.com/gin-gonic/gin"
)

func SignMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.GetRawData()
		/*bodyBytes,exists := ctx.Get("rwaBody")
		if !exists {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "中间件配置错误"})
			return
		}


		appId := ctx.GetHeader("X-App-id")
		appSecret := getAppSecret(appId)

		params := ctx.

		//signStr := sign.Generate()//*/

	}
}

func getAppSecret(appId string) string {
	return appId
}
