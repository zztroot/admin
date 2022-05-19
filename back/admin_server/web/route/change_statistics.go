package route

import (
	"goods/web/form"
	"goods/web/handler"
	"goods/web/middleware"

	"github.com/gin-gonic/gin"
)

type ChangeStatistics struct{}

func (c *ChangeStatistics) ChangeStatisticsRoute(group *gin.RouterGroup) {
	group.GET("/get-list", new(handler.ChangeStatistics).GetList)                                                         // 获取变动成本数据
	group.POST("/save", middleware.Bind(form.SaveChangeStatisticsOneForm{}), new(handler.ChangeStatistics).SaveOne)       // 保存变动成本数据
	group.POST("/modify", middleware.Bind(form.ModifyChangeStatisticsOneForm{}), new(handler.ChangeStatistics).ModifyOne) // 修改变动成本数据
	group.POST("/delete", middleware.Bind(form.DeleteChangeStatisticsOneForm{}), new(handler.ChangeStatistics).DeleteOne) // 删除变动成本数据
}
