package constant

const (
	// 成功
	ErrorOk = 0

	// 系统 - 1
	ErrorSystemError     = 1000 // 系统异常
	ErrorSystemRateLimit = 1001 // 请求过于频繁
	ErrorParams          = 1002 // 参数错误

	// 数据库 - 2
	ErrorDatabase = 2000 // 数据库错误

	// 用户 - 3
	ErrorUserRepeat             = 3000 // 用户已存在
	ErrorUserNameContainChinese = 3002 // 账号包含中文
	ErrorPwdContainChinese      = 3003 // 密码包含中文
	ErrorUserNameOutOfRange     = 3004 // 账号超出范围
	ErrorPwdOutOfRange          = 3005 // 密码超出范围
	ErrorUserNameOrPwdFail      = 3006 // 账号或密码错误
	ErrorUserTokenFail          = 3007 // token错误，无效的登录
	ErrorTwoPasswordNotEqual    = 3008 // 两次密码不相等
	ErrorNoPermission           = 3009 // 没有权限

	// 角色 - 4
	ErrorRoleRepeat = 4000 // 角色已经存在

	// 仓库 - 5
	ErrorWarehouseRepeat = 5000 // 仓库已经存在

	// 商品 - 6
	ErrorGoodsNotFind = 6000 // 没有该商品
)

// error和message映射表
var errorMap = map[int]string{
	ErrorOk:                     "成功",
	ErrorSystemError:            "系统异常",
	ErrorDatabase:               "数据库错误",
	ErrorUserRepeat:             "用户已存在",
	ErrorParams:                 "参数错误",
	ErrorUserNameContainChinese: "账号不能包含中文",
	ErrorPwdContainChinese:      "密码不能包含中文",
	ErrorUserNameOutOfRange:     "账号不能小于5位或大于15位",
	ErrorPwdOutOfRange:          "密码不能小于5位",
	ErrorUserNameOrPwdFail:      "账号或密码错误",
	ErrorUserTokenFail:          "登录已过期，请重新登录",
	ErrorRoleRepeat:             "角色已经存在",
	ErrorTwoPasswordNotEqual:    "两次密码不相等",
	ErrorNoPermission:           "没有权限访问",
	ErrorWarehouseRepeat:        "仓库已经存在",
	ErrorGoodsNotFind:           "没有该商品",
}

type Error struct{}

func (e *Error) ErrorMessage(code int) string {
	if message, ok := errorMap[code]; ok {
		return message
	}
	return errorMap[ErrorSystemError]
}
