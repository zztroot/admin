package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/zztroot/zztlog"

	"goods/common/constant"
	"goods/common/model"
	"goods/common/util"
)

type PlatformStatistics struct{}

func (p *PlatformStatistics) SaveOne(ctx context.Context, platformId, authorId int32, price float32, code, comments string) error {
	// 查询商品
	goods, err := new(model.Goods).GetOneByCode(ctx, code)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return util.New(constant.ErrorGoodsNotFind)
		} else {
			return util.New(constant.ErrorDatabase)
		}
	}
	_, err = new(model.PlatformStatistics).SaveOne(ctx, platformId, authorId, goods.Id, price, comments)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (p *PlatformStatistics) ModifyOne(ctx context.Context, id, platformId, authorId, salesTimes int32, price float32, code, comments string) error {
	// 查询商品
	goods, err := new(model.Goods).GetOneByCode(ctx, code)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return util.New(constant.ErrorGoodsNotFind)
		} else {
			return util.New(constant.ErrorDatabase)
		}
	}
	_, err = new(model.PlatformStatistics).ModifyOne(ctx, id, platformId, authorId, salesTimes, goods.Id, price, comments)
	if err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (p *PlatformStatistics) GetList(ctx context.Context, page, pageSize, authorId, platformId int32, goodsName string) (interface{}, error) {
	list, total, err := new(model.PlatformStatistics).GetList(ctx, page, pageSize, authorId, platformId, goodsName)
	if err != nil {
		return nil, util.New(constant.ErrorDatabase, err)
	}
	type platForm struct {
		PlatForm struct {
			PlatformModel       *model.PlatformStatistics `json:"platform_model"`
			PlatformSaleVolume  float32                   `json:"platform_sale_volume"`  // 平台销售额
			PlatformGrossProfit float32                   `json:"platform_gross_profit"` // 平台毛利额
			PlatformProfitRate  string                    `json:"platform_profit_rate"`  // 平台毛利率
			PlatformStr         string                    `json:"platform_str"`          // 平台
			Author              string                    `json:"author"`                // 创建者
		} `json:"plat_form"`
		Goods struct {
			GoodsModel     *model.Goods `json:"goods_model"`
			GoodsCostTotal float32      `json:"goods_cost_total"` // 商品成本
			Warehouse      string       `json:"warehouse"`        // 仓库
		} `json:"goods"`
		Company struct {
			CompanySaleVolume     float32 `json:"company_sale_volume"`      // 公司销售额
			CompanyOneGrossProfit float32 `json:"company_one_gross_profit"` // 公司单袋毛利额
			CompanyGrossProfit    float32 `json:"company_gross_profit"`     // 公司毛利额
			CompanyProfitRate     string  `json:"company_profit_rate"`      // 公司毛利率
		} `json:"company"`
	}
	results := make([]interface{}, 0)
	for _, v := range list {
		p := platForm{}
		// 平台基础数据
		p.PlatForm.PlatformModel = v
		// 创建者
		user, err := new(model.User).GetUserById(ctx, v.AuthorId)
		if err != nil {
			zztlog.Error(err)
		}
		p.PlatForm.Author = user.UserRealName
		// 平台
		pf, err := new(model.Platform).GetOne(ctx, v.PlatformId)
		if err != nil {
			zztlog.Error(err)
		}
		p.PlatForm.PlatformStr = pf.PlatformName
		// 商品
		goods, err := new(model.Goods).GetOneById(ctx, v.GoodsId)
		if err != nil {
			zztlog.Error(err)
		}
		p.Goods.GoodsModel = goods
		// 商品成本
		goodsCostTotal := goods.GoodsCostPrice * float32(goods.GoodsSaleNumber)
		p.Goods.GoodsCostTotal = util.Float32ToFloat32(2, goodsCostTotal)
		// 平台销售额
		platformSaleVolume := float32(goods.GoodsSaleNumber) * v.PlatformPrice
		p.PlatForm.PlatformSaleVolume = util.Float32ToFloat32(2, platformSaleVolume)
		// 平台毛利额
		platformGrossProfit := float32(goods.GoodsSaleNumber)*v.PlatformPrice - goods.GoodsSupplyPrice*float32(goods.GoodsSaleNumber)
		p.PlatForm.PlatformGrossProfit = util.Float32ToFloat32(2, platformGrossProfit)
		// 平台毛利率
		p.PlatForm.PlatformProfitRate = util.Float32ToString(2, platformGrossProfit/platformSaleVolume*100)
		// 公司销售额
		companySaleVolume := goods.GoodsSupplyPrice * float32(goods.GoodsSaleNumber)
		p.Company.CompanySaleVolume = util.Float32ToFloat32(2, companySaleVolume)
		// 公司单袋毛利额
		p.Company.CompanyOneGrossProfit = util.Float32ToFloat32(2, goods.GoodsSupplyPrice-goods.GoodsCostPrice)
		// 公司毛利额
		companyGrossProfit := companySaleVolume - goodsCostTotal
		p.Company.CompanyGrossProfit = util.Float32ToFloat32(2, companyGrossProfit)
		// 公司毛利率
		p.Company.CompanyProfitRate = util.Float32ToString(2, companyGrossProfit/companySaleVolume*100)
		// 仓库
		warehouse, err := new(model.Warehouse).GetOne(ctx, goods.GoodsWarehouseId)
		if err != nil {
			zztlog.Error(err)
		}
		p.Goods.Warehouse = warehouse.WarehouseName
		results = append(results, p)
	}
	m := make(map[string]interface{})
	m["data"] = results
	m["total"] = total
	return m, nil
}

func (p *PlatformStatistics) DeleteOne(ctx context.Context, id int32) error {
	if err := new(model.PlatformStatistics).DeleteById(ctx, id); err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}
