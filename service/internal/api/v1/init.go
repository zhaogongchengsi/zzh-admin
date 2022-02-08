package v1

import (
	"github.com/gin-gonic/gin"
	i "github/gin-react-admin/pkg/init"
	"net/http"
)


type InitService struct {
	DBname string `json:"db_name"`
	DBPort string `json:"db_port"`
	DBtype string `json:"db_type"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Host string `json:"host"`
	TablePrefix string `json:"table_prefix"`
}

// AutoMigrated 数据库迁移接口
func AutoMigrated(c *gin.Context) {
	err := i.AutoMigrate()
	i.Menus()
	user, er := i.UserAdmin()
	if err != nil || er != nil {
		c.JSON(http.StatusOK, gin.H{"mag": "数据库初始化失败", "err": err, "user": user})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mag": "数据库初始化成功"})
}
