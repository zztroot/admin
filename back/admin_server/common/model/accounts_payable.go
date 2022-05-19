package model

import (
	"context"

	"goods/common/constant"
	"goods/common/util"

	"github.com/jinzhu/gorm"
)

type AccountsPayable struct {
	Base
	MysqlConn
	AccountsPayableDate     string  `json:"accounts_payable_data"`      // 日期 精确到天
	AccountsPayableGoodsId  int32   `json:"accounts_payable_goods_id"`  // 商品ID
	AccountsPayableUnit     string  `json:"accounts_payable_unit"`      // 单位
	AccountsPayableNumber   int     `json:"accounts_payable_times"`     // 数量
	AccountsPayablePrice    float32 `json:"accounts_payable_price"`     // 价格
	AccountsPayableComments string  `json:"accounts_payable_comments"`  // 备注
	AccountsPayableAuthorId int32   `json:"accounts_payable_author_id"` // 创建者-用户ID
}

func (a *AccountsPayable) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("create_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

func (a *AccountsPayable) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

// 保存数据
func (a *AccountsPayable) SaveOne(ctx context.Context, date, unit, comments string, number int, goodsId, authorId int32, price float32) (*AccountsPayable, error) {
	a.AccountsPayableDate = date
	a.AccountsPayableUnit = unit
	a.AccountsPayableComments = comments
	a.AccountsPayableNumber = number
	a.AccountsPayableGoodsId = goodsId
	a.AccountsPayableAuthorId = authorId
	a.AccountsPayablePrice = price
	a.IsDelete = constant.UserIsDelNo
	if err := a.GetMysqlDB(ctx).Model(a).Save(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// 修改数据
func (a *AccountsPayable) ModifyOne(ctx context.Context, date, unit, comments string, number int, goodsId, authorId, id int32, price float32) (*AccountsPayable, error) {
	a.Id = id
	a.AccountsPayableDate = date
	a.AccountsPayableUnit = unit
	a.AccountsPayableComments = comments
	a.AccountsPayableNumber = number
	a.AccountsPayableGoodsId = goodsId
	a.AccountsPayableAuthorId = authorId
	a.AccountsPayablePrice = price
	a.IsDelete = constant.UserIsDelNo
	if err := a.GetMysqlDB(ctx).Model(a).Updates(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// 查询所有
func (a *AccountsPayable) GetList(ctx context.Context, page, pageSize, authorId int32, startDate, endDate, goodsName string) ([]*AccountsPayable, error) {
	offset := (page - 1) * pageSize
	accountsPayable := make([]*AccountsPayable, 0)
	db := a.GetMysqlDB(ctx).Model(a)
	if goodsName != "" {
		db = db.Joins("left join goods on accounts_payable.accounts_payable_goods_id = goods.id").Where("goods.goods_name like ?", "%"+goodsName+"%")
	}
	if startDate != "" {
		db = db.Where("accounts_payable_data between ? and ?", startDate, endDate)
	}
	if authorId != 0 {
		db = db.Where("accounts_payable_author_id = ?", authorId)
	}
	if err := db.Where("is_delete = ?", constant.UserIsDelNo).Offset(offset).Limit(pageSize).Order("id desc").Find(&accountsPayable).Error; err != nil {
		return nil, err
	}
	return accountsPayable, nil
}

// 删除
func (a *AccountsPayable) DeleteById(ctx context.Context, id int32) error {
	if err := a.GetMysqlDB(ctx).Model(a).Where("id = ?", id).Update("is_delete", constant.UserIsDelYes).Error; err != nil {
		return err
	}
	return nil
}

// 根据ID查询一条
func (a *AccountsPayable) GetOne(ctx context.Context, id int32) (*AccountsPayable, error) {
	if err := a.GetMysqlDB(ctx).Where("id = ?", id).Find(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}
