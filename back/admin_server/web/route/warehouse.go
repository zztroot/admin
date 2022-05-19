package route

import (
	"github.com/gin-gonic/gin"
	"goods/web/form"
	"goods/web/middleware"

	"goods/web/handler"
)

type Warehouse struct{}

func (w *Warehouse) WarehouseRoute(group *gin.RouterGroup) {
	group.GET("/get-list", new(handler.Warehouse).GetList)                                                  // 获取全部仓库
	group.POST("/save", middleware.Bind(form.WarehouseSaveOneForm{}), new(handler.Warehouse).SaveOne)       // 保存一个仓库
	group.POST("/modify", middleware.Bind(form.WarehouseModifyOneForm{}), new(handler.Warehouse).ModifyOne) // 修改一个仓库
}
