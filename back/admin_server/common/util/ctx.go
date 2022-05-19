package util

import (
	"context"
	"github.com/gin-gonic/gin"
)

// gin context返回一个context
func ContextTracer(g *gin.Context) context.Context {
	v, exist := g.Get("tracer")
	if !exist {
		return context.TODO()
	}
	return v.(context.Context)
}
