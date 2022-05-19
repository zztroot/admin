package form

type ModifyRoleForm struct {
	RoleName   string `json:"role_name" binding:"required"`   // 用户账号
	RoleExtend string `json:"role_extend" binding:"required"` // 用户密码
	Id         int32  `json:"id"`                             // ID
}

type DeleteRoleForm struct {
	Id int32 `json:"id" binding:"required"` // ID
}

type GetOneByNameForm struct {
	Name string `json:"name" binding:"required"` // 名称
}
