package middleware

import (
	"context"
	"encoding/json"

	"goods/common/constant"
	"goods/common/model"
	"goods/common/util"
	"goods/common/util/redis"

	"github.com/gin-gonic/gin"
	"github.com/zztroot/zztlog"
)

const XToken = "X-TOKEN"

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := util.BuildRedisKey(constant.RedisKeyUserToken, c.Request.Header.Get(XToken))
		b, err := redis.Exists(context.Background(), key)
		if err != nil {
			zztlog.Error(constant.ErrorSystemError, "系统错误")
			util.Fail(c, util.New(constant.ErrorSystemError))
			c.Abort()
			return
		}
		if !b {
			util.Fail(c, util.New(constant.ErrorUserTokenFail))
			c.Abort()
			return
		}
		// 获取用户数据
		userByte, err := redis.Bytes(context.Background(), key)
		if err != nil {
			util.Fail(c, util.New(constant.ErrorUserTokenFail))
			c.Abort()
			return
		}
		// 更新token过期时间
		err = redis.Set(context.Background(), key, userByte, constant.TimeHour*3)
		if err != nil {
			zztlog.Error(err)
			c.Abort()
			return
		}
		user := new(model.User)
		if err := json.Unmarshal(userByte, user); err != nil {
			zztlog.Error(constant.ErrorSystemError, "系统错误")
			util.Fail(c, util.New(constant.ErrorSystemError))
			c.Abort()
			return
		}
		c.Set(constant.GinBindTokenUserKey, user)
		c.Next()
	}
}

type RedisUser struct {
	UserName     string `json:"user_name"`      // 用户账号
	UserRealName string `json:"user_real_name"` // 用户真实名字
	UserPhone    string `json:"user_phone"`     // 用户手机号
	UserAge      int    `json:"user_age"`       // 用户年龄
	UserSex      int    `json:"user_sex"`       // 用户性别 1男|2女
	UserRoleId   int32  `json:"user_role_id"`   // 用户角色ID
}

func GetUser(c *gin.Context) *model.User {
	return c.MustGet(constant.GinBindTokenUserKey).(*model.User)
}
