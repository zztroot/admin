package model

import (
	"context"

	"goods/common/constant"
	"goods/common/util"

	"github.com/jinzhu/gorm"
)

type Role struct {
	Base
	MysqlConn
	RoleName   string `json:"role_name"`                    // 角色名称
	RoleExtend string `json:"role_extend" gorm:"type:text"` // 角色的权限数据
}

func (r *Role) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("create_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

func (r *Role) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

// 查询所有角色
func (r *Role) GetRoleList(ctx context.Context) ([]*Role, error) {
	roleList := make([]*Role, 0)
	if err := r.GetMysqlDB(ctx).Where("is_delete = ?", constant.UserIsDelNo).Model(r).Find(&roleList).Error; err != nil {
		return nil, err
	}
	return roleList, nil
}

// 修改数据-这个方法主要用于修改角色数据
func (r *Role) ModifyOne(ctx context.Context, roleName, roleExtend string, id int32) (*Role, error) {
	if err := r.GetMysqlDB(ctx).Model(r).Where("id = ?", id).Updates(map[string]interface{}{
		"role_name":   roleName,
		"role_extend": roleExtend,
	}).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

// 新增一个角色
func (r *Role) CreateOne(ctx context.Context, name string) (*Role, error) {
	r.RoleName = name
	r.IsDelete = constant.UserIsDelNo
	r.RoleExtend = "[]"
	if err := r.GetMysqlDB(ctx).Model(r).Create(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

// 根据名称查询
func (r *Role) GetOneByName(ctx context.Context, name string) (*Role, error) {
	if err := r.GetMysqlDB(ctx).Where("role_name = ? and is_delete = ?", name, constant.UserIsDelNo).Find(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

// 根据ID查询
func (r *Role) GetOneById(ctx context.Context, id int32) (*Role, error) {
	if err := r.GetMysqlDB(ctx).Model(r).Where("id = ? and is_delete = ?", id, constant.UserIsDelNo).Find(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

// 删除
func (r *Role) DeleteById(ctx context.Context, id int32) error {
	if err := r.GetMysqlDB(ctx).Model(r).Where("id = ?", id).Update("is_delete", constant.UserIsDelYes).Error; err != nil {
		return err
	}
	return nil
}
