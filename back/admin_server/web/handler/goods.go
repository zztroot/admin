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

type Goods struct{}

func (g *Goods) SaveOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.GoodsSaveOneForm)
	if err := new(service.Goods).SaveOne(ctx, params.WarehouseId, user.Id, params.Name, params.Date, params.Unit, params.CostPrice, params.SupplyPrice, params.SaleNumber, params.ReturnNumber, params.LinkReturnNumber); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (g *Goods) ModifyOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.GoodsModifyOneForm)
	if err := new(service.Goods).ModifyOne(ctx, params.Id, params.WarehouseId, user.Id, params.Name, params.Date, params.Unit, params.CostPrice, params.SupplyPrice, params.SaleNumber, params.ReturnNumber, params.LinkReturnNumber); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (g *Goods) DeleteOne(c *gin.Context) {
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.GoodsDeleteOneForm)
	if err := new(service.Goods).DeleteOne(ctx, params.Id); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (g *Goods) GetList(c *gin.Context) {
	ctx := util.ContextTracer(c)
	// 获取参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	authorId, _ := strconv.Atoi(c.DefaultQuery("author_id", "0"))
	warehouseId, _ := strconv.Atoi(c.DefaultQuery("warehouse_id", "0"))
	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")
	name := c.DefaultQuery("name", "")
	data, err := new(service.Goods).GetList(ctx, int32(page), int32(pageSize), int32(authorId), int32(warehouseId), startDate, endDate, name)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, data)
}
