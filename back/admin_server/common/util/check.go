package util

import (
	"regexp"
	"unicode"

	"goods/common/constant"
)

// 校验手机号
func CheckPhone(phone string) bool {
	ok, _ := regexp.MatchString("^1[0-9]{10}$", phone)
	return ok
}

// 校验账号和密码
func CheckUserNameAndPwd(userName, pwd string) error {
	if IsChineseChar(userName) {
		return New(constant.ErrorUserNameContainChinese)
	}
	if IsChineseChar(pwd) {
		return New(constant.ErrorPwdContainChinese)
	}
	// 校验账号长度
	if len(userName) < 5 || len(userName) > 15 {
		return New(constant.ErrorUserNameOutOfRange)
	}
	// 校验密码长度
	if len(pwd) < 5 {
		return New(constant.ErrorPwdOutOfRange)
	}
	return nil
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
