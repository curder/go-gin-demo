package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")                     // 允许跨域的域名
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")                      // 缓存时间
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATH,DELETE") // 允许请求的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")                    // 请求的头
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")             // 是否允许https

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
		} else {
			ctx.Next()
		}
	}
}
