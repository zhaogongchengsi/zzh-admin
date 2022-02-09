package request

type Menus struct {
	ParentId    int    `json:"parent_id"`
	Component   string `json:"component" binding:"required"`
	MenuPath    string `json:"path" binding:"required"`
	MenuName    string `json:"name" binding:"required"`
	MenuIcon    string `json:"icon"`
	MenuDisable bool   `json:"disabled"`
	Label       string `json:"label" binding:"required"`
	Sort        int    `json:"sort" binding:"required"`
	Remarks     string `json:"remarks"`
	Id float64 `json:"ID"`
}

