package service

import (
	"context"
	"encoding/json"

	"goods/common/constant"
	"goods/common/model"
	"goods/common/util"
	"goods/common/util/redis"
	"goods/web/form"

	"github.com/jinzhu/gorm"
	"github.com/zztroot/zztlog"
)

type User struct{}

func (u *User) Login(ctx context.Context, userName, pwd string) (interface{}, error) {
	// 校验账号和密码
	if err := util.CheckUserNameAndPwd(userName, pwd); err != nil {
		return nil, err
	}
	// 查询数据
	user, err := new(model.User).GetUserByUserName(ctx, userName)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, util.New(constant.ErrorUserNameOrPwdFail, err)
		}
		return nil, util.New(constant.ErrorDatabase, err)
	}
	// 判断密码
	if !util.ComparePasswords(user.UserPassword, []byte(pwd)) {
		return nil, util.New(constant.ErrorUserNameOrPwdFail, err)
	}
	// 生成token
	token := util.RandomString(constant.TokenGenerate)
	key := util.BuildRedisKey(constant.RedisKeyUserToken, token)
	userByte, _ := json.Marshal(user)
	err = redis.Set(ctx, key, userByte, constant.TimeHour*3)
	if err != nil {
		zztlog.Error(err)
		return nil, util.New(constant.ErrorDatabase, err)
	}
	// 查询用户权限
	var permissions []constant.WebMenu
	if user.UserRoleId != 0 {
		role, err := new(model.Role).GetOneById(ctx, user.UserRoleId)
		if err != nil {
			return nil, util.New(constant.ErrorDatabase)
		}
		rolePermissions := make([]constant.WebMenu, 0)
		if err := json.Unmarshal([]byte(role.RoleExtend), &rolePermissions); err != nil {
			return nil, util.New(constant.ErrorSystemError)
		}
		permissions = rolePermissions
	} else {
		permissions = constant.ListWebMenu
	}

	data := make(map[string]interface{})
	data["token"] = token
	data["user"] = map[string]interface{}{
		"user_name":   user.UserName,
		"real_name":   user.UserRealName,
		"id":          user.Id,
		"age":         user.UserAge,
		"sex":         constant.SexMap[user.UserSex],
		"phone":       user.UserPhone,
		"role":        user.UserRoleId, // 管理员role为0
		"permissions": permissions,
	}
	return data, nil
}

func (u *User) SaveOne(ctx context.Context, params *form.SaveUserForm) error {
	// 查询两次密码是否一样
	if params.UserPassword != params.AgainPassword {
		return util.New(constant.ErrorTwoPasswordNotEqual)
	}
	// 校验账号和密码
	if err := util.CheckUserNameAndPwd(params.UserName, params.UserPassword); err != nil {
		return err
	}
	// 判断用户是否存在
	user, err := new(model.User).GetUserByUserName(ctx, params.UserName)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		zztlog.Error(err)
		return util.New(constant.ErrorDatabase, err)
	}
	if user != nil && user.Id != 0 {
		return util.New(constant.ErrorUserRepeat)
	}
	// 创建用户
	// 密码加密
	password := util.HashAndSalt([]byte(params.UserPassword))
	_, err = new(model.User).SaveOne(ctx, params.UserName, password, params.UserRealName, params.UserPhone, util.EncodeStr(params.UserPassword), params.UserAge, params.UserSex, params.UserRoleId)
	if err != nil {
		zztlog.Error(err)
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (u *User) ModifyOne(ctx context.Context, params *form.ModifyUserForm, user *model.User) error {
	// 查询两次密码是否一样
	if params.NewUserPassword != params.AgainPassword {
		return util.New(constant.ErrorTwoPasswordNotEqual)
	}
	// 如果不是管理员修改密码，需要验证登录密码
	if user.UserRoleId != 0 {
		// 不是管理员
		if params.UserPassword == "" {
			return util.New(constant.ErrorParams)
		}
		u, err := new(model.User).GetUserByUserName(ctx, user.UserName)
		if err != nil {
			return util.New(constant.ErrorDatabase, err)
		}
		// 判断密码
		if !util.ComparePasswords(u.UserPassword, []byte(params.UserPassword)) {
			return util.New(constant.ErrorUserNameOrPwdFail, err)
		}
	}
	// 判断用户名是否存在
	user, err := new(model.User).GetUserByUserName(ctx, params.UserName)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		zztlog.Error(err)
		return util.New(constant.ErrorDatabase, err)
	}
	if user != nil && user.Id != 0 && user.Id != params.Id {
		return util.New(constant.ErrorUserRepeat)
	}
	var (
		password string
		extend   string
	)
	if params.NewUserPassword != "" {
		password = util.HashAndSalt([]byte(params.NewUserPassword))
		extend = util.EncodeStr(params.NewUserPassword)
	}
	err = new(model.User).ModifyOne(ctx, params.UserName, password, params.UserRealName, params.UserPhone, extend, params.UserAge, params.UserSex, params.Id, params.UserRoleId)
	if err != nil {
		zztlog.Error(err)
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

type getUserResponse struct {
	model.Base
	UserName     string `json:"user_name"`      // 用户账号
	UserRealName string `json:"user_real_name"` // 用户真实名字
	UserPhone    string `json:"user_phone"`     // 用户手机号
	UserAge      int    `json:"user_age"`       // 用户年龄
	UserSex      string `json:"user_sex"`       // 用户性别 1男|2女
	UserRoleId   int32  `json:"user_role_id"`   // 用户角色ID
	UserRole     string `json:"user_role"`      // 用户角色
}

func (u *User) GetUserList(ctx context.Context, page, pageSize int32, userName string, sex int) (interface{}, error) {
	data, total, err := new(model.User).GetUserList(ctx, page, pageSize, userName, sex)
	if err != nil {
		return nil, util.New(constant.ErrorDatabase, err)
	}
	getUserResponseList := make([]getUserResponse, 0)
	for _, v := range data {
		// 查询角色
		var roleName string
		if v.UserRoleId != 0 {
			role, err := new(model.Role).GetOneById(ctx, v.UserRoleId)
			if err != nil {
				zztlog.Error(err)
			}
			roleName = ""
			if role != nil {
				roleName = role.RoleName
			}
		} else {
			roleName = "管理员"
		}
		getUserResponseList = append(getUserResponseList, getUserResponse{
			Base:         v.Base,
			UserName:     v.UserName,
			UserRealName: v.UserRealName,
			UserPhone:    v.UserPhone,
			UserAge:      v.UserAge,
			UserSex:      constant.SexMap[v.UserSex],
			UserRoleId:   v.UserRoleId,
			UserRole:     roleName,
		})
	}
	results := make(map[string]interface{})
	results["data"] = getUserResponseList
	results["total"] = total
	return results, nil
}

func (u *User) GetUserOne(ctx context.Context, id int32) (interface{}, error) {
	user, err := new(model.User).GetUserById(ctx, id)
	if err != nil {
		return nil, util.New(constant.ErrorDatabase, err)
	}
	userResponse := getUserResponse{}
	userResponse.Id = user.Id
	userResponse.UserName = user.UserName
	userResponse.UserSex = constant.SexMap[user.UserSex]
	userResponse.UserAge = user.UserAge
	userResponse.UserRealName = user.UserRealName
	userResponse.UserPhone = user.UserPhone
	userResponse.UserRoleId = user.UserRoleId
	// 查询角色
	var roleName string
	if user.UserRoleId != 0 {
		role, err := new(model.Role).GetOneById(ctx, user.UserRoleId)
		if err != nil {
			zztlog.Error(err)
		}
		roleName = role.RoleName
	} else {
		roleName = "管理员"
	}
	userResponse.UserRole = roleName
	return userResponse, nil
}

func (u *User) DeleteUserOne(ctx context.Context, id int32) error {
	if err := new(model.User).DeleteUserById(ctx, id); err != nil {
		return util.New(constant.ErrorDatabase, err)
	}
	return nil
}

func (u *User) GetUserAll(ctx context.Context) (interface{}, error) {
	exists, err := redis.Exists(ctx, constant.RedisKeyUserData)
	if err != nil {
		zztlog.Error(err)
	}
	results := make([]interface{}, 0)
	if exists {
		users := make([]*model.User, 0)
		bytes, err := redis.Bytes(ctx, constant.RedisKeyUserData)
		if err != nil {
			zztlog.Error(err)
		}
		if err := json.Unmarshal(bytes, &users); err != nil {
			zztlog.Error(err)
		}
		for i, _ := range users {
			m := make(map[string]interface{})
			m["name"] = users[i].UserRealName
			m["id"] = users[i].Id
			results = append(results, m)
		}
	} else {
		// 查询数据
		users, err := new(model.User).GetUserAll(ctx)
		if err != nil {
			return nil, util.New(constant.ErrorDatabase, err)
		}
		// 将数据存入数据库
		userByte, _ := json.Marshal(users)
		if err := redis.Set(ctx, constant.RedisKeyUserData, userByte, constant.TimeMinute*10); err != nil {
			zztlog.Error(err)
		}
		for i, _ := range users {
			m := make(map[string]interface{})
			m["name"] = users[i].UserRealName
			m["id"] = users[i].Id
			results = append(results, m)
		}
	}
	return results, nil
}
