package route

import (
	"github.com/gin-gonic/gin"
	"goods/web/form"
	"goods/web/handler"
	"goods/web/middleware"
)

type Role struct{}

func (r *Role) RoleRoute(group *gin.RouterGroup) {
	// 角色路由
	group.GET("/get-list", new(handler.Role).GetList)                                                // 获取全部
	group.POST("/modify", middleware.Bind(form.ModifyRoleForm{}), new(handler.Role).ModifyOne)       // 修改角色
	group.POST("/del", middleware.Bind(form.DeleteRoleForm{}), new(handler.Role).Delete)             // 删除角色
	group.POST("/get-one", middleware.Bind(form.GetOneByNameForm{}), new(handler.Role).GetOneByName) // 查询一条数据
	group.POST("/save", middleware.Bind(form.GetOneByNameForm{}), new(handler.Role).CreateOne)       // 创建一条数据
}
