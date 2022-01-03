package v1

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/internal/service"
	"github/gin-react-admin/pkg/errcode"
	"github/gin-react-admin/pkg/request"
	"github/gin-react-admin/pkg/response"
)

// @Tags Menu
// @Summary 创建菜单
// @Produce  application/json
// @Param data body request.Menus true "是否开启， 菜单名， 菜单文件路径， 菜单路径， icon图标， 排序标记，展示名， 父级ID"
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /api/v1/role/create_role [post]
func CreateMenu(c *gin.Context) {
	var newMenu request.Menus
	err := c.ShouldBindJSON(&newMenu)
	if  err != nil {
		res := response.Response{Err: err}
		res.SendInputParameterError(c)
		return
	}
	menuW := &model.Menu{
		Disabled:  newMenu.MenuDisable,
		Name:      newMenu.MenuName,
		Path:      newMenu.MenuPath,
		Component: newMenu.Component,
		Icon:      newMenu.MenuIcon,
		Sort:      newMenu.Sort,
		Label:     newMenu.Label,
		Remarks:   newMenu.Remarks,
		ParentId:  newMenu.ParentId,
	}
	menuItem, errMe := service.CreateMenu(menuW)
	if errMe != nil {
		Re := response.Response{Err: errMe, State: response.State{
			Code: errcode.CreateMenuError,
			Message: "角色创建失败",
		}}
		Re.SendError(c)
		return
	} else {
		ReToken := response.Response{Data: menuItem}
		ReToken.Send(c)
		return
	}
}

// @Tags Menu
// @Summary 获取菜单列表
// @Produce  application/json
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /api/v1/menu/get_menu [get]
func GetMenus(c *gin.Context) {
	menus, errMe := service.GetMenus()
	if errMe != nil {
		Re := response.Response{Data: errMe, State: response.State{
			Code: errcode.NotFound,
			Message: "菜单数据获取失败",
		}}
		Re.SendError(c)
		return
	} else {
		ReToken := response.Response{Data: menus}
		ReToken.Send(c)
		return
	}
}


// @Tags Menu
// @Summary 删除菜单
// @Produce  application/json
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /api/v1/menu/get_menu [post]
func DeleteMenu(c *gin.Context) {
	var menuId request.GetGormById
	_ = c.ShouldBindJSON(&menuId)
	err := service.DeleteMenuHandle(menuId.ID)
	if err != nil {
		deErr := response.Response{Err: err, State: response.State{
			Code: errcode.DeleteError,
			Message: "删除失败",
		}}
		deErr.SendError(c)
		return
	}
	deOk := response.Response{}
	deOk.Send(c)
}

// @Tags Menu
// @Summary 根据ID查找菜单
// @Produce  application/json
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /api/v1/menu/find_menu [post]
func GetMenuByID(c *gin.Context)  {
	var menuId request.GetGormById
	_ = c.ShouldBindJSON(&menuId)
	var menu = model.Menu{

	}
	err := service.GetMenuByID(menuId.ID, &menu)
	if err != nil {
		deErr := response.Response{Err: err, State: response.State{
			Code: errcode.NotFound,
			Message: "查找失败",
		}}
		deErr.SendError(c)
	}
	deOk := response.Response{
		Data: menu,
	}
	deOk.Send(c)
}
