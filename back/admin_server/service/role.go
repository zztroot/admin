package service

import (
	"context"
	"encoding/json"
	"github.com/zztroot/zztlog"
	"goods/common/constant"
	"goods/common/model"
	"goods/common/util"

	"github.com/jinzhu/gorm"
)

type Role struct{}

// 查询全部
func (r *Role) GetList(ctx context.Context) (interface{}, error) {
	list, err := new(model.Role).GetRoleList(ctx)
	if err != nil {
		return nil, util.New(constant.ErrorDatabase)
	}
	return list, nil
}

// 修改角色 -- 主要用于修改
func (r *Role) ModifyOne(ctx context.Context, roleName, roleExtend string, id int32) error {
	webMenuList := make([]constant.WebMenu, 0)
	if err := json.Unmarshal([]byte(roleExtend), &webMenuList); err != nil {
		zztlog.Error(err)
		return util.New(constant.ErrorSystemError)
	}
	menu := make([]constant.WebMenu, 0)
	for _, v := range constant.Menu {
		if v.Id == 1 {
			menu = append(menu, v)
			continue
		}
		isPush := false
		for i, _ := range webMenuList {
			if webMenuList[i].FatherId == v.Id {
				isPush = true
				v.Children = append(v.Children, &webMenuList[i])
			} else if webMenuList[i].Id == v.Id {
				isPush = true
			}
		}
		if isPush {
			menu = append(menu, v)
		}
	}
	extend, err := json.Marshal(menu)
	if err != nil {
		zztlog.Error(err)
	}
	// 判断角色是否存在
	role, err := new(model.Role).GetOneByName(ctx, roleName)
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return util.New(constant.ErrorDatabase)
		}
	}
	if role != nil && role.Id != 0 && role.Id != id {
		return util.New(constant.ErrorRoleRepeat)
	}
	// 保存
	_, err = new(model.Role).ModifyOne(ctx, roleName, string(extend), id)
	if err != nil {
		return util.New(constant.ErrorDatabase)
	}
	return nil
}

// 删除角色
func (r *Role) Delete(ctx context.Context, id int32) error {
	if err := new(model.Role).DeleteById(ctx, id); err != nil {
		return util.New(constant.ErrorDatabase)
	}
	return nil
}

// 根据名称查询
func (r *Role) GetOneByName(ctx context.Context, name string) (interface{}, error) {
	role, err := new(model.Role).GetOneByName(ctx, name)
	if err != nil {
		return nil, util.New(constant.ErrorDatabase)
	}
	roleExtend := role.RoleExtend
	menu := new([]constant.WebMenu)
	if err := json.Unmarshal([]byte(roleExtend), menu); err != nil {
		zztlog.Error(err)
		return nil, err
	}
	m := make(map[string]interface{})
	m["auth"] = menu
	return m, nil
}

// 创建一个角色
func (r *Role) CreateOne(ctx context.Context, name string) error {
	role, err := new(model.Role).GetOneByName(ctx, name)
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return util.New(constant.ErrorDatabase)
		}
	}
	if role != nil && role.Id != 0 {
		return util.New(constant.ErrorRoleRepeat)
	}
	_, err = new(model.Role).CreateOne(ctx, name)
	if err != nil {
		return util.New(constant.ErrorDatabase)
	}
	return nil
}
