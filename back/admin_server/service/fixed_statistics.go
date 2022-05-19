package service

import (
	"context"

	"goods/common/constant"
	"goods/common/model"
	"goods/common/util"
)

type FixedStatistics struct{}

func (f *FixedStatistics) SaveOne(ctx context.Context, date string, wages, rent, socialSecurity, pcCost, otherCost float32, authorId int32) error {
	_, err := new(model.FixedStatistics).SaveOne(ctx, date, wages, rent, socialSecurity, pcCost, otherCost, authorId)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (f *FixedStatistics) ModifyOne(ctx context.Context, date string, wages, rent, socialSecurity, pcCost, otherCost float32, authorId, id int32) error {
	_, err := new(model.FixedStatistics).ModifyOne(ctx, date, wages, rent, socialSecurity, pcCost, otherCost, authorId, id)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (f *FixedStatistics) GetList(ctx context.Context, page, pageSize, authorId int32, startDate, endDate string) (interface{}, error) {
	list, err := new(model.FixedStatistics).GetList(ctx, page, pageSize, authorId, startDate, endDate)
	if err != nil {
		return nil, util.New(constant.ErrorDatabase, err)
	}
	return list, nil
}

func (f *FixedStatistics) DeleteOne(ctx context.Context, id int32) error {
	if err := new(model.FixedStatistics).DeleteById(ctx, id); err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}
