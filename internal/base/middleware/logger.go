package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-sven/layout/internal/base/constant"
	"io"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.GetHeader(constant.H_REQUEST_ID)
		if requestId == "" {
			requestId = ""
		}
		ctx.Set(constant.H_REQUEST_ID, requestId)
		bodyBytes, err := ctx.GetRawData()
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "无效的请求体"})
			return
		}
		ctx.Set(constant.RAW_BODY, bodyBytes)
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		ctx.Next()

	}
}
