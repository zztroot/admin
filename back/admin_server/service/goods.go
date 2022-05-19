package service

import (
	"context"
	"github.com/zztroot/zztlog"

	"goods/common/constant"
	"goods/common/model"
	"goods/common/util"
)

type Goods struct{}

func (g *Goods) SaveOne(ctx context.Context, warehouseId, authorId int32, name, date, unit string, costPrice, supplyPrice float32, saleNumber, returnNumber, linkReturnNumber int) error {
	_, err := new(model.Goods).SaveOne(ctx, warehouseId, authorId, util.GenerateGoodsCode(), name, date, unit, costPrice, supplyPrice, saleNumber, returnNumber, linkReturnNumber)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (g *Goods) ModifyOne(ctx context.Context, id, warehouseId, authorId int32, name, date, unit string, costPrice, supplyPrice float32, saleNumber, returnNumber, linkReturnNumber int) error {
	_, err := new(model.Goods).ModifyOne(ctx, id, warehouseId, authorId, name, date, unit, costPrice, supplyPrice, saleNumber, returnNumber, linkReturnNumber)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (g *Goods) GetList(ctx context.Context, page, pageSize, authorId, warehouseId int32, startDate, endDate, name string) (interface{}, error) {
	list, total, err := new(model.Goods).GetList(ctx, page, pageSize, authorId, warehouseId, startDate, endDate, name)
	if err != nil {
		return nil, util.New(constant.ErrorDatabase, err)
	}
	type goods struct {
		Goods     *model.Goods `json:"goods"`     // 商品
		Warehouse string       `json:"warehouse"` // 仓库
		Author    string       `json:"author"`    // 创建者
		Income    struct {
			SalesVolume            float32 `json:"sales_volume"`              // 销售额
			ReturnVolume           float32 `json:"return_volume"`             // 客退额
			LinkReturnVolume       float32 `json:"link_return_volume"`        // 链路客退额
			SalesVolumeAfterReturn float32 `json:"sales_volume_after_return"` // 扣退后销售额
		} `json:"income"` // 业务收入
		Cost struct {
			GoodsCost            float32 `json:"goods_cost"`              // 商品成本
			ReturnCost           float32 `json:"return_cost"`             // 客退成本
			LinkReturnCost       float32 `json:"link_return_cost"`        // 链路客退成本
			GoodsCostAfterReturn float32 `json:"goods_cost_after_return"` // 扣退后产品成本
		} `json:"cost"` // 业务成本
		Profit struct {
			GrossProfit     float32 `json:"gross_profit"`      // 毛利额
			GrossProfitRate string  `json:"gross_profit_rate"` // 毛利率
			//NetProfit       float32 `json:"net_profit"`          // 净利额
			//NetInterestRate string  `json:"net_interest_rate"`   // 净利率
		} `json:"profit"` // 利润
		ReturnRate string `json:"return_rate"` // 客退率
	}
	results := make([]interface{}, 0)
	for _, v := range list {
		req := goods{}
		// 商品信息
		req.Goods = v
		// 查询作者
		user, err := new(model.User).GetUserById(ctx, v.AuthorId)
		if err != nil {
			zztlog.Error(err)
			return nil, nil
		}
		req.Author = user.UserRealName
		// 查询仓库
		warehouse, err := new(model.Warehouse).GetOne(ctx, v.GoodsWarehouseId)
		if err != nil {
			zztlog.Error(err)
			return nil, nil
		}
		req.Warehouse = warehouse.WarehouseName
		// 业务收入
		// 销售额
		req.Income.SalesVolume = v.GoodsSupplyPrice * float32(v.GoodsSaleNumber)
		// 客退额
		req.Income.ReturnVolume = v.GoodsSupplyPrice * float32(v.GoodsReturnNumber)
		// 链路客退
		req.Income.LinkReturnVolume = v.GoodsSupplyPrice * float32(v.GoodsLinkReturnNumber)
		// 扣退后销售额
		salesVolumeAfterReturn := v.GoodsSupplyPrice*float32(v.GoodsSaleNumber) - v.GoodsSupplyPrice*float32(v.GoodsReturnNumber) - v.GoodsSupplyPrice*float32(v.GoodsLinkReturnNumber)
		req.Income.SalesVolumeAfterReturn = salesVolumeAfterReturn
		// 业务成本
		// 产品成本
		req.Cost.GoodsCost = util.Float32ToFloat32(2, v.GoodsCostPrice*float32(v.GoodsSaleNumber))
		// 客退成本
		req.Cost.ReturnCost = util.Float32ToFloat32(2, v.GoodsCostPrice*float32(v.GoodsReturnNumber))
		// 链路客退成本
		req.Cost.LinkReturnCost = v.GoodsCostPrice * float32(v.GoodsLinkReturnNumber)
		// 扣退后产品成本
		goodsCostAfterReturn := util.Float32ToFloat32(2, v.GoodsCostPrice*float32(v.GoodsSaleNumber)-v.GoodsCostPrice*float32(v.GoodsReturnNumber)-v.GoodsCostPrice*float32(v.GoodsLinkReturnNumber))
		req.Cost.GoodsCostAfterReturn = goodsCostAfterReturn
		// 客退率
		req.ReturnRate = util.Float32ToString(2, float32((v.GoodsReturnNumber+v.GoodsLinkReturnNumber)/v.GoodsSaleNumber*100))
		// 毛利额
		req.Profit.GrossProfit = util.Float32ToFloat32(2, salesVolumeAfterReturn-goodsCostAfterReturn)
		// 毛利率
		req.Profit.GrossProfitRate = util.Float32ToString(2, util.Float32ToFloat32(2, salesVolumeAfterReturn-goodsCostAfterReturn)/salesVolumeAfterReturn*100)
		results = append(results, req)
	}
	m := make(map[string]interface{})
	m["data"] = results
	m["total"] = total
	return m, nil
}

func (g *Goods) DeleteOne(ctx context.Context, id int32) error {
	if err := new(model.Goods).DeleteById(ctx, id); err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}
