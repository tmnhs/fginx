package system

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/model"
	"github.com/tmnhs/fginx/server/internal/model/common/response"
	"github.com/tmnhs/fginx/server/internal/model/system"
	systemReq "github.com/tmnhs/fginx/server/internal/model/system/request"
	systemRes "github.com/tmnhs/fginx/server/internal/model/system/response"
	"github.com/tmnhs/fginx/server/pkg/errcode"
	"github.com/tmnhs/fginx/server/pkg/timer"
	"github.com/tmnhs/fginx/server/pkg/utils"
	"go.uber.org/zap"
	"strings"
)

//通过用户名加密码或者邮箱加密码登录
func (b *BaseApi) Login(c *gin.Context) {
	var req systemReq.Login
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(errcode.ErrorRequestParameter, "请求参数有误", c)
		return
	}
	//两种情况，一种是邮箱加密码登录，另一种是用户名加密码登录
	isEmail := utils.IsEmail(req.Account)
	if isEmail == true {
		//情况一:使用邮箱登录
		if err := utils.Verify(req, utils.LoginByEmailVerify); err != nil {
			response.FailWithMessage(errcode.ErrorEmailFormat, err.Error(), c)
			return
		}
		if store.Verify(req.CaptchaId, req.Captcha, true) {
			u := &system.User{
				Password: req.Password,
				Email:    req.Account,
			}
			if err, user := userService.LoginByUserEmail(u); err != nil {
				global.GV_LOG.Error("登陆失败! 邮箱不存在或者密码错误!", zap.Error(err))
				response.FailWithMessage(errcode.ErrorLoginByEmail, "邮箱不存在或者密码错误", c)
			} else {
				status, err := userService.GetUserBanStatus(user.ID)
				if err != nil || status == nil {
					response.FailWithMessage(errcode.ErrorLoginByEmail, "登录失败", c)
					return
				}
				if *status == model.StatusBan {
					response.FailWithMessage(errcode.ErrorUserBan, "用户已被封禁", c)
					return
				}
				global.GV_LOG.Debug("user login by email and password")
				b.tokenNext(c, *user)
			}
		} else {
			response.FailWithMessage(errcode.ErrorLoginCaptcha, "验证码错误", c)
		}
	} else {
		//情况二:使用用户名登录
		if err := utils.Verify(req, utils.LoginByUserNameVerify); err != nil {
			response.FailWithMessage(errcode.ErrorUserNameFormat, err.Error(), c)
			return
		}
		if store.Verify(req.CaptchaId, req.Captcha, true) {
			u := &system.User{
				Password: req.Password,
				UserName: req.Account,
			}
			if err, user := userService.LoginByUserName(u); err != nil {
				global.GV_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
				response.FailWithMessage(errcode.ErrorLoginByUserName, "用户名不存在或者密码错误", c)
			} else {
				status, err := userService.GetUserBanStatus(user.ID)
				if err != nil || status == nil {
					response.FailWithMessage(errcode.ErrorLoginByEmail, "登录失败", c)
					return
				}
				if *status == model.StatusBan {
					response.FailWithMessage(errcode.ErrorUserBan, "用户已被封禁", c)
					return
				}
				global.GV_LOG.Debug("user login by username and password")
				b.tokenNext(c, *user)
			}
		} else {
			response.FailWithMessage(errcode.ErrorLoginCaptcha, "验证码错误", c)
		}
	}
}

//通过邮箱验证登录
func (b *BaseApi) SendEmail(c *gin.Context) {
	var req systemReq.LoginByEmail
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(errcode.ErrorRequestParameter, "请求参数有误", c)
		return
	}
	user, err := userService.FindUserByEmail(req.Email)
	if err != nil || user == nil {
		response.FailWithMessage(errcode.ErrorEmailNotRegister, "该邮箱还未注册，请先注册", c)
		return
	}
	//生成验证码
	code, err := utils.CreateVerificationCode()
	//将验证码存入redis中
	err = redisService.SetToRedis(context.Background(), fmt.Sprintf(model.KeyEmailCode, user.ID, user.Email), code, timer.CodeExpirationTime)
	if err != nil {
		response.FailWithMessage(errcode.ErrorRedis, "发送邮箱失败，请重新发送", c)
		return
	}
	mailTo := []string{
		req.Email,
	}
	//邮件主题为"Hello"
	subject := "hust_mall客服"
	// 邮件正文
	content, err := utils.ReadFile(utils.VerificationCodeFile)
	content = strings.Replace(content, "codeSlot", code, 1)
	content = strings.Replace(content, "emailSlot", req.Email, 1)
	body := strings.Replace(content, "userSlot", user.UserName, 1)

	//发送邮件
	err = utils.SendMail(mailTo, subject, body)
	if err != nil {
		response.FailWithMessage(errcode.ErrorEmailSend, "发送邮箱失败，请重新发送", c)
		return
	}
	response.OkWithMessage("发送成功！", c)
}

//通过邮箱验证登录
func (b *BaseApi) LoginByEmail(c *gin.Context) {
	var req systemReq.LoginByEmail
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(errcode.ErrorRequestParameter, "请求参数有误", c)
		return
	}
	user, err := userService.FindUserByEmail(req.Email)
	if err != nil || user == nil {
		response.FailWithMessage(errcode.ErrorLoginByEmail, "登录失败！", c)
		return
	}
	status, err := userService.GetUserBanStatus(user.ID)
	if err != nil || status == nil {
		response.FailWithMessage(errcode.ErrorLoginByEmail, "登录失败", c)
		return
	}
	if *status == model.StatusBan {
		response.FailWithMessage(errcode.ErrorUserBan, "用户已被封禁", c)
		return
	}
	var code string
	err = redisService.GetFromRedis(context.Background(), fmt.Sprintf(model.KeyEmailCode, user.ID, user.Email), &code)
	if err == model.ErrRedisNotFound {
		response.FailWithMessage(errcode.ErrorEmailOverdue, "验证码已失效！", c)
		return
	}
	if err != nil {
		response.FailWithMessage(errcode.ErrorRedis, "登录失败！", c)
		return
	}
	if code != req.Code {
		response.FailWithMessage(errcode.ErrorEmailCode, "验证码错误！", c)
		return
	}

	b.tokenNext(c, *user)
}

// 登录以后签发jwt
func (b *BaseApi) tokenNext(c *gin.Context, user system.User) {
	j := &utils.JWT{SigningKey: []byte(global.GV_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GV_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage(errcode.ErrorTokenGenerate, "获取token失败", c)
		return
	}
	//todo
	err = redisService.SetToRedis(context.Background(), fmt.Sprintf(model.KeyToken, user.ID), token, 0)
	//err = redisService.SetToRedis(context.Background(), fmt.Sprintf(model.KeyToken, user.ID), token, timer.UserTokenExpireTime)
	if err != nil {
		global.GV_LOG.Error("存储token失败!", zap.Error(err))
		response.FailWithMessage(errcode.ErrorRedis, "设置token失败", c)
		return
	}
	ip := c.ClientIP()
	log := system.UserLoginLog{
		UID:       user.ID,
		UserName:  user.UserName,
		Ip:        ip,
		LoginType: model.UserLogin,
		City:      *utils.GetIpAddress(ip),
		LoginTime: timer.GetNowUnix(),
	}
	err = userService.SetUserLoginLog(log)
	if err != nil {
		global.GV_LOG.Error("存储token失败!", zap.Error(err))
		response.FailWithMessage(errcode.ErrorLoginLog, "设置token失败", c)
		return
	}
	/*//设置用户在线状态
	err = userService.SetUserOnline(context.Background(), claims.ID, model.UserOnline)
	if err != nil {
		global.GV_LOG.Error("SetUserOnline err:", zap.Error(err))
		response.FailWithDetailed(errcode.ERROR, gin.H{"reload": true}, "未登录或非法访问", c)
		c.Abort()
		return
	}*/
	//登录次数+1
	err = userService.IncUserStatisCount(user.ID, model.IncLoginCount)
	if err != nil {
		global.GV_LOG.Error("IncUserStatisCount err:", zap.Error(err))
		response.FailWithMessage(errcode.ERROR, "登录失败", c)
		return
	}
	if !global.GV_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			TokenID:   fmt.Sprintf(model.KeyToken, user.ID),
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	if err, jwtStr := jwtService.GetRedisJWT(user.UserName); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.UserName); err != nil {
			global.GV_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage(errcode.ErrorLoginStatusSet, "设置登录状态失败", c)
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			TokenID:   fmt.Sprintf(model.KeyToken, user.ID),
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GV_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage(errcode.ErrorLoginStatusSet, "设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JoinInBlacklist(blackJWT); err != nil {
			response.FailWithMessage(errcode.ErrorJwtInvalid, "jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.UserName); err != nil {
			response.FailWithMessage(errcode.ErrorLoginStatusSet, "设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			TokenID:   fmt.Sprintf(model.KeyToken, user.ID),
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

//用户注册
func (b *BaseApi) Register(c *gin.Context) {
	var req systemReq.Register
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(errcode.ErrorRequestParameter, "请求参数有误", c)
		return
	}
	if err := utils.Verify(req, utils.RegisterVerify); err != nil {
		response.FailWithMessage(errcode.ErrorRegisterFormat, err.Error(), c)
		return
	}
	if userService.CheckUserNameExist(req.UserName) == true {
		response.FailWithMessage(errcode.ErrorUserNameExist, "用户名已存在", c)
		return
	}
	if userService.CheckEmailExist(req.Email) == true {
		response.FailWithMessage(errcode.ErrorEmailExist, "邮箱已注册", c)
		return
	}

	user := &system.User{UserName: req.UserName, Password: req.Password, Email: req.Email}
	err, userReturn := userService.Register(*user)
	if err != nil {
		global.GV_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(errcode.ErrorRegister, systemRes.UserResponse{User: userReturn}, "注册失败", c)
		return
	}

	profile := &system.Profile{UID: userReturn.ID, UserName: userReturn.UserName, Email: userReturn.Email}
	err = userService.SetProfile(*profile)
	if err != nil {
		global.GV_LOG.Error("SetProfile err:", zap.Error(err))
		response.FailWithDetailed(errcode.ErrorRegister, systemRes.UserResponse{User: userReturn}, "注册失败", c)
		return
	}
	statis := system.UserStatisInfo{
		UID:        userReturn.ID,
		OrderCount: 0,
		LoginCount: 0,
	}
	err = userService.SetUserStatisInfo(statis)
	if err != nil {
		global.GV_LOG.Error("SetUserStatisInfo err:", zap.Error(err))
		response.FailWithDetailed(errcode.ErrorSetUserStatisInfo, systemRes.UserResponse{User: userReturn}, "注册失败", c)
		return
	}

	response.OkWithDetailed(systemRes.UserResponse{User: userReturn}, "注册成功", c)
}

//修改密码  ChangePassword
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordStruct
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(errcode.ErrorRequestParameter, "请求参数有误", c)
		return
	}
	user := utils.GetUserInfo(c)

	if err := utils.Verify(req, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(errcode.ErrorChangePasswordFormat, err.Error(), c)
		return
	}
	u := &system.User{ID: user.ID, Password: req.Password}
	if err, _ := userService.ChangePassword(u, req.NewPassword); err != nil {
		global.GV_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(errcode.ErrorChangePassword, "修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

//更新用户详细信息 UpdateProfile
func (b *BaseApi) UpdateProfile(c *gin.Context) {
	var profile systemReq.Profile
	user := utils.GetUserInfo(c)

	err := c.ShouldBindJSON(&profile)
	if err != nil {
		response.FailWithMessage(errcode.ErrorRequestParameter, "请求参数有误", c)
		return
	}

	profileUpdate := system.ProfileUpdate{
		UID:      user.ID,
		UserName: profile.UserName,
		Email:    profile.Email,
		Gender:   profile.Gender,
		Phone:    profile.Phone,
		QQ:       profile.QQ,
		Wechat:   profile.Wechat,
		Avatar:   profile.Avatar,
		BgImage:  profile.BgImage,
	}
	err = userService.UpdateProfile(profileUpdate)
	if err != nil {
		response.FailWithMessage(errcode.ErrorProfileUpdate, "更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

//获取用户详细信息 GetProfile
func (b *BaseApi) GetProfile(c *gin.Context) {
	user := utils.GetUserInfo(c)

	profile, err := userService.GetProfile(user.ID)
	if err != nil {
		response.FailWithMessage(errcode.ErrorProfileGet, "获取失败", c)
		return
	}

	response.OkWithDetailed(systemRes.ProfileResponse{
		Profile: profile,
	}, "获取成功", c)
}

//退出登录 LoginOut
func (b *BaseApi) LoginOut(c *gin.Context) {
	user := utils.GetUserInfo(c)

	err := redisService.DelFromRedis(context.Background(), fmt.Sprintf(model.KeyProfile, user.ID))
	if err != nil {
		response.FailWithMessage(errcode.ErrorRedis, "退出登录失败", c)
		return
	}
	err = redisService.DelFromRedis(context.Background(), fmt.Sprintf(model.KeyOnline, user.ID))
	if err != nil {
		response.FailWithMessage(errcode.ErrorRedis, "退出登录失败", c)
		return
	}
	response.OkWithMessage("退出成功", c)
}

//获取用户数据
func (b *BaseApi) GetUserStatis(c *gin.Context) {
	user := utils.GetUserInfo(c)
	var resultStatis systemRes.UserStatis

	userStatis, err := userService.GetUserStatisInfo(user.ID)
	if err != nil {
		response.FailWithMessage(errcode.ErrorRedis, "获取失败", c)
		return
	}
	resultStatis.FinishOrderCount = userStatis.OrderCount

	_ = redisService.GetFromRedis(context.Background(), fmt.Sprintf(model.KeyLastLogin, user.ID), &resultStatis.LastTimestamp)

	err = redisService.GetFromRedis(context.Background(), fmt.Sprintf(model.KeyOnline, user.ID), &resultStatis.Status)

	if err == model.ErrRedisNotFound {
		//不在线
		resultStatis.Status = model.UserNotOnline
		response.OkWithDetailed(resultStatis, "获取成功", c)
		return
	}
	if err != nil {
		response.FailWithMessage(errcode.ErrorRedis, "获取失败", c)
		return
	}
	//在线
	response.OkWithDetailed(resultStatis, "获取成功", c)
}
