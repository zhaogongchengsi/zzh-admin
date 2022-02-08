package init

import (
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/internal/service"
)

func UserAdmin () (model.AdminUser, error) {
	nu := &model.AdminUser{
		UserAdmin: "root",
		Username: "admin",
		Password: "12345",
		NickName: "超级管理员（默认）",
		AuthorityId: "1",
	}
	return service.Register(*nu)
}
