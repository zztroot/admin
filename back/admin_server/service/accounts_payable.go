package service

import (
	"context"
	"goods/common/constant"
	"goods/common/model"
	"goods/common/util"
)

type AccountsPayable struct{}

func (a *AccountsPayable) SaveOne(ctx context.Context, date, unit, comments string, number int, goodsId, authorId int32, price float32) error {
	_, err := new(model.AccountsPayable).SaveOne(ctx, date, unit, comments, number, goodsId, authorId, price)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (a *AccountsPayable) ModifyOne(ctx context.Context, date, unit, comments string, number int, goodsId, authorId, id int32, price float32) error {
	_, err := new(model.AccountsPayable).ModifyOne(ctx, date, unit, comments, number, goodsId, authorId, id, price)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (a *AccountsPayable) GetList(ctx context.Context, page, pageSize, authorId int32, startDate, endDate, goodsName string) (interface{}, error) {
	list, err := new(model.AccountsPayable).GetList(ctx, page, pageSize, authorId, startDate, endDate, goodsName)
	if err != nil {
		return nil, util.New(constant.ErrorDatabase, err)
	}
	return list, nil
}

func (a *AccountsPayable) DeleteOne(ctx context.Context, id int32) error {
	if err := new(model.AccountsPayable).DeleteById(ctx, id); err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}
