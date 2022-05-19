package form

type WarehouseSaveOneForm struct {
	Name string `json:"name" binding:"required"` // 仓库名称
}

type WarehouseModifyOneForm struct {
	Id   int32  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"` // 仓库名称
}
