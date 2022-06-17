package request

// User register structure
type Register struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User login structure
type Login struct {
	Account   string `json:"account"`   //账户：用户名或者邮箱
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type LoginByEmail struct {
	Email string `json:"email" form:"email"`
	Code  string `json:"code" form:"code"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}
