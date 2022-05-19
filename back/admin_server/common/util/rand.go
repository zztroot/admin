package util

import (
	"encoding/base64"
	"math/rand"
)

const sourceStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	// 初始化随机种子
	rand.Seed(NowTime().UnixNano())
}

// 获取随机字符串
func RandomString(length int) string {
	bytes := []byte(sourceStr)
	var result []byte
	for i := 0; i < length; i++ {
		key := RandomInt(len(bytes))
		result = append(result, bytes[key])
	}
	return base64.URLEncoding.EncodeToString(result)[:length]
}

// 获取随机整数
func RandomInt(n int) int {
	return rand.Intn(n)
}
