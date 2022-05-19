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

type PlatformStatistics struct{}

func (p *PlatformStatistics) SaveOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.PlatformStatisticsSaveOneForm)
	if err := new(service.PlatformStatistics).SaveOne(ctx, params.PlatformId, user.Id, params.Price, params.Goods, params.Comments); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (p *PlatformStatistics) ModifyOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.PlatformStatisticsModifyOneForm)
	if err := new(service.PlatformStatistics).ModifyOne(ctx, params.Id, params.PlatformId, user.Id, params.SalesTimes, params.Price, params.Goods, params.Comments); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (p *PlatformStatistics) DeleteOne(c *gin.Context) {
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.PlatformStatisticsDeleteOneForm)
	if err := new(service.PlatformStatistics).DeleteOne(ctx, params.Id); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (p *PlatformStatistics) GetList(c *gin.Context) {
	ctx := util.ContextTracer(c)
	// 获取参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	authorId, _ := strconv.Atoi(c.DefaultQuery("author_id", "0"))
	platformId, _ := strconv.Atoi(c.DefaultQuery("platform_id", "0"))
	goodsName := c.DefaultQuery("goods_name", "")
	data, err := new(service.PlatformStatistics).GetList(ctx, int32(page), int32(pageSize), int32(authorId), int32(platformId), goodsName)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, data)
}
