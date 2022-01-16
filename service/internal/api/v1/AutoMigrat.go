package v1

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
	"net/http"
)

// AutoMigrated 数据库迁移接口
func AutoMigrated(c *gin.Context) {
	var modelType model.File
	err := global.DBEngine.AutoMigrate(&modelType)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"mag": "数据库迁移成功"})
}
