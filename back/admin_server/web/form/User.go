package form

type LoginForm struct {
	UserName     string `json:"user_name" binding:"required"`     // 用户账号
	UserPassword string `json:"user_password" binding:"required"` // 用户密码
}

type SaveUserForm struct {
	UserName      string `json:"user_name" binding:"required"`      // 用户账号
	UserPassword  string `json:"user_password" binding:"required"`  // 用户密码
	UserRealName  string `json:"user_real_name" binding:"required"` // 用户真实名字
	UserPhone     string `json:"user_phone"  binding:"required"`    // 用户手机号
	UserAge       int    `json:"user_age"  binding:"required"`      // 用户年龄
	UserSex       int    `json:"user_sex"  binding:"required"`      // 用户性别
	UserRoleId    int32  `json:"user_role_id"`                      // 用户角色ID
	AgainPassword string `json:"again_password" binding:"required"` // 再次输入密码
}

type ModifyUserForm struct {
	UserName        string `json:"user_name"`             // 用户账号
	UserPassword    string `json:"user_password"`         // 用户旧密码密码
	NewUserPassword string `json:"new_user_password"`     // 新密码
	UserRealName    string `json:"user_real_name" `       // 用户真实名字
	UserPhone       string `json:"user_phone"  `          // 用户手机号
	UserAge         int    `json:"user_age" `             // 用户年龄
	UserSex         int    `json:"user_sex"  `            // 用户性别
	UserRoleId      int32  `json:"user_role_id"`          // 用户角色ID
	Id              int32  `json:"id" binding:"required"` // 用户ID
	AgainPassword   string `json:"again_password"`        // 再次输入密码
}

type GetUserOneForm struct {
	Id int32 `json:"id" binding:"required"` // 用户ID
}
