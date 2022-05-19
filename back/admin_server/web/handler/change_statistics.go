package handler

import (
	"goods/common/constant"
	"goods/common/util"
	"goods/service"
	"goods/web/form"
	"goods/web/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChangeStatistics struct{}

func (s *ChangeStatistics) SaveOne(c *gin.Context) {
	user := middleware.GetUser(c)
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.SaveChangeStatisticsOneForm)
	if err := new(service.ChangeStatistics).SaveOne(ctx, params, user.Id); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (s *ChangeStatistics) ModifyOne(c *gin.Context) {
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.ModifyChangeStatisticsOneForm)
	if err := new(service.ChangeStatistics).ModifyOne(ctx, params); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (s *ChangeStatistics) DeleteOne(c *gin.Context) {
	ctx := util.ContextTracer(c)
	params := c.MustGet(constant.GinBindFormKey).(*form.DeleteChangeStatisticsOneForm)
	if err := new(service.ChangeStatistics).DeleteOne(ctx, params.Id); err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, gin.H{})
}

func (s *ChangeStatistics) GetList(c *gin.Context) {
	ctx := util.ContextTracer(c)
	// 获取参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	authorId, _ := strconv.Atoi(c.DefaultQuery("author_id", "0"))
	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")
	data, err := new(service.ChangeStatistics).GetList(ctx, int32(page), int32(pageSize), int32(authorId), startDate, endDate)
	if err != nil {
		util.Fail(c, err)
		return
	}
	util.Success(c, data)
}
