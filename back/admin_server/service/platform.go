package service

import (
	"context"

	"goods/common/constant"
	"goods/common/model"
	"goods/common/util"
)

type Platform struct{}

func (p *Platform) SaveOne(ctx context.Context, authorId int32, name string) error {
	_, err := new(model.Platform).SaveOne(ctx, authorId, name)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (p *Platform) ModifyOne(ctx context.Context, id, authorId int32, name string) error {
	_, err := new(model.Platform).ModifyOne(ctx, id, authorId, name)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (p *Platform) GetList(ctx context.Context) (interface{}, error) {
	list, err := new(model.Platform).GetList(ctx)
	if err != nil {
		return nil, util.New(constant.ErrorDatabase, err)
	}
	return list, nil
}
