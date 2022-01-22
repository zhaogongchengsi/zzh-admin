package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type AdminUser struct {
	gorm.Model
	UUID         uuid.UUID `json:"uuid" gorm:"comment:用户UUID"` // 用户UUID
	UserAdmin    string    `json:"useradmin" gorm:"default:grading;comment:用户账号"`
	Username     string    `json:"userName" gorm:"comment:用户登录名"`                                                                          // 用户登录名
	Password     string    `json:"-"  gorm:"comment:用户登录密码"`                                                                        // 用户登录密码
	NickName     string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                              // 用户昵称
	SideMode     string    `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                                            // 用户侧边主题
	AvatarImg    string    `json:"avatarImg" gorm:"default:https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png;comment:用户头像"` // 用户头像
	BaseColor    string    `json:"baseColor" gorm:"default:#333;comment:基础颜色"`                                                         // 基础颜色
	BaseTextColor string  `json:"baseTextColor" gorm:"default: #fff; comment: 基础字体颜色"`
	ActiveColor  string    `json:"activeColor" gorm:"default:#409EFF;comment:活跃颜色"`                                                      // 活跃颜色
	AuthorityId  string    `json:"authorityId" gorm:"default:1;comment:用户角色ID"`                                                            // 用户角色ID
	ParentNodeId string    `json:"parentnodeid" gorm:"default:1; comment: 用户角色"`
}
