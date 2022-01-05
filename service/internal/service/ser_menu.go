package service

import (
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/pkg/request"
	"gorm.io/gorm"
)

func CreateMenu(m *model.Menu) (me *model.Menu, err error) {
	//var menus model.Menu
	// 节点不能重复 需要做一层判断
	//if !errors.Is(global.DBEngine.Where("root = ?", m.Root).First(&menus).Error, gorm.ErrRecordNotFound) {
	//	return &menus, errors.New("根节点")
	//}
	err = global.DBEngine.Create(&m).Error
	return m, err
}

func GetMenus() (me []model.Menu, err error) {
	var m model.Menu
	var menus []model.Menu
	er := global.DBEngine.Model(&m).Find(&menus).Error
	return menus, er
}

func DeleteMenuHandle(id float64) (err error)  {
	err = global.DBEngine.Delete(&model.Menu{}, id).Error
	return err
}

func GetMenuByID (id float64, menu *model.Menu) (err error) {
	err = global.DBEngine.First(&menu, id).Error
	return err
}

func UpdateMenu (menu request.Menus) (err error) {
	var oldMenu model.Menu
	var newMenu = make(map[string]interface{})
	newMenu["parent_id"] = menu.ParentId
	newMenu["path"] = menu.MenuPath
	newMenu["name"] = menu.MenuName
	newMenu["component"] = menu.Component
	newMenu["icon"] = menu.MenuIcon
	newMenu["sort"] = menu.Sort
	newMenu["disabled"] = menu.MenuDisable
	newMenu["label"] = menu.Label
	newMenu["remarks"] = menu.Remarks
	// 开启事务 更新菜单项
	err = global.DBEngine.Transaction(func(tx *gorm.DB) error {
		upErr := tx.Where("id = ?", menu.Id).Find(&oldMenu).Updates(newMenu).Error
		if upErr != nil {
			return upErr
		}
		return nil
	})
	return  err
}



func GetSubMenu (id float64) (menus  []model.Menu,err error) {
	var subMenus []model.Menu
	err = global.DBEngine.Where("parent_id = ?", id).Find(&subMenus).Error
	return  subMenus, err
}

func DeletMenu(id float64) error {
	var oldMenu model.Menu
	menuLIst, suberr := GetSubMenu(id)
	if suberr != nil {
		return  suberr
	}
	if len(menuLIst) > 0 {
		err := global.DBEngine.Where("parent_id = ?", id).First(&oldMenu).Delete(&oldMenu).Error
		return err
	} else {
		err := global.DBEngine.Where("id = ?", id).First(&oldMenu).Delete(&oldMenu).Error
		return err
	}
}