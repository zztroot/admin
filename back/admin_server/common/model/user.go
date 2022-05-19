package model

import (
	"context"

	"goods/common/constant"
	"goods/common/util"

	"github.com/jinzhu/gorm"
)

// 用户表
type User struct {
	Base
	MysqlConn
	UserName     string `json:"user_name"`      // 用户账号
	UserPassword string `json:"user_password"`  // 用户密码
	UserRealName string `json:"user_real_name"` // 用户真实名字
	UserPhone    string `json:"user_phone"`     // 用户手机号
	UserAge      int    `json:"user_age"`       // 用户年龄
	UserSex      int    `json:"user_sex"`       // 用户性别 1男|2女
	UserRoleId   int32  `json:"user_role_id"`   // 用户角色ID
	UserExtend   string `json:"user_extend"`    // 加密密码
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("create_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("update_time", util.Now()); err != nil {
		return util.Wrap(err)
	}
	return nil
}

// 新增数据
func (u *User) SaveOne(ctx context.Context, userName, userPwd, userRealName, userPhone, userExtend string, userAge, userSex int, userRoleId int32) (*User, error) {
	u.UserName = userName
	u.UserPassword = userPwd
	u.UserRealName = userRealName
	u.UserPhone = userPhone
	u.UserAge = userAge
	u.UserSex = userSex
	u.IsDelete = constant.UserIsDelNo
	u.UserRoleId = userRoleId
	u.UserExtend = userExtend
	if err := u.GetMysqlDB(ctx).Model(u).Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// 修改数据
func (u *User) ModifyOne(ctx context.Context, userName, userPwd, userRealName, userPhone, userExtend string, userAge, userSex int, id, userRoleId int32) error {
	user := User{}
	user.UserName = userName
	user.UserPassword = userPwd
	user.UserExtend = userExtend
	user.UserRealName = userRealName
	user.UserPhone = userPhone
	user.UserAge = userAge
	user.UserSex = userSex
	user.IsDelete = constant.UserIsDelNo
	user.UserRoleId = userRoleId
	if err := u.GetMysqlDB(ctx).Model(u).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// 根据用户账号查询数据
func (u *User) GetUserByUserName(ctx context.Context, userName string) (*User, error) {
	if err := u.GetMysqlDB(ctx).Model(u).Where("user_name = ? and is_delete = ?", userName, constant.UserIsDelNo).Find(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// 删除用户
func (u *User) DeleteUserById(ctx context.Context, id int32) error {
	if err := u.GetMysqlDB(ctx).Model(u).Where("id = ?", id).Update("is_delete", constant.UserIsDelYes).Error; err != nil {
		return err
	}
	return nil
}

// 查询用户分页
func (u *User) GetUserList(ctx context.Context, page, pageSize int32, userName string, sex int) ([]*User, int, error) {
	offset := (page - 1) * pageSize
	user := make([]*User, 0)
	total := 0
	db := u.GetMysqlDB(ctx).Model(u).Where("is_delete = ?", constant.UserIsDelNo)
	if userName != "" {
		db = db.Where("user_real_name like ?", "%"+userName+"%")
	}
	if sex != 0 {
		db = db.Where("user_sex = ?", sex)
	}
	if err := db.Offset(offset).Limit(pageSize).Order("id desc").Find(&user).Error; err != nil {
		return nil, 0, err
	}
	// 查询count
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return user, total, nil
}

// 根据ID查询用户
func (u *User) GetUserById(ctx context.Context, id int32) (*User, error) {
	if err := u.GetMysqlDB(ctx).Model(u).Where("id = ? and is_delete = ?", id, constant.UserIsDelNo).Find(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// 查询全部用户
func (u *User) GetUserAll(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err := u.GetMysqlDB(ctx).Model(u).Where("is_delete = ?", constant.UserIsDelNo).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
