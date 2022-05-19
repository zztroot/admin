package handler

import (
	"goods/common/constant"
	"goods/common/util"
	"goods/service"
	"goods/web/form"
	"goods/web/middleware"

	"github.com/gin-gonic/gin"
)

type Role struct{}

func (r *Role) GetList(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	if user.UserRoleId != 0 {
		util.Fail(c, util.New(constant.ErrorNoPermission))
		return
	}
	data, err := new(service.Role).GetList(ctx)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, data)
}

func (r *Role) ModifyOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	if user.UserRoleId != 0 {
		util.Fail(c, util.New(constant.ErrorNoPermission))
		return
	}
	params := c.MustGet(constant.GinBindFormKey).(*form.ModifyRoleForm)
	if err := new(service.Role).ModifyOne(ctx, params.RoleName, params.RoleExtend, params.Id); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (r *Role) Delete(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	if user.UserRoleId != 0 {
		util.Fail(c, util.New(constant.ErrorNoPermission))
		return
	}
	params := c.MustGet(constant.GinBindFormKey).(*form.DeleteRoleForm)
	if err := new(service.Role).Delete(ctx, params.Id); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (r *Role) GetOneByName(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	if user.UserRoleId != 0 {
		util.Fail(c, util.New(constant.ErrorNoPermission))
		return
	}
	params := c.MustGet(constant.GinBindFormKey).(*form.GetOneByNameForm)
	data, err := new(service.Role).GetOneByName(ctx, params.Name)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, data)
}

func (r *Role) CreateOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	if user.UserRoleId != 0 {
		util.Fail(c, util.New(constant.ErrorNoPermission))
		return
	}
	params := c.MustGet(constant.GinBindFormKey).(*form.GetOneByNameForm)
	if err := new(service.Role).CreateOne(ctx, params.Name); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}
