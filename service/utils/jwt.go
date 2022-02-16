package utils

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github/gin-react-admin/global"
	"github/gin-react-admin/pkg/errcode"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// var jwtSecret = []byte(global.Jwt.SigningKey)

type Claims struct {
	UUID     uuid.UUID
	ID       uint
	Username string
	NickName string
	jwt.StandardClaims
}

// CreateToken 创建token
func CreateToken(uuid uuid.UUID, id uint, username string, nickname string) (string, error) {
	nowTime := time.Now()                                                      // token 生效时间
	expireTime := nowTime.Add(time.Duration(global.Jwt.ExpiresAt) * time.Hour) // token 有效时间 小时为单位
	claims := Claims{
		UUID:     uuid,
		ID:       id,
		Username: username,
		NickName: nickname,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: global.Jwt.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString([]byte(global.Jwt.SigningKey))
	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, int) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Jwt.SigningKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors & jwt.ValidationErrorMalformed != 0 {
				return nil, errcode.UnauthorizedTokenNotEffective
			} else if ve.Errors & jwt.ValidationErrorExpired != 0 {
				return nil, errcode.UnauthorizedTokenTimeout
			} else if ve.Errors & jwt.ValidationErrorNotValidYet != 0 {
				return nil, errcode.UnauthorizedTokenError
			} else {
				return nil, errcode.UnauthorizedTokenError
			}
		}
	}

	if tokenClaims != nil {
		if claims, token := tokenClaims.Claims.(*Claims); token && tokenClaims.Valid {
			return claims, 200
		}
		return nil, errcode.UnauthorizedTokenError
	} else {
		return nil, errcode.UnauthorizedTokenError
	}
}

func GetHeaderUser(c *gin.Context) (*Claims, int) {
	authHeader := c.Request.Header.Get("z_token")
	return ParseToken(authHeader)
}