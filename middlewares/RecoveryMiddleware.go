package middlewares

import (
	"fmt"
	"github.com/curder/go-gin-demo/responses"
	"github.com/gin-gonic/gin"
)

// 处理panic中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				responses.Fail(ctx, fmt.Sprint(err), nil)
			}
		}()
		ctx.Next()
	}
}
