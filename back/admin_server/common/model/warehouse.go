package model

import (
	"context"

	"goods/common/constant"
	"goods/common/util"

	"github.com/jinzhu/gorm"
)

type Warehouse struct {
	Base
	MysqlConn
	WarehouseName     string `json:"warehouse_name"`      // 仓库名称
	WarehouseAuthorId int32  `json:"warehouse_author_id"` // 创建者
}

func (w *Warehouse) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("create_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

func (w *Warehouse) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

// 保存数据
func (w *Warehouse) SaveOne(ctx context.Context, name string, authorId int32) (*Warehouse, error) {
	w.IsDelete = constant.UserIsDelNo
	w.WarehouseName = name
	w.WarehouseAuthorId = authorId
	if err := w.GetMysqlDB(ctx).Model(w).Create(w).Error; err != nil {
		return nil, err
	}
	return w, nil
}

// 修改数据
func (w *Warehouse) ModifyOne(ctx context.Context, name string, authorId, id int32) (*Warehouse, error) {
	w.IsDelete = constant.UserIsDelNo
	w.Id = id
	w.WarehouseName = name
	w.WarehouseAuthorId = authorId
	if err := w.GetMysqlDB(ctx).Model(w).Updates(w).Error; err != nil {
		return nil, err
	}
	return w, nil
}

// 查询全部数据
func (w *Warehouse) GetList(ctx context.Context) ([]*Warehouse, error) {
	warehouseList := make([]*Warehouse, 0)
	if err := w.GetMysqlDB(ctx).Where("is_delete = ?", constant.UserIsDelNo).Model(w).Find(&warehouseList).Error; err != nil {
		return nil, err
	}
	return warehouseList, nil
}

// 根据ID查询单条数据
func (w *Warehouse) GetOne(ctx context.Context, id int32) (*Warehouse, error) {
	if err := w.GetMysqlDB(ctx).Model(w).Where("id = ? and is_delete = ?", id, constant.UserIsDelNo).Find(w).Error; err != nil {
		return nil, err
	}
	return w, nil
}

// 根据name查询
func (w *Warehouse) GetOneByName(ctx context.Context, name string) (*Warehouse, error) {
	if err := w.GetMysqlDB(ctx).Model(w).Where("warehouse_name = ?", name).Find(w).Error; err != nil {
		return nil, err
	}
	return w, nil
}
