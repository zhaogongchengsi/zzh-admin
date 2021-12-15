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
		path := c.FullPath() // 此处打算提出来一个保安函数 减查是否要进行token 解析防守
		if path == "/api/v1/user/register" || path == "/api/v1/user/login" || path == "/api/v1/verify_code" {
			c.Next()
			return
		}
		if global.ServerSetting.RenMode == "debug" {
			c.Next()
			return
		}
		authHeader := c.Request.Header.Get("z_token")
		if authHeader == "" {
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
		if err > 300 {
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
