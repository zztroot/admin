package util

import "goods/common/constant"

func GenerateGoodsCode() string {
	return NowTime().Format(constant.DateLayoutStr) + RandomString(3) + NowTime().Format("0405")
}
