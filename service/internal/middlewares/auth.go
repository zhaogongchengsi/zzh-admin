package middlewares

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/pkg/errcode"
	"github/gin-react-admin/pkg/response"
	"github/gin-react-admin/utils"
	"time"
)

// JWTAuth  token 解析中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		authHeader := c.Request.Header.Get("z_token")
		isw := door(path)
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
		Claims, err := utils.ParseToken(authHeader) // 解析出token
		// 判断当前token 是否在当前到了缓冲时间内 若是到了缓冲时间内 则更新一个新token 此时 前端会有一个新token 和 老token 若注重安全可以选择讲老token 拉黑
		if Claims.ExpiresAt-time.Now().Unix() < Claims.BufferTime {
			newToken, _ :=  utils.CreateToken(Claims.UUID, Claims.ID, Claims.Username,Claims.NickName)
			c.Header("new-token", newToken)
		}
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


		//CompareTime(Claims.StandardClaims.ExpiresAt)


		c.Set("claims", Claims)
		c.Next()
	}
}

var whiteList = [3]string{"/api/v1/user/register","/api/v1/user/login","/api/v1/verify_code"}
func door(path string) bool {
	var tag bool
	for _, item := range whiteList {
		if item == path {
			tag = true
			break
		} else {
			tag = false
		}
	}
	return tag
}

