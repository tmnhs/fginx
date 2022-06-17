package errcode

//错误码
const (
	SUCCESS = 200
	ERROR   = 1000

	ErrorRequestParameter     = 1001
	ErrorEmailFormat          = 1002
	ErrorLoginByEmail         = 1003
	ErrorLoginByUserName      = 1004
	ErrorLoginCaptcha         = 1005
	ErrorEmailCode            = 1006
	ErrorUserNameFormat       = 1007
	ErrorEmailNotRegister     = 1008
	ErrorEmailSend            = 1009
	ErrorEmailOverdue         = 1010
	ErrorJwtInvalid           = 1011
	ErrorCaptchaGenerate      = 1012
	ErrorUserBan              = 1013
	ErrorTokenGenerate        = 1014
	ErrorLoginStatusSet       = 1015
	ErrorRegisterFormat       = 1016
	ErrorRegister             = 1017
	ErrorUserNameExist        = 1018
	ErrorEmailExist           = 1019
	ErrorChangePasswordFormat = 1020
	ErrorChangePassword       = 1021
	ErrorAddress              = 1022
	ErrorProfileUpdate        = 1023
	ErrorProfileGet           = 1023
	ErrorNotLogin             = 1024
	ErrorTokenExpiration      = 1025
	ErrorLoginLog             = 1026
	ErrorGetAddress           = 1027
	ErrorSetUserStatisInfo    = 1028
	ErrorFileDelete           = 1026
	ErrorFile                 = 1027

	ErrorRedis         = 2001
	ErrorRedisNotFound = 2002

	ErrorHasCollected = 3001
	ErrorIdleGet      = 3002
	ErrorPurchaseGet  = 3003
	ErrorHttpUpgrade  = 3004
	ErrorOrderSelf    = 3005
)
