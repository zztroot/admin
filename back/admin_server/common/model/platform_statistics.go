package model

import (
	"context"

	"goods/common/constant"
	"goods/common/util"

	"github.com/jinzhu/gorm"
)

type PlatformStatistics struct {
	Base
	MysqlConn
	GoodsId       int32   `json:"goods_id"`       // 商品ID
	PlatformId    int32   `json:"platform_id"`    // 平台ID
	PlatformPrice float32 `json:"platform_price"` // 平台售价
	//PlatformSalesTimes int32   `json:"platform_sales_times"` // 平台销售数量
	PlatformComments string `json:"platform_comments"` // 平台备注
	AuthorId         int32  `json:"author_id"`         // 用户ID
}

func (p *PlatformStatistics) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("create_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

func (p *PlatformStatistics) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

// 保存数据
func (p *PlatformStatistics) SaveOne(ctx context.Context, platformId, authorId, goodsId int32, price float32, comments string) (*PlatformStatistics, error) {
	p.IsDelete = constant.UserIsDelNo
	p.PlatformId = platformId
	p.AuthorId = authorId
	p.PlatformPrice = price
	//p.PlatformSalesTimes = salesTimes
	p.PlatformComments = comments
	p.GoodsId = goodsId
	if err := p.GetMysqlDB(ctx).Model(p).Save(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

// 修改数据
func (p *PlatformStatistics) ModifyOne(ctx context.Context, id, platformId, authorId, salesTimes, goodsId int32, price float32, comments string) (*PlatformStatistics, error) {
	p.Id = id
	p.IsDelete = constant.UserIsDelNo
	p.PlatformId = platformId
	p.AuthorId = authorId
	p.PlatformPrice = price
	//p.PlatformSalesTimes = salesTimes
	p.PlatformComments = comments
	p.GoodsId = goodsId
	if err := p.GetMysqlDB(ctx).Model(p).Updates(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

// 查询全部数据
func (p *PlatformStatistics) GetList(ctx context.Context, page, pageSize, authorId, platformId int32, goodsName string) ([]*PlatformStatistics, int, error) {
	var total int
	offset := (page - 1) * pageSize
	platformStatistics := make([]*PlatformStatistics, 0)
	db := p.GetMysqlDB(ctx).Model(p)
	if authorId != 0 {
		db = db.Where("platform_statistics.author_id = ?", authorId)
	}
	if platformId != 0 {
		db = db.Where("platform_statistics.platform_id = ?", platformId)
	}
	if goodsName != "" {
		db = db.Joins("left join goods on platform_statistics.goods_id = goods.id").Where("goods.goods_name like ?", "%"+goodsName+"%")
	}
	if err := db.Where("platform_statistics.is_delete = ?", constant.UserIsDelNo).Offset(offset).Limit(pageSize).Order("platform_statistics.id desc").Find(&platformStatistics).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Where("platform_statistics.is_delete = ?", constant.UserIsDelNo).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return platformStatistics, total, nil
}

// 删除
func (p *PlatformStatistics) DeleteById(ctx context.Context, id int32) error {
	if err := p.GetMysqlDB(ctx).Model(p).Where("id = ?", id).Update("is_delete", constant.UserIsDelYes).Error; err != nil {
		return err
	}
	return nil
}
