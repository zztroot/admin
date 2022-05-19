package handler

import (
	"goods/common/constant"
	"goods/common/util"
	"goods/service"
	"goods/web/form"
	"goods/web/middleware"

	"github.com/gin-gonic/gin"
)

type Warehouse struct{}

func (w *Warehouse) SaveOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.WarehouseSaveOneForm)
	if err := new(service.Warehouse).SaveOne(ctx, params.Name, user.Id); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (w *Warehouse) ModifyOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.WarehouseModifyOneForm)
	if err := new(service.Warehouse).ModifyOne(ctx, params.Name, user.Id, params.Id); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (w *Warehouse) GetList(c *gin.Context) {
	ctx := util.ContextTracer(c)
	data, err := new(service.Warehouse).GetList(ctx)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, data)
}
