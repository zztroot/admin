package middleware

import (
	"io"
	"reflect"

	"goods/common/constant"
	"goods/common/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Bind(val interface{}) gin.HandlerFunc {
	return func(g *gin.Context) {
		value := reflect.ValueOf(val)
		if value.Kind() == reflect.Ptr {
			panic(`Bind bean can not be a pointer. Example: Use: gin.Bind(Struct{}) instead of gin.Bind(&Struct{})`)
		}
		obj := reflect.New(value.Type()).Interface()
		err := g.ShouldBindBodyWith(obj, binding.JSON)
		if err != nil && err != io.EOF {
			util.Fail(g, util.New(constant.ErrorParams, err))
			g.Abort()
			return
		}
		g.Set(constant.GinBindFormKey, obj)
		g.Next()
	}
}
