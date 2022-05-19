package form

type GoodsSaveOneForm struct {
	Id               int32   `json:"id"`                              // 商品ID
	WarehouseId      int32   `json:"warehouse_id" binding:"required"` // 仓库ID
	Name             string  `json:"name" binding:"required"`         // 商品名称
	Date             string  `json:"date" binding:"required"`         // 销售日期
	Unit             string  `json:"unit" binding:"required"`         // 商品单位
	CostPrice        float32 `json:"cost_price" binding:"required"`   // 商品成本价
	SupplyPrice      float32 `json:"supply_price" binding:"required"` // 供货价格
	SaleNumber       int     `json:"sale_number" binding:"required"`  // 销售数量
	ReturnNumber     int     `json:"return_number" `                  // 退货数量
	LinkReturnNumber int     `json:"link_return_number" `             // 链路客退数量
}

type GoodsModifyOneForm struct {
	Id               int32   `json:"id"  binding:"required"`          // 商品ID
	WarehouseId      int32   `json:"warehouse_id" binding:"required"` // 仓库ID
	Name             string  `json:"name" binding:"required"`         // 商品名称
	Date             string  `json:"date" binding:"required"`         // 销售日期
	Unit             string  `json:"unit" binding:"required"`         // 商品单位
	CostPrice        float32 `json:"cost_price" binding:"required"`   // 商品成本价
	SupplyPrice      float32 `json:"supply_price" binding:"required"` // 供货价格
	SaleNumber       int     `json:"sale_number" `                    // 销售数量
	ReturnNumber     int     `json:"return_number"`                   // 退货数量
	LinkReturnNumber int     `json:"link_return_number"`              // 链路客退数量
}

type GoodsDeleteOneForm struct {
	Id int32 `json:"id"  binding:"required"` // 商品ID
}
