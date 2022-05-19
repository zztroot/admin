package form

type PlatformStatisticsSaveOneForm struct {
	PlatformId int32   `json:"platform_id" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
	Comments   string  `json:"comments"`
	Goods      string  `json:"goods" binding:"required"`
}

type PlatformStatisticsModifyOneForm struct {
	PlatformId int32   `json:"platform_id" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
	SalesTimes int32   `json:"sales_times"`
	Comments   string  `json:"comments"`
	Goods      string  `json:"goods" binding:"required"`
	Id         int32   `json:"id" binding:"required"`
}

type PlatformStatisticsDeleteOneForm struct {
	Id int32 `json:"id" binding:"required"`
}
