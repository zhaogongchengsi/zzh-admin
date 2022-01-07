package v1

import (
	"github/gin-react-admin/global"
	"github/gin-react-admin/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

// @Tags CreateVerificationCode
// @Summary 获取验证码
// @Produce  application/json
// @Success 200 {string} string "{"code":"200","data":{ id:"", b64s:"" },"msg":"验证码获取成功"}"
// @Router /api/v1/verify_code [get]
func CreateVerificationCode(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(
		global.Captcha.ImgHeight,
		global.Captcha.ImgWidth,
		global.Captcha.KeyLong,
		global.Captcha.MaxSkew,
		global.Captcha.DotCount,
	)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		var rep = response.Response{
			Code: http.StatusOK,
			Msg:  "验证码获取失败",
			Data: map[string]interface{}{"err": err},
		}
		rep.Send(c)
	} else {
		var repOk = response.Response{
			State: response.State{
				Code:    200,
				Success: true,
				Message: "成功",
			},
			Code: http.StatusOK,
			Msg:  "验证码获取成功",
			Data: map[string]string{"id": id, "b64s": b64s},
		}
		repOk.Send(c)
	}
}
