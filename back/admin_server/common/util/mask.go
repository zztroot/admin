package util

import (
	"encoding/base64"

	"github.com/zztroot/zztlog"

	"golang.org/x/crypto/bcrypt"
)

// 加密
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		zztlog.Error(err)
	}
	return string(hash)
}

// 密码对比
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		zztlog.Debug(err)
		return false
	}
	return true
}

var prefix = "$#6-6#$"
var suffix = "$#520#$"

// 加密
func EncodeStr(str string) string {
	tempStr := prefix + str + suffix
	return base64.StdEncoding.EncodeToString([]byte(tempStr))
}

// 解密
func DecodeStr(str string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		zztlog.Error(err)
		return ""
	}
	return string(decodeBytes)
}
