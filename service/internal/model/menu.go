package model

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model `json:"-"`
	Disabled  bool   `gorm:"default:false;comment:是否禁用" json:"disabled"`
	Label     string `gorm:"comment:标签" json:"label"`
	Name      string `gorm:"comment:菜单名字" json:"name"`
	Path      string `gorm:"comment:路由路径" json:"path"`
	Component string `gorm:"comment:前端对应的组件" json:"component"`
	Icon      string `gorm:"comment:菜单图标" json:"icon"`
	Sort      int    `gorm:"comment:排序标记" json:"sort"`
	ParentId  int    `gorm:"comment:父节点路径" json:"parent_id"`
	Remarks   string `gorm:"comment:备注" json:"remarks"`
}
