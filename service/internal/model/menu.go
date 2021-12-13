package model

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Disabled  bool   `gorm:"default:false;comment:是否禁用"`
	Label     string `gorm:"comment:标签"`
	Name      string `gorm:"comment:菜单名字"`
	Path      string `gorm:"comment:路由路径"`
	Component string `gorm:"comment:前端对应的组件"`
	Icon      string `gorm:"comment:菜单图标"`
	Sort      int    `gorm:"comment:排序标记"`
	ParentId  int    `gorm:"comment:父节点路径"`
	Remarks   string `gorm:"comment:备注"`
}
