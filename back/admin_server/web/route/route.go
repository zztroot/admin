package route

import (
	"goods/common/constant"
	"goods/common/util"
	"goods/web/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(engine *gin.Engine) {
	group := engine.Group("/goods-admin")
	// 用户路由
	new(User).UserRoute(group.Group("/user"))
	// 以下路由需要检查token
	group.Use(middleware.CheckToken())
	{
		// 商品路由
		new(Goods).GoodsRoute(group.Group("/goods"))
		// 角色路由
		new(Role).RoleRoute(group.Group("/role"))
		// 仓库路由
		new(Warehouse).WarehouseRoute(group.Group("/warehouse"))
		// 平台统计路由
		new(PlatformStatistics).PlatformStatisticsRoute(group.Group("/platform-statistics"))
		// 平台路由
		new(Platform).PlatformRoute(group.Group("/platform"))
	}
	// 获取全部菜单
	group.GET("/menu", func(c *gin.Context) {
		util.Success(c, constant.ListWebMenu)
	})
	// 测试接口
	group.GET("/test", func(c *gin.Context) {
		util.Success(c, gin.H{})
	})
}
