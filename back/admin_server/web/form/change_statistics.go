package form

type SaveChangeStatisticsOneForm struct {
	Date string `json:"date" binding:"required"` // 日期
	// 仓费用估算情况（汇总量）
	WarehouseCost          float32 `json:"warehouse_cost"`            // 仓库租金
	HandWarehouseGoodsCost float32 `json:"warehouse_goods_cost"`      // 交仓货费
	BackWarehouseGoodsCost float32 `json:"back_warehouse_goods_cost"` // 退仓货费
	ChangeWarehouseCost    float32 `json:"change_warehouse_cost"`     // 转仓运费
	PackCost               float32 `json:"pack_cost"`                 // 打包/标签
	// 协同仓费用估算情况（汇总量）
	CoordinationWarehouseCost  float32 `json:"coordination_warehouse_cost"`    // 协同仓仓库租金
	OperationAndUnloadingCost  float32 `json:"operation_and_unloading_cost"`   // 收货操作费/收货装卸费
	ReplenishmentCost          float32 `json:"replenishment_cost"`             // 补货费
	ManyBackWarehouseGoodsCost float32 `json:"many_back_warehouse_goods_cost"` // 多货返仓费
	OperationBackGoodsCost     float32 `json:"operation_back_goods_cost"`      // 操作退货费
	ChangeCost                 float32 `json:"change_cost"`                    // 调拨转运费
	UserBackGoodsCost          float32 `json:"user_back_goods_cost"`           // 客退调拨费
	// 当天采购进货费用支出
	DeliveryCost  float32 `json:"delivery_cost"`    // 私人交货费
	LogisticsCost float32 `json:"logistics_cost"`   // 物流费
	GoodsLaLaCost float32 `json:"goods_la_la_cost"` // 货拉拉费
	OtherCost     float32 `json:"other_cost"`       // 设计费等杂费
}

type ModifyChangeStatisticsOneForm struct {
	Id   int32  `json:"id" binding:"required"`   // ID
	Date string `json:"date" binding:"required"` // 日期
	// 仓费用估算情况（汇总量）
	WarehouseCost          float32 `json:"warehouse_cost"`            // 仓库租金
	HandWarehouseGoodsCost float32 `json:"warehouse_goods_cost"`      // 交仓货费
	BackWarehouseGoodsCost float32 `json:"back_warehouse_goods_cost"` // 退仓货费
	ChangeWarehouseCost    float32 `json:"change_warehouse_cost"`     // 转仓运费
	PackCost               float32 `json:"pack_cost"`                 // 打包/标签
	// 协同仓费用估算情况（汇总量）
	CoordinationWarehouseCost  float32 `json:"coordination_warehouse_cost"`    // 协同仓仓库租金
	OperationAndUnloadingCost  float32 `json:"operation_and_unloading_cost"`   // 收货操作费/收货装卸费
	ReplenishmentCost          float32 `json:"replenishment_cost"`             // 补货费
	ManyBackWarehouseGoodsCost float32 `json:"many_back_warehouse_goods_cost"` // 多货返仓费
	OperationBackGoodsCost     float32 `json:"operation_back_goods_cost"`      // 操作退货费
	ChangeCost                 float32 `json:"change_cost"`                    // 调拨转运费
	UserBackGoodsCost          float32 `json:"user_back_goods_cost"`           // 客退调拨费
	// 当天采购进货费用支出
	DeliveryCost  float32 `json:"delivery_cost"`    // 私人交货费
	LogisticsCost float32 `json:"logistics_cost"`   // 物流费
	GoodsLaLaCost float32 `json:"goods_la_la_cost"` // 货拉拉费
	OtherCost     float32 `json:"other_cost"`       // 设计费等杂费
}

type DeleteChangeStatisticsOneForm struct {
	Id int32 `json:"id" binding:"required"` // ID
}
