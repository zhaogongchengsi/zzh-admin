package service

import (
	"errors"
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Login(u *model.AdminUser) (usr *model.AdminUser, err error) {
	var user model.AdminUser
	u.Password = utils.MD5([]byte(u.Password))
	err = global.DBEngine.Where("user_admin = ? AND password = ?", u.UserAdmin, u.Password).First(&user).Error
	return &user, err
}
func Register(u model.AdminUser) (usr model.AdminUser, err error) {
	var user model.AdminUser
	if !errors.Is(global.DBEngine.Where("user_admin = ?", u.UserAdmin).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return u, errors.New("用户名已注册")
	}
	u.UUID = uuid.NewV4()
	u.Password = utils.MD5([]byte(u.Password))
	err = global.DBEngine.Create(&u).Error
	return u, err
}

// func (s *ServiceUser) ChangePassword(user model.AdminUser) (err error, token string, usr model.AdminUser) {

// }

type AdminUserApi struct {
	UUID         uuid.UUID
	UserAdmin    string
	Username     string
	NickName     string
	SideMode     string
	AvatarImg    string
	BaseColor    string
	ActiveColor  string
	AuthorityId  string
	ParentNodeId string
}

// GetUsers 查询自己用户权限下所以的用户
func GetUsers(u model.AdminUser) (users []AdminUserApi, er error) {
	var user []AdminUserApi
	err := global.DBEngine.Model(&model.AdminUser{}).Where("parent_node_id = ?", u.AuthorityId).Find(&user).Error
	return user, err
}
