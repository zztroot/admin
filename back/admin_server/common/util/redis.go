package util

import (
	"bytes"
	"strconv"
)

func BuildRedisKey(str ...interface{}) string {
	symbol := ":"
	if len(str) < 1 {
		return ""
	}
	build := bytes.Buffer{}
	for i, v := range str {
		if i == 0 {
			// 第一个str必须是字符串
			build.WriteString(v.(string))
			continue
		}
		build.WriteString(symbol)
		switch v := v.(type) {
		case string:
			build.WriteString(v)
		case int:
			build.WriteString(strconv.Itoa(v))
		case int32:
			build.WriteString(strconv.Itoa(int(v)))
		}
	}
	return build.String()
}
