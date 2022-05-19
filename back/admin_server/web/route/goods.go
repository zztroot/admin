package route

import (
	"github.com/gin-gonic/gin"
	"goods/web/form"
	"goods/web/handler"
	"goods/web/middleware"
)

type Goods struct{}

func (g *Goods) GoodsRoute(group *gin.RouterGroup) {
	// /goods
	group.GET("/get-list", new(handler.Goods).GetList)                                              // 获取商品列表
	group.POST("/save", middleware.Bind(form.GoodsSaveOneForm{}), new(handler.Goods).SaveOne)       // 保存商品
	group.POST("/modify", middleware.Bind(form.GoodsModifyOneForm{}), new(handler.Goods).ModifyOne) // 修改商品
	group.POST("/delete", middleware.Bind(form.GoodsDeleteOneForm{}), new(handler.Goods).DeleteOne) // 删除商品
}
