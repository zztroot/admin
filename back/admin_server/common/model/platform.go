package model

import (
	"context"

	"goods/common/constant"
	"goods/common/util"

	"github.com/jinzhu/gorm"
)

type Platform struct {
	Base
	MysqlConn
	PlatformName     string `json:"platform_name"`      // 平台名称
	PlatformAuthorId int32  `json:"platform_author_id"` // 平台创建者
}

func (p *Platform) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("create_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

func (p *Platform) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

// 保存数据
func (p *Platform) SaveOne(ctx context.Context, authorId int32, name string) (*Platform, error) {
	p.IsDelete = constant.UserIsDelNo
	p.PlatformAuthorId = authorId
	p.PlatformName = name
	if err := p.GetMysqlDB(ctx).Model(p).Save(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

// 修改数据
func (p *Platform) ModifyOne(ctx context.Context, id, authorId int32, name string) (*Platform, error) {
	p.IsDelete = constant.UserIsDelNo
	p.Id = id
	p.PlatformAuthorId = authorId
	p.PlatformName = name
	if err := p.GetMysqlDB(ctx).Model(p).Updates(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

// 查询全部数据
func (p *Platform) GetList(ctx context.Context) ([]*Platform, error) {
	platformList := make([]*Platform, 0)
	if err := p.GetMysqlDB(ctx).Where("is_delete = ?", constant.UserIsDelNo).Model(p).Find(&platformList).Error; err != nil {
		return nil, err
	}
	return platformList, nil
}

// 根据ID查询单条数据
func (p *Platform) GetOne(ctx context.Context, id int32) (*Platform, error) {
	if err := p.GetMysqlDB(ctx).Model(p).Where("id = ? and is_delete = ?", id, constant.UserIsDelNo).Find(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}
