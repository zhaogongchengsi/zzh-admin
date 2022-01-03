package request

type Menus struct {
	ParentId    int    `json:"parent_id"`
	Component   string `json:"component" binding:"required"`
	MenuPath    string `json:"menu_path" binding:"required"`
	MenuName    string `json:"menu_name" binding:"required"`
	MenuIcon    string `json:"menu_icon"`
	MenuDisable bool   `json:"menu_disable"`
	Label       string `json:"Label" binding:"required"`
	Sort        int    `json:"sort" binding:"required"`
	Remarks     string `json:"remarks"`
	Id float64 `json:"id"`
}

