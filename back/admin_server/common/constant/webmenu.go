package constant

// 前端菜单数据
type WebMenu struct {
	Id       int32      `json:"id"`
	Name     string     `json:"name"`
	Path     string     `json:"path"`
	Title    string     `json:"title"`
	Icon     string     `json:"icon"`
	Status   bool       `json:"status"`
	Children []*WebMenu `json:"children"`
	FatherId int32      `json:"is_father"` // 父级ID
}

// 主菜单
var Menu = []WebMenu{
	{
		Id:       1,
		Name:     "home",
		Path:     "/home",
		Title:    "首页",
		Icon:     "HomeFilled",
		Status:   false,
		Children: []*WebMenu{},
		FatherId: 0,
	},
	{
		Id:       2,
		Name:     "goods",
		Path:     "/goods",
		Title:    "商品管理",
		Icon:     "GoodsFilled",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{},
	},
	{
		Id:       3,
		Name:     "platform",
		Path:     "/platform",
		Title:    "平台统计",
		Icon:     "HelpFilled",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{},
	},
	{
		Id:       4,
		Name:     "account",
		Path:     "/account",
		Title:    "应付帐款",
		Icon:     "Stamp",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{},
	},
	{
		Id:       5,
		Name:     "system",
		Path:     "/system",
		Title:    "系统管理",
		Icon:     "Menu",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{},
	},
	{
		Id:       8,
		Name:     "change",
		Path:     "/change",
		Title:    "变动成本",
		Icon:     "TrendCharts",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{},
	},
	{
		Id:       9,
		Name:     "fixed",
		Path:     "/fixed",
		Title:    "固定成本",
		Icon:     "Crop",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{},
	},
}

var ListWebMenu = []WebMenu{
	{
		Id:       1,
		Name:     "home",
		Path:     "/home",
		Title:    "首页",
		Icon:     "HomeFilled",
		Status:   false,
		Children: []*WebMenu{},
		FatherId: 0,
	},
	{
		Id:       2,
		Name:     "goods",
		Path:     "/goods",
		Title:    "商品统计",
		Icon:     "GoodsFilled",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{},
	},
	{
		Id:       3,
		Name:     "platform",
		Path:     "/platform",
		Title:    "平台统计",
		Icon:     "HelpFilled",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{},
	},
	{
		Id:       4,
		Name:     "account",
		Path:     "/account",
		Title:    "应付帐款",
		Icon:     "Stamp",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{},
	},
	{
		Id:       8,
		Name:     "change",
		Path:     "/change",
		Title:    "变动成本",
		Icon:     "TrendCharts",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{},
	},
	{
		Id:       9,
		Name:     "fixed",
		Path:     "/fixed",
		Title:    "固定成本",
		Icon:     "Crop",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{},
	},
	{
		Id:       5,
		Name:     "system",
		Path:     "/system",
		Title:    "系统管理",
		Icon:     "Menu",
		Status:   false,
		FatherId: 0,
		Children: []*WebMenu{
			{
				Id:       6,
				Name:     "user",
				Path:     "/user",
				Title:    "用户管理",
				Icon:     "UserFilled",
				Status:   false,
				FatherId: 5,
				Children: []*WebMenu{},
			},
			{
				Id:       7,
				Name:     "role",
				Path:     "/role",
				Title:    "角色管理",
				Icon:     "Opportunity",
				Status:   false,
				FatherId: 5,
				Children: []*WebMenu{},
			},
		},
	},
}
