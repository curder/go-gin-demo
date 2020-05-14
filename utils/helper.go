package utils

import (
	"math/rand"
	"time"
)

// 生成随机字符串
func RandomString(n int) string {
	var (
		letters []byte
		index   int
		result  []byte
	)

	letters = []byte("abcdefhijklmnopqrstuvwxyzABCDEFHIJKLMNOPQRSTUVWXYZ1234597890")
	result = make([]byte, n)
	rand.Seed(time.Now().Unix())

	for index = range result {
		result[index] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
