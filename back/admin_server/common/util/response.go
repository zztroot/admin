package util

import (
	"github.com/gin-gonic/gin"
	"goods/common/constant"
	"net/http"
)

// 成功
func Success(g *gin.Context, data interface{}) {
	body := gin.H{
		"code":    constant.ErrorOk,
		"message": new(constant.Error).ErrorMessage(constant.ErrorOk),
		"data":    data,
	}
	g.JSON(http.StatusOK, body)
}

// 失败
func Fail(g *gin.Context, err error) {
	ztErr := WrapBySkip(err, 3)
	body := gin.H{
		"code":    ztErr.Code,
		"message": ztErr.Message,
		"data":    gin.H{},
	}
	g.JSON(http.StatusOK, body)
}
