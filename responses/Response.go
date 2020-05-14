package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 成功响应
func Success(ctx *gin.Context, data gin.H, message string) {
	Response(ctx, http.StatusOK, http.StatusOK, data, message)
}

// 失败响应
func Fail(ctx *gin.Context, message string, data gin.H) {
	Response(ctx, http.StatusBadRequest, http.StatusBadRequest, data, message)
}

// 统一响应处理
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, message string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "message": message})
}
