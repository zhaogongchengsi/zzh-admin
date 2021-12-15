package errcode

const (
	Success                          = 200 // 成功
	ServerError                      = 500 // 服务器内部错误
	InvalidParams                    = 226 // 入参错误
	NotFound                         = 404 // 找不到
	UnauthorizedAuthExist            = 301 // 鉴权失败 没有token
	UnauthorizedTokenError           = 302 // 鉴权失败 token 错误
	UnauthorizedTokenNotEffective    = 303 // 鉴权失败 token 未生效
	UnauthorizedTokenTimeout         = 304 // 鉴权失败 token 超时
	UnauthorizedTokenTimeoutGenerate = 305 // 鉴权失败 token 生成失败
	TooManyRequests                  = 350 // 请求过多
	LoginError                       = 204 // 登录失败 账号密码 错误
	VerifyError                      = 205 // 验证码错误
	UserAlreadyExists                = 206 // 用户已存在
	DataDoesNotExist                 = 504 // 数据查找失败
	CreateAuthError                  = 405 // 角色创建失败
	CreateMenuError                  = 406 // 角色创建失败
)
