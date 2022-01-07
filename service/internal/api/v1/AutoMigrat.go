package v1

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
	"net/http"
)

// AutoMigrated 数据库迁移接口
func AutoMigrated(c *gin.Context) {
	var user model.Menu
	err := global.DBEngine.AutoMigrate(&user)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"mag": "数据库迁移成功"})
}
