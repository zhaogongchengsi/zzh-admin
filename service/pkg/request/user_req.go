package request

type Register struct {
	UserAdmin   string `json:"useradmin" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NickName    string `json:"nickname"`
	HeaderImg   string `json:"headerImg"`
	AuthorityId string `json:"authorityId"`
}

type Login struct {
	UserAdmin string `json:"useradmin" binding:"required"`
	Username  string `json:"username"`                     // 用户名
	Password  string `json:"password" binding:"required"`  // 密码
	Captcha   string `json:"captcha" binding:"required"`   // 验证码
	CaptchaId string `json:"captchaId" binding:"required"` // 验证码ID
}

type ChangePasswordStruct struct {
	UserAdmin   string `json:"useradmin" binding:"required"`
	Username    string `json:"username" binding:"required"`    // 用户名
	Password    string `json:"password" binding:"required"`    // 密码
	NewPassword string `json:"newPassword" binding:"required"` // 新密码
}

type Users struct {
	UserAdmin   string `json:"useradmin"`
	AuthorityId string `json:"authorityId"  binding:"required"`
}

//type MenuNode struct {
//	 Key string `json:"key"`
//	 Title string `json:"title"`
//	 Children []MenuNode `json:"children"`
//	 Label string `json:"label"`
//}
