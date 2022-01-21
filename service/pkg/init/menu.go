package init

import (
	"github/gin-react-admin/internal/model"
	"gorm.io/gorm"
	"time"
)

//Disabled  bool   `gorm:"default:false;comment:是否禁用"`
//Label     string `gorm:"comment:标签"`
//Name      string `gorm:"comment:菜单名字"`
//Path      string `gorm:"comment:路由路径"`
//Component string `gorm:"comment:前端对应的组件"`
//Icon      string `gorm:"comment:菜单图标"`
//Sort      int    `gorm:"comment:排序标记"`
//ParentId  int    `gorm:"comment:父节点路径"`
//Remarks   string `gorm:"comment:备注"`

var InitMenu = []model.Menu{
	{
		Model:gorm.Model{
			ID: 1,
			CreatedAt: time.Now(),
			UpdatedAt:time.Now(),
		},
		Label: "系统管理",
		Name: "system",
		Path: "/system",
		Component: "views/system/index",
		Icon: "ServerIcon",
		Sort: 1,
		ParentId: 0,
		Disabled: false,
	},
	{
		Model:gorm.Model{
			ID: 2,
			CreatedAt: time.Now(),
			UpdatedAt:time.Now(),
		},
		Label: "权限管理",
		Name: "admin",
		Path: "/system/admin",
		Component: "views/system/admin",
		Icon: "KeyIcon",
		Sort: 3,
		ParentId: 1,
		Disabled: false,
	},
	{
		Model:gorm.Model{
			ID: 3,
			CreatedAt: time.Now(),
			UpdatedAt:time.Now(),
		},
		Label: "用户管理",
		Name: "admin_user",
		Path: "/system/user",
		Component: "views/system/admin/Users",
		Icon: "UsersIcon",
		Sort: 4,
		ParentId: 2,
		Disabled: false,
	},
	{
		Model:gorm.Model{
			ID: 4,
			CreatedAt: time.Now(),
			UpdatedAt:time.Now(),
		},
		Label: "角色管理",
		Name: "admin_role",
		Path: "/system/user",
		Component: "views/system/admin/Role",
		Icon: "ChartSquareBarIcon",
		Sort: 6,
		ParentId: 2,
		Disabled: false,
	},
	{
		Model:gorm.Model{
			ID: 5,
			CreatedAt: time.Now(),
			UpdatedAt:time.Now(),
		},
		Label: "菜单管理",
		Name: "admin_menus",
		Path: "/system/menus",
		Component: "views/system/admin/Menus",
		Icon: "FingerPrintIcon",
		Sort: 7,
		ParentId: 2,
		Disabled: false,
	},
	{
		Model:gorm.Model{
			ID: 6,
			CreatedAt: time.Now(),
			UpdatedAt:time.Now(),
		},
		Label: "网站管理",
		Name: "web",
		Path: "/web",
		Component: "views/web/index",
		Icon: "ChartBarIcon",
		Sort: 1,
		ParentId: 0,
		Disabled: false,
	},
}

func InItMenu () {

}