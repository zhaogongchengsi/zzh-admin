package middlewares

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/global"
	"github/gin-react-admin/pkg/errcode"
	"github/gin-react-admin/pkg/response"
	"github/gin-react-admin/utils"
)

// JWTAuth  token 解析中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		authHeader := c.Request.Header.Get("z_token")
		isw := door(path)
		if authHeader == "" && isw == false {
			terr := response.Response{
				State: response.State{
					Code:    errcode.UnauthorizedAuthExist,
					Message: "token 不存在",
				},
			}
			terr.SendError(c)
			c.Abort()
			return
		}
		Claims, err := utils.ParseToken(authHeader)
		if err != 200 && isw == false {
			terr := response.Response{
				State: response.State{
					Code:    err,
					Message: "token 解析失败",
				},
			}
			terr.SendError(c)
			c.Abort()
			return
		}
		c.Set("claims", Claims)
		c.Next()
	}
}

var whiteList = [3]string{"/api/v1/user/register","/api/v1/user/login","/api/v1/verify_code"}
func door(path string) bool {
	var tag bool
	if global.ServerSetting.RenMode == "debug" {
		tag = true
	} else {
		for _, item := range whiteList {
			if item == path {
				tag = true
				break;
			} else {
				tag = false
			}
		}
	}
	return tag
}
