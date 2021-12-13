package service

import (
	"errors"
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
	"gorm.io/gorm"
)

// CreateRole 创建角色
func CreateRole(a *model.AuthRole) (au *model.AuthRole, err error) {
	var AuthRole model.AuthRole
	if !errors.Is(global.DBEngine.Where("auth_role_id = ?", a.AuthRoleID).First(&AuthRole).Error, gorm.ErrRecordNotFound) {
		return a, errors.New("存在相同角色id")
	}
	err = global.DBEngine.Create(&a).Error
	return a, err
}

// GetRoles 获取角色列表
func GetRoles() ([]model.AuthRole, error) {
	var mode model.AuthRole
	var roles []model.AuthRole
	er := global.DBEngine.Model(&mode).Find(&roles).Error
	return roles, er
}
