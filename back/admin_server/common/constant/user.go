package constant

const (
	// 用户是否删除
	UserIsDelNo  = 1
	UserIsDelYes = 2

	// 用户性别
	UserSexMale   = 1 // 男
	UserSexFemale = 2 // 女
)

// token位数
const (
	TokenGenerate = 48
)

// 性别映射表
var SexMap = map[int]string{
	UserSexMale:   "男",
	UserSexFemale: "女",
}
