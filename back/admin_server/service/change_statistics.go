package service

import (
	"context"

	"goods/common/constant"
	"goods/common/model"
	"goods/common/util"
	"goods/web/form"
)

type ChangeStatistics struct{}

func (c *ChangeStatistics) SaveOne(ctx context.Context, params *form.SaveChangeStatisticsOneForm, authorId int32) error {
	changeStatistics := model.ChangeStatistics{
		Date:                       params.Date,
		WarehouseCost:              params.WarehouseCost,
		HandWarehouseGoodsCost:     params.HandWarehouseGoodsCost,
		BackWarehouseGoodsCost:     params.BackWarehouseGoodsCost,
		ChangeWarehouseCost:        params.ChangeWarehouseCost,
		PackCost:                   params.PackCost,
		CoordinationWarehouseCost:  params.CoordinationWarehouseCost,
		OperationAndUnloadingCost:  params.OperationAndUnloadingCost,
		ReplenishmentCost:          params.ReplenishmentCost,
		ManyBackWarehouseGoodsCost: params.ManyBackWarehouseGoodsCost,
		OperationBackGoodsCost:     params.OperationBackGoodsCost,
		ChangeCost:                 params.ChangeCost,
		UserBackGoodsCost:          params.UserBackGoodsCost,
		DeliveryCost:               params.DeliveryCost,
		LogisticsCost:              params.LogisticsCost,
		GoodsLaLaCost:              params.GoodsLaLaCost,
		OtherCost:                  params.OtherCost,
	}
	_, err := new(model.ChangeStatistics).SaveOne(ctx, &changeStatistics)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (c *ChangeStatistics) ModifyOne(ctx context.Context, params *form.ModifyChangeStatisticsOneForm) error {
	changeStatistics := model.ChangeStatistics{
		Base:                       model.Base{Id: params.Id},
		Date:                       params.Date,
		WarehouseCost:              params.WarehouseCost,
		HandWarehouseGoodsCost:     params.HandWarehouseGoodsCost,
		BackWarehouseGoodsCost:     params.BackWarehouseGoodsCost,
		ChangeWarehouseCost:        params.ChangeWarehouseCost,
		PackCost:                   params.PackCost,
		CoordinationWarehouseCost:  params.CoordinationWarehouseCost,
		OperationAndUnloadingCost:  params.OperationAndUnloadingCost,
		ReplenishmentCost:          params.ReplenishmentCost,
		ManyBackWarehouseGoodsCost: params.ManyBackWarehouseGoodsCost,
		OperationBackGoodsCost:     params.OperationBackGoodsCost,
		ChangeCost:                 params.ChangeCost,
		UserBackGoodsCost:          params.UserBackGoodsCost,
		DeliveryCost:               params.DeliveryCost,
		LogisticsCost:              params.LogisticsCost,
		GoodsLaLaCost:              params.GoodsLaLaCost,
		OtherCost:                  params.OtherCost,
	}
	_, err := new(model.ChangeStatistics).ModifyOne(ctx, &changeStatistics)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (c *ChangeStatistics) GetList(ctx context.Context, page, pageSize, authorId int32, startDate, endDate string) (interface{}, error) {
	list, err := new(model.ChangeStatistics).GetList(ctx, page, pageSize, authorId, startDate, endDate)
	if err != nil {
		return nil, util.New(constant.ErrorDatabase, err)
	}
	return list, nil
}

func (c *ChangeStatistics) DeleteOne(ctx context.Context, id int32) error {
	if err := new(model.ChangeStatistics).DeleteById(ctx, id); err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}
