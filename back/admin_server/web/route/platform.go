package route

import (
	"github.com/gin-gonic/gin"
	"goods/web/form"
	"goods/web/handler"
	"goods/web/middleware"
)

type Platform struct{}

func (p *Platform) PlatformRoute(group *gin.RouterGroup) {
	group.GET("/get-list", new(handler.Platform).GetList)                                                 // 获取全部平台
	group.POST("/save", middleware.Bind(form.PlatformSaveOneForm{}), new(handler.Platform).SaveOne)       // 保存一个平台
	group.POST("/modify", middleware.Bind(form.PlatformModifyOneForm{}), new(handler.Platform).ModifyOne) // 修改一个平台
}
