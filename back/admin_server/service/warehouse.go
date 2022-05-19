package service

import (
	"context"

	"goods/common/constant"
	"goods/common/model"
	"goods/common/util"

	"github.com/jinzhu/gorm"
)

type Warehouse struct{}

func (w *Warehouse) SaveOne(ctx context.Context, name string, authorId int32) error {
	warehouse, err := new(model.Warehouse).GetOneByName(ctx, name)
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return util.New(constant.ErrorDatabase, err)
		}
	}
	if warehouse != nil && warehouse.Id != 0 {
		// 仓库已经存在
		return util.New(constant.ErrorWarehouseRepeat)
	}
	_, err = new(model.Warehouse).SaveOne(ctx, name, authorId)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (w *Warehouse) ModifyOne(ctx context.Context, name string, authorId, id int32) error {
	_, err := new(model.Warehouse).ModifyOne(ctx, name, authorId, id)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (w *Warehouse) GetList(ctx context.Context) (interface{}, error) {
	list, err := new(model.Warehouse).GetList(ctx)
	if err != nil {
		return nil, util.New(constant.ErrorDatabase, err)
	}
	return list, nil
}
