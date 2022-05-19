package route

import (
	"goods/web/form"
	"goods/web/handler"
	"goods/web/middleware"

	"github.com/gin-gonic/gin"
)

type PlatformStatistics struct{}

func (g *PlatformStatistics) PlatformStatisticsRoute(group *gin.RouterGroup) {
	group.GET("/get-list", new(handler.PlatformStatistics).GetList)                                                           // 获取平台数据
	group.POST("/save", middleware.Bind(form.PlatformStatisticsSaveOneForm{}), new(handler.PlatformStatistics).SaveOne)       // 保存平台数据
	group.POST("/modify", middleware.Bind(form.PlatformStatisticsModifyOneForm{}), new(handler.PlatformStatistics).ModifyOne) // 修改及平台数据
	group.POST("/delete", middleware.Bind(form.PlatformStatisticsDeleteOneForm{}), new(handler.PlatformStatistics).DeleteOne) // 删除平台数据
}
