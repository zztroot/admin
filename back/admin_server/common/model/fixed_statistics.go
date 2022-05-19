package model

import (
	"context"

	"goods/common/constant"
	"goods/common/util"

	"github.com/jinzhu/gorm"
)

type FixedStatistics struct {
	Base
	MysqlConn
	Date           string  `json:"date"`            // 日期
	Wages          float32 `json:"wages"`           // 人工工资和福利费
	Rent           float32 `json:"rent"`            // 房租
	SocialSecurity float32 `json:"social_security"` // 社保费
	PcCost         float32 `json:"pc_cost"`         // 电脑费用
	OtherCost      float32 `json:"other_cost"`      // 水电网费
	AuthorId       int32   `json:"author_id"`       // 用户ID
}

func (f *FixedStatistics) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("create_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

func (f *FixedStatistics) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

// 保存数据
func (f *FixedStatistics) SaveOne(ctx context.Context, date string, wages, rent, socialSecurity, pcCost, otherCost float32, authorId int32) (*FixedStatistics, error) {
	f.Date = date
	f.Wages = wages
	f.Rent = rent
	f.SocialSecurity = socialSecurity
	f.PcCost = pcCost
	f.OtherCost = otherCost
	f.AuthorId = authorId
	f.IsDelete = constant.UserIsDelNo
	if err := f.GetMysqlDB(ctx).Model(f).Save(f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

// 修改数据
func (f *FixedStatistics) ModifyOne(ctx context.Context, date string, wages, rent, socialSecurity, pcCost, otherCost float32, authorId, id int32) (*FixedStatistics, error) {
	f.Id = id
	f.Date = date
	f.Wages = wages
	f.Rent = rent
	f.SocialSecurity = socialSecurity
	f.PcCost = pcCost
	f.OtherCost = otherCost
	f.AuthorId = authorId
	f.IsDelete = constant.UserIsDelNo
	if err := f.GetMysqlDB(ctx).Model(f).Updates(f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

// 查询全部数据
func (f *FixedStatistics) GetList(ctx context.Context, page, pageSize, authorId int32, startDate, endDate string) ([]*FixedStatistics, error) {
	offset := (page - 1) * pageSize
	fixedStatistics := make([]*FixedStatistics, 0)
	db := f.GetMysqlDB(ctx).Model(f)
	if startDate != "" {
		db = db.Where("date between ? and ?", startDate, endDate)
	}
	if authorId != 0 {
		db = db.Where("author_id = ?", authorId)
	}
	if err := db.Where("is_delete = ?", constant.UserIsDelNo).Offset(offset).Limit(pageSize).Order("id desc").Find(&fixedStatistics).Error; err != nil {
		return nil, err
	}
	return fixedStatistics, nil
}

// 删除
func (f *FixedStatistics) DeleteById(ctx context.Context, id int32) error {
	if err := f.GetMysqlDB(ctx).Model(f).Where("id = ?", id).Update("is_delete", constant.UserIsDelYes).Error; err != nil {
		return err
	}
	return nil
}
