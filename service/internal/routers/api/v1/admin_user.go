package v1

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/internal/service"
	"github/gin-react-admin/pkg/errcode"
	"github/gin-react-admin/pkg/request"
	"github/gin-react-admin/pkg/response"
	"github/gin-react-admin/utils"
)

// @Tags User
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户账号, 密码, 验证码ID 验证码"
// @Success 200 {string} string "{"state": { code: 0, msg: "成功" },"data":{ token:"" }}"
// @Router /api/v1/user/login [post]
func Login(c *gin.Context) {
	var newUsr request.Login
	if err := c.ShouldBindJSON(&newUsr); err != nil {
		res := response.Response{Err: err}
		res.SendInputParameterError(c)
		return
	}
	if store.Verify(newUsr.CaptchaId, newUsr.Captcha, true) {
		nu := &model.AdminUser{UserAdmin: newUsr.UserAdmin, Username: newUsr.Username, Password: newUsr.Password}
		userReturn, errRe := service.Login(nu)
		if errRe != nil {
			Re := response.Response{State: response.State{
				Code:    errcode.LoginError,
				Message: "账号或验证码错误",
			}, Err: errRe}
			Re.SendError(c)
			return
		} else {
			// 用户验证成功后 签发token
			if errToken := IssueToken(c, userReturn); errToken != nil {
				ReToken := response.Response{State: response.State{
					Code:    errcode.UnauthorizedTokenTimeoutGenerate,
					Message: "token 生成失败",
				}, Err: errToken}
				ReToken.SendError(c)
				return
			}

		}
	} else {
		Recode := response.Response{State: response.State{
			Code:    errcode.VerifyError,
			Message: "验证码错误",
		}}
		Recode.SendError(c)
		return
	}
}

// @Tags User
// @Summary 用户注册
// @Produce  application/json
// @Param data body request.Register true "用户名, 密码, 用户ID"
// @Success 200 {string} string "{"state": { code: 0, msg: "成功" },"data":{ user: {}}}"
// @Router /api/v1/user/register [post]
func Register(c *gin.Context) {
	var newUsr request.Register
	if err := c.ShouldBindJSON(&newUsr); err != nil {
		res := response.Response{Err: err}
		res.Send(c)
		return
	}
	nu := &model.AdminUser{UserAdmin: newUsr.UserAdmin, Username: newUsr.Username, Password: newUsr.Password, NickName: newUsr.NickName, AuthorityId: newUsr.AuthorityId}
	userReturn, errRe := service.Register(*nu)
	if errRe != nil {
		Re := response.Response{Err: errRe, State: response.State{
			Code:    errcode.CreateAuthError,
			Message: "用户创建失败",
		}}
		Re.SendError(c)
		return
	}
	restock := response.Response{Data: map[string]interface{}{"user": userReturn.UserAdmin}} // 注册成功后
	restock.Send(c)
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {

}

// IssueToken 签发token
func IssueToken(c *gin.Context, user *model.AdminUser) error {
	tokenString, err := utils.CreateToken(user.UUID, user.ID, user.Username, user.NickName)
	if err != nil {
		return err
	}
	Re := response.Response{Data: map[string]string{"user": user.UserAdmin, "token": tokenString}}
	Re.Send(c)
	return nil
}

// @Tags Userlist
// @Summary 获取用户列表
// @Produce  application/json
// @Param data body request.Users true "用户账号, 用户ID"
// @Success 200 {string} string "{"state": { code: 0, msg: "成功" },"data":{ user: {}}}"
// @Router /api/v1/user/users [post]
func GetUserList(c *gin.Context) {
	var newUsr request.Users
	if err := c.ShouldBindJSON(&newUsr); err != nil {
		res := response.Response{Err: err}
		res.SendInputParameterError(c)
		return
	}
	nu := &model.AdminUser{UserAdmin: newUsr.UserAdmin, AuthorityId: newUsr.AuthorityId}
	userReturn, errRe := service.GetUsers(*nu)
	if errRe != nil {
		Re := response.Response{Err: errRe, State: response.State{
			Code:    errcode.DataDoesNotExist,
			Message: "用户数据查找失败",
		}}
		Re.SendError(c)
		return
	}
	restock := response.Response{Data: map[string]interface{}{"users": userReturn}}
	restock.Send(c)
}
