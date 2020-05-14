package middlewares

import (
	"github.com/curder/go-gin-demo/commons"
	"github.com/curder/go-gin-demo/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	http "net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			user        models.Users
			DB          *gorm.DB
			userID      uint
			tokenString string
			claims      *commons.Claims
			token       *jwt.Token
			err         error
		)
		// 获取 authorization header
		tokenString = ctx.GetHeader("Authorization")

		// 验证 tokenString
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:] // 获取传入的 token 有效部分，除去"Bearer "后的字符

		if token, claims, err = commons.ParseToken(tokenString); err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "权限不足"})
			ctx.Abort()
			return
		}

		// token中的UserID
		userID = claims.UserID
		DB = commons.GetDB()

		DB.First(&user, userID)
		// 用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "权限不足"})
			ctx.Abort()
			return
		}

		ctx.Set("user", user) // 写入gin上下文

		ctx.Next()
	}
}
