package route

import (
	"goods/web/form"
	"goods/web/handler"
	"goods/web/middleware"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (u *User) UserRoute(group *gin.RouterGroup) {
	// 路由/user
	group.POST("/login", middleware.Bind(form.LoginForm{}), new(handler.User).Login) // 登录

	group.POST("/register", middleware.CheckToken(), middleware.Bind(form.SaveUserForm{}), new(handler.User).SaveUser)          // 注册
	group.POST("/modify", middleware.CheckToken(), middleware.Bind(form.ModifyUserForm{}), new(handler.User).ModifyOne)         // 修改
	group.GET("/get-list", middleware.CheckToken(), new(handler.User).GetUserList)                                              // 获取用户列表
	group.POST("/get-one", middleware.CheckToken(), middleware.Bind(form.GetUserOneForm{}), new(handler.User).GetUserOne)       // 获取单个用户
	group.POST("/delete-one", middleware.CheckToken(), middleware.Bind(form.GetUserOneForm{}), new(handler.User).DeleteUserOne) // 删除单个用户
	group.GET("/get-all", middleware.CheckToken(), new(handler.User).GetUserAll)
}
