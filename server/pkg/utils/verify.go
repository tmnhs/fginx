package utils

import "regexp"

//常用正则表达式
type Regexp struct{}

var regexpImpl Regexp

var (
	IdVerify               = Rules{"ID": {NotEmpty()}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	LoginByUserNameVerify  = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Account": {NotEmpty()}, "Password": {Gt("5")}}
	LoginByEmailVerify     = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Account": {RegexpMatch(regexpImpl.Email())}, "Password": {Gt("5")}}
	RegisterVerify         = Rules{"UserName": {NotEmpty()}, "Password": {Gt("5")}, "Email": {RegexpMatch(regexpImpl.Email())}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}, "ParentId": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Password": {NotEmpty()}, "NewPassword": {Gt("5")}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}
)

//邮箱 Email
func (regexp *Regexp) Email() string {
	return "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
}

//电话号码 Phone
func (regexp *Regexp) Phone() string {
	return "^(\\+?0?86\\-?)?1[3-9]\\d{9}$"
}

//腾讯qq QQ 腾讯QQ号从10000开始)
func (regexp *Regexp) QQ() string {
	return "[1-9][0-9]{4,10}"
}

//ip地址 IP
func (regexp *Regexp) IP() string {
	return "\\d+\\.\\d+\\.\\d+\\.\\d+"
}

//密码 Password  由数字和字母组成，并且要同时含有数字和字母，且长度要在6-16位之间。
func (regexp *Regexp) Password() string {
	return "^[0-9a-zA-Z_]{6,16}$"
}
func IsEmail(matchStr string) bool {
	return regexp.MustCompile(regexpImpl.Email()).MatchString(matchStr)
}
