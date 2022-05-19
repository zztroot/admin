package handler

import (
	"goods/common/constant"
	"goods/common/util"
	"goods/service"
	"goods/web/form"
	"goods/web/middleware"

	"github.com/gin-gonic/gin"
)

type Platform struct{}

func (p *Platform) SaveOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.PlatformSaveOneForm)
	if err := new(service.Platform).SaveOne(ctx, user.Id, params.Name); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (p *Platform) ModifyOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.PlatformModifyOneForm)
	if err := new(service.Platform).ModifyOne(ctx, user.Id, params.Id, params.Name); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (p *Platform) GetList(c *gin.Context) {
	ctx := util.ContextTracer(c)
	data, err := new(service.Platform).GetList(ctx)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, data)
}
