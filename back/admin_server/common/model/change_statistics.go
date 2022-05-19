package model

import (
	"context"

	"goods/common/constant"
	"goods/common/util"

	"github.com/jinzhu/gorm"
)

type ChangeStatistics struct {
	Base
	MysqlConn
	Date string `json:"date"` // 日期
	// 仓费用估算情况（汇总量）
	WarehouseCost          float32 `json:"warehouse_cost"`            // 仓库租金
	HandWarehouseGoodsCost float32 `json:"warehouse_goods_cost"`      // 交仓货费
	BackWarehouseGoodsCost float32 `json:"back_warehouse_goods_cost"` // 退仓货费
	ChangeWarehouseCost    float32 `json:"change_warehouse_cost"`     // 转仓运费
	PackCost               float32 `json:"pack_cost"`                 // 打包/标签
	// 协同仓费用估算情况（汇总量）
	CoordinationWarehouseCost  float32 `json:"coordination_warehouse_cost"`    // 协同仓仓库租金
	OperationAndUnloadingCost  float32 `json:"operation_and_unloading_cost"`   // 收货操作费/收货装卸费
	ReplenishmentCost          float32 `json:"replenishment_cost"`             // 补货费
	ManyBackWarehouseGoodsCost float32 `json:"many_back_warehouse_goods_cost"` // 多货返仓费
	OperationBackGoodsCost     float32 `json:"operation_back_goods_cost"`      // 操作退货费
	ChangeCost                 float32 `json:"change_cost"`                    // 调拨转运费
	UserBackGoodsCost          float32 `json:"user_back_goods_cost"`           // 客退调拨费
	// 当天采购进货费用支出
	DeliveryCost  float32 `json:"delivery_cost"`    // 私人交货费
	LogisticsCost float32 `json:"logistics_cost"`   // 物流费
	GoodsLaLaCost float32 `json:"goods_la_la_cost"` // 货拉拉费
	OtherCost     float32 `json:"other_cost"`       // 设计费等杂费
	AuthorId      int32   `json:"author_id"`        // 用户ID
}

func (c *ChangeStatistics) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("create_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

func (c *ChangeStatistics) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

// 保存数据
func (c *ChangeStatistics) SaveOne(ctx context.Context, change *ChangeStatistics) (*ChangeStatistics, error) {
	change.IsDelete = constant.UserIsDelNo
	if err := c.GetMysqlDB(ctx).Model(change).Save(change).Error; err != nil {
		return nil, err
	}
	return change, nil
}

// 修改数据
func (c *ChangeStatistics) ModifyOne(ctx context.Context, change *ChangeStatistics) (*ChangeStatistics, error) {
	change.IsDelete = constant.UserIsDelNo
	if err := c.GetMysqlDB(ctx).Model(change).Updates(change).Error; err != nil {
		return nil, err
	}
	return change, nil
}

// 查询所有数据
func (c *ChangeStatistics) GetList(ctx context.Context, page, pageSize, authorId int32, startDate, endDate string) ([]*ChangeStatistics, error) {
	offset := (page - 1) * pageSize
	changeWarehouseCost := make([]*ChangeStatistics, 0)
	db := c.GetMysqlDB(ctx).Model(c)
	if startDate != "" {
		db = db.Where("date between ? and ?", startDate, endDate)
	}
	if authorId != 0 {
		db = db.Where("author_id = ?", authorId)
	}
	if err := db.Where("is_delete = ?", constant.UserIsDelNo).Offset(offset).Limit(pageSize).Order("id desc").Find(&changeWarehouseCost).Error; err != nil {
		return nil, err
	}
	return changeWarehouseCost, nil
}

// 删除
func (c *ChangeStatistics) DeleteById(ctx context.Context, id int32) error {
	if err := c.GetMysqlDB(ctx).Model(c).Where("id = ?", id).Update("is_delete", constant.UserIsDelYes).Error; err != nil {
		return err
	}
	return nil
}
