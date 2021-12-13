package model

import "gorm.io/gorm"

type AuthRole struct {
	gorm.Model
	ParentID        int    `gorm:"comment:父角色ID"`
	AuthRoleID      int    `gorm:"not null;unique;primary_key;comment:角色ID"`
	AuthRoleName    string `gorm:"comment:角色名"`
	AuthRoleRemarks string `gorm:"comment: 角色描述"`
	AuthRoleMenus   []Menu `gorm:"comment: 角色拥有的功能权限; many2many:auto_role_menus;"`
}
