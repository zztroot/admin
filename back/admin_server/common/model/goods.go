package model

import (
	"context"

	"goods/common/constant"
	"goods/common/util"

	"github.com/jinzhu/gorm"
)

type Goods struct {
	Base
	MysqlConn
	GoodsCode             string  `json:"goods_code"`               // 商品编号
	GoodsName             string  `json:"goods_name"`               // 商品名称
	GoodsSaleDate         string  `json:"goods_sale_date"`          // 商品销售日期
	GoodsWarehouseId      int32   `json:"goods_warehouse_id"`       // 仓库ID
	GoodsUnit             string  `json:"goods_unit"`               // 商品单位
	GoodsCostPrice        float32 `json:"goods_cost_price"`         // 商品成本价
	GoodsSupplyPrice      float32 `json:"goods_supply_price"`       // 供货价格
	GoodsSaleNumber       int     `json:"goods_sale_number"`        // 销售数量
	GoodsReturnNumber     int     `json:"goods_return_number"`      // 退货数量
	GoodsLinkReturnNumber int     `json:"goods_link_return_number"` // 链路客退数量
	AuthorId              int32   `json:"author_id"`                // 创建用户ID
}

func (g *Goods) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("create_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

func (g *Goods) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

// 保存数据
func (g *Goods) SaveOne(ctx context.Context, warehouseId, authorId int32, code, name, date, unit string, costPrice, supplyPrice float32, saleNumber, returnNumber, linkReturnNumber int) (*Goods, error) {
	g.GoodsWarehouseId = warehouseId
	g.AuthorId = authorId
	g.GoodsCode = code
	g.GoodsName = name
	g.GoodsSaleDate = date
	g.GoodsUnit = unit
	g.GoodsCostPrice = costPrice
	g.GoodsSupplyPrice = supplyPrice
	g.GoodsSaleNumber = saleNumber
	g.GoodsReturnNumber = returnNumber
	g.GoodsLinkReturnNumber = linkReturnNumber
	g.IsDelete = constant.UserIsDelNo
	if err := g.GetMysqlDB(ctx).Model(g).Save(g).Error; err != nil {
		return nil, err
	}
	return g, nil
}

// 修改数据
func (g *Goods) ModifyOne(ctx context.Context, id, warehouseId, authorId int32, name, date, unit string, costPrice, supplyPrice float32, saleNumber, returnNumber, linkReturnNumber int) (*Goods, error) {
	g.Id = id
	g.GoodsWarehouseId = warehouseId
	g.AuthorId = authorId
	g.GoodsName = name
	g.GoodsSaleDate = date
	g.GoodsUnit = unit
	g.GoodsCostPrice = costPrice
	g.GoodsSupplyPrice = supplyPrice
	g.GoodsSaleNumber = saleNumber
	g.GoodsReturnNumber = returnNumber
	g.GoodsLinkReturnNumber = linkReturnNumber
	g.IsDelete = constant.UserIsDelNo
	if err := g.GetMysqlDB(ctx).Model(g).Updates(g).Error; err != nil {
		return nil, err
	}
	return g, nil
}

// 查询全部数据
func (g *Goods) GetList(ctx context.Context, page, pageSize, authorId, warehouseId int32, startDate, endDate, name string) ([]*Goods, int, error) {
	var total int
	offset := (page - 1) * pageSize
	goodsList := make([]*Goods, 0)
	db := g.GetMysqlDB(ctx).Model(g)
	if startDate != "" {
		db = db.Where("goods_sale_date between ? and ?", startDate, endDate)
	}
	if authorId != 0 {
		db = db.Where("author_id = ?", authorId)
	}
	if name != "" {
		db = db.Where("goods_name like ?", "%"+name+"%")
	}
	if warehouseId != 0 {
		db = db.Where("goods_warehouse_id = ?", warehouseId)
	}
	if err := db.Where("is_delete = ?", constant.UserIsDelNo).Offset(offset).Limit(pageSize).Order("id desc").Find(&goodsList).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Where("is_delete = ?", constant.UserIsDelNo).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return goodsList, total, nil
}

// 删除
func (g *Goods) DeleteById(ctx context.Context, id int32) error {
	if err := g.GetMysqlDB(ctx).Model(g).Where("id = ?", id).Update("is_delete", constant.UserIsDelYes).Error; err != nil {
		return err
	}
	return nil
}

// 获取单条
func (g *Goods) GetOneById(ctx context.Context, id int32) (*Goods, error) {
	if err := g.GetMysqlDB(ctx).Model(g).Where("id = ?", id).Find(g).Error; err != nil {
		return nil, err
	}
	return g, nil
}

// 根据code获取
func (g *Goods) GetOneByCode(ctx context.Context, code string) (*Goods, error) {
	if err := g.GetMysqlDB(ctx).Model(g).Where("goods_code = ?", code).Find(g).Error; err != nil {
		return nil, err
	}
	return g, nil
}
