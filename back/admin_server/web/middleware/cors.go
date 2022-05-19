package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if c.Request.Header.Get("Origin") != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Set("Content-type", "application/json")
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			//c.Header("Access-Control-Allow-Origin", "http://www.zztdd.cn") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, X-TOKEN")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "false")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
