package global

import (
	"github/gin-react-admin/pkg/setting"

	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine        *gorm.DB
	Captcha         *setting.Captcha
	Jwt             *setting.JWT
)
