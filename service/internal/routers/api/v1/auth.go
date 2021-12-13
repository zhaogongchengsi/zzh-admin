package v1

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/internal/service"
	"github/gin-react-admin/pkg/request"
	"github/gin-react-admin/pkg/response"
)

// @Tags Auth
// @Summary 创建角色
// @Produce  application/json
// @Param data body request.Role true "父角色id， 角色名， 角色备注， 角色ID"
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /api/v1/role/create_role [post]
func CreateRole(c *gin.Context) {
	var newRole request.Role
	if err := c.ShouldBindJSON(&newRole); err != nil {
		res := response.Response{Err: err}
		res.Send(c)
		return
	}
	role := model.AuthRole{
		ParentID:        newRole.ParentID,
		AuthRoleName:    newRole.AuthRoleName,
		AuthRoleRemarks: newRole.AuthRoleRemarks,
		AuthRoleID:      newRole.AuthRoleID,
	}
	r, er := service.CreateRole(&role)
	if er != nil {
		resEr := response.Response{Err: er}
		resEr.SendError(c)
		return
	}
	res := response.Response{Data: r}
	res.Send(c)
}

// @Tags Auth
// @Summary 获取角色列表
// @Produce  application/json
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /api/v1/role/get_roles [get]
func GetAuthLIst(c *gin.Context) {
	r, er := service.GetRoles()
	if er != nil {
		resEr := response.Response{Data: er}
		resEr.Send(c)
		return
	}
	res := response.Response{Data: r}
	res.Send(c)
}
