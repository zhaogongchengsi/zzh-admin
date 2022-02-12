package init

import (
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
)

var modelList = []interface{}{
	model.Menu{},
	model.Article{},
	model.AuthRole{},
	model.File{},
	model.AdminUser{},
}

func AutoMigrate () error {
	var err error = nil
	for _, mod := range modelList {
		err = global.DBEngine.AutoMigrate(&mod)
	}
	return  err
}