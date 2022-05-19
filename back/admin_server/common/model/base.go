package model

type Base struct {
	Id         int32  `json:"id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	IsDelete   int    `json:"is_delete"` // 是否删除 1未删除|2删除
}

func GetObjects() []interface{} {
	return []interface{}{User{}, Role{}, Goods{}, Warehouse{}, PlatformStatistics{}, Platform{}}
}
