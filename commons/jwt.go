package commons

import (
	"github.com/curder/go-gin-demo/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte(`RAQFEAirRV9k7IxJ0QjkkgJY`)

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// 发放token
func ReleaseToken(user models.Users) (tokenString string, err error) {
	var (
		expirationTime time.Time
		claims         *Claims
		token          *jwt.Token
	)

	expirationTime = time.Now().Add(7 * 24 * time.Hour) // token有效时间：7 天

	claims = &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    `go-gin-demo`,
			Subject:   `user token`,
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(jwtKey)

	return
}

// 解析token
func ParseToken(tokenString string) (token *jwt.Token, claims *Claims, err error) {
	claims = &Claims{}

	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return
}
