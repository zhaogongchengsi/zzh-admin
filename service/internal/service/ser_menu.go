package service

import (
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
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