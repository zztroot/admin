package handler

import (
	"strconv"

	"goods/common/constant"
	"goods/common/util"
	"goods/service"
	"goods/web/form"
	"goods/web/middleware"

	"github.com/gin-gonic/gin"
)

type User struct{}

// 创建用户
func (u *User) SaveUser(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	if user.UserRoleId != 0 {
		util.Fail(c, util.New(constant.ErrorNoPermission))
		return
	}
	params := c.MustGet(constant.GinBindFormKey).(*form.SaveUserForm)
	if err := new(service.User).SaveOne(ctx, params); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

// 用户登录
func (u *User) Login(c *gin.Context) {
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.LoginForm)
	data, err := new(service.User).Login(ctx, params.UserName, params.UserPassword)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, data)
}

// 用户修改
func (u *User) ModifyOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.ModifyUserForm)
	if err := new(service.User).ModifyOne(ctx, params, user); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

// 获取所有用户
func (u *User) GetUserList(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	if user.UserRoleId != 0 {
		util.Fail(c, util.New(constant.ErrorNoPermission))
		return
	}
	// 获取参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	userName := c.DefaultQuery("user_name", "")
	sex, _ := strconv.Atoi(c.DefaultQuery("sex", "0"))
	data, err := new(service.User).GetUserList(ctx, int32(page), int32(pageSize), userName, sex)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, data)
}

// 根据ID获取单个用户
func (u *User) GetUserOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	if user.UserRoleId != 0 {
		util.Fail(c, util.New(constant.ErrorNoPermission))
		return
	}
	params := c.MustGet(constant.GinBindFormKey).(*form.GetUserOneForm)
	data, err := new(service.User).GetUserOne(ctx, params.Id)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, data)
}

// 根据ID删除单个用户
func (u *User) DeleteUserOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	if user.UserRoleId != 0 {
		util.Fail(c, util.New(constant.ErrorNoPermission))
		return
	}
	params := c.MustGet(constant.GinBindFormKey).(*form.GetUserOneForm)
	err := new(service.User).DeleteUserOne(ctx, params.Id)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (u *User) GetUserAll(c *gin.Context) {
	ctx := util.ContextTracer(c)
	data, err := new(service.User).GetUserAll(ctx)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, data)
}
