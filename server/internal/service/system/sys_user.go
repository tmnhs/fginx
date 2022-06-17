package system

import (
	"context"
	"fmt"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/model"
	"github.com/tmnhs/fginx/server/internal/model/system"
	redisClient "github.com/tmnhs/fginx/server/internal/service/redis"
	"github.com/tmnhs/fginx/server/pkg/timer"
	"github.com/tmnhs/fginx/server/pkg/utils"
	"go.uber.org/zap"
	"time"
)

//@function: Register
//@description: 用户注册
//@param: u model.User
//@return: err error, userInter model.User

type UserService struct{}

var redisService = redisClient.RedisService{}

func (userService *UserService) Register(u system.User) (error, *system.User) {
	if nil == global.GV_DB {
		return fmt.Errorf("db not init"), nil
	}
	u.CreatedTime = timer.GetNowUnix()
	// 否则 加密 注册
	u.Password = utils.ScryptPaassword(u.Password)
	err := global.GV_DB.Table(model.UserTableName).Create(&u).Error
	if err != nil {
		global.GV_LOG.Error("Register err:", zap.Error(err))
		return err, nil
	}
	return nil, &u
}

func (userService *UserService) LoginByUserName(u *system.User) (error, *system.User) {
	if nil == global.GV_DB {
		return fmt.Errorf("db not init"), nil
	}

	var user system.User
	u.Password = utils.ScryptPaassword(u.Password)
	err := global.GV_DB.Table(model.UserTableName).Where("username = ? AND password = ?", u.UserName, u.Password).First(&user).Error
	if err != nil {
		global.GV_LOG.Error("LoginByUserName err:", zap.Error(err))
		return err, nil
	}
	return nil, &user
}
func (userService *UserService) LoginByUserEmail(u *system.User) (error, *system.User) {
	if nil == global.GV_DB {
		return fmt.Errorf("db not init"), nil
	}

	var user system.User
	u.Password = utils.ScryptPaassword(u.Password)
	err := global.GV_DB.Table(model.UserTableName).Where("email = ? AND password = ?", u.Email, u.Password).First(&user).Error
	if err != nil {
		global.GV_LOG.Error("LoginByUserEmail err:", zap.Error(err))
		return err, nil
	}
	return nil, &user
}

func (userService *UserService) ChangePassword(u *system.User, newPassword string) (error, *system.User) {
	var user system.User

	u.Password = utils.ScryptPaassword(u.Password)

	err := global.GV_DB.Table(model.UserTableName).Where("id = ? AND password = ?", u.ID, u.Password).First(&user).Update("password", utils.ScryptPaassword(newPassword)).Error
	if err != nil {
		global.GV_LOG.Error("ChangePassword err:", zap.Error(err))
		return err, nil
	}
	return err, u
}

//通过邮箱获取用户 FindUserByEmail
func (userService *UserService) FindUserByEmail(email string) (*system.User, error) {
	var user system.User
	err := global.GV_DB.Table(model.UserTableName).Where("email = ?", email).First(&user).Error
	if err != nil {
		global.GV_LOG.Error("FindUserByEmail err:", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

//获取封禁状态 GetUserBanStatus
func (userService *UserService) GetUserBanStatus(uid int64) (*int, error) {
	var status int
	err := global.GV_DB.Table(model.ProfileTableName).Select("status").Where("uid = ?", uid).Find(&status).Error
	if err != nil {
		global.GV_LOG.Error("GetUserBanStatus err:", zap.Error(err))
		return nil, err
	}
	return &status, nil
}

//获取用户具体信息 GetProfile
func (userService *UserService) GetProfile(uid int64) (*system.Profile, error) {
	var profile system.Profile
	//get from Redis
	err := redisService.GetFromRedis(context.Background(), fmt.Sprintf(model.KeyProfile, uid), &profile)
	if err == nil { // Redis Found
		global.GV_LOG.Debug("GetProfile by redis:", zap.Any("profile", profile))
		return &profile, nil
	}
	if err != model.ErrRedisNotFound { // Redis Error
		return nil, err
	}
	//Redis not Found
	err = global.GV_DB.Table(model.ProfileTableName).Where("uid = ?", uid).First(&profile).Error
	if err != nil {
		global.GV_LOG.Error("GetProfile err:", zap.Error(err))
		return nil, err
	}
	//set to Redis
	err = redisService.SetToRedis(context.Background(), fmt.Sprintf(model.KeyProfile, profile.UID), profile, 0)
	if err != nil {
		global.GV_LOG.Error("GetProfile redis  err:", zap.Error(err))
		return nil, err
	}
	global.GV_LOG.Debug("GetProfile set to redis:", zap.Any("profile", profile))
	return &profile, nil
}

//设置用户的具体信息  SetProfile
func (userService *UserService) SetProfile(profile system.Profile) error {
	profile.CreatedTime = timer.GetNowUnix()
	profile.Status = model.StatusNormal
	//set to Redis
	err := redisService.SetToRedis(context.Background(), fmt.Sprintf(model.KeyProfile, profile.UID), profile, 0)
	if err != nil {
		global.GV_LOG.Error("SetProfile redis  err:", zap.Error(err))
		return err
	}
	//set to Mysql
	err = global.GV_DB.Table(model.ProfileTableName).Create(&profile).Error
	if err != nil {
		global.GV_LOG.Error("SetProfile mysql err:", zap.Error(err))
		return err
	}
	return nil
}

//设置用户的具体信息  SetProfile
func (userService *UserService) UpdateProfile(profile system.ProfileUpdate) error {
	profile.UpdatedTime = timer.GetNowUnix()

	err := global.GV_DB.Table(model.ProfileTableName).Where("uid = ?", profile.UID).Save(&profile).Error
	if err != nil {
		global.GV_LOG.Error("UpdateProfile err:", zap.Error(err))
		return err
	}
	//update to Redis
	err = redisService.SetToRedis(context.Background(), fmt.Sprintf(model.KeyProfile, profile.UID), profile, 0)
	if err != nil {
		global.GV_LOG.Error("UpdateProfile redis  err:", zap.Error(err))
		return err
	}
	return nil
}

//检验用户名是否存在  CheckUserNameExist
func (userService *UserService) CheckUserNameExist(username string) bool {
	var user system.User

	err := global.GV_DB.Table(model.UserTableName).Where("username = ?", username).First(&user).Error
	if err != nil { //mysql没有找到记录也会报错
		return false
	}
	if user.ID > 0 {
		return true
	}
	return false
}

//检验用邮箱是否存在  CheckEmailExist
func (userService *UserService) CheckEmailExist(email string) bool {
	var user system.User

	err := global.GV_DB.Table(model.UserTableName).Where("email = ?", email).First(&user).Error
	if err != nil {
		return false
	}
	if user.ID > 0 {
		return true
	}
	return false
}

//设置用户在线状态
func (userService *UserService) SetUserOnline(ctx context.Context, uid int64, status int) error {
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	key := fmt.Sprintf(model.KeyOnline, uid)
	err := redisService.SetToRedis(ctx, key, status, timer.UserOnlineExpireTime)
	if err != nil {
		global.GV_LOG.Error("SetUserOnline redis  err:", zap.Error(err))
		return err
	}

	return nil
}

//设置用户登录日志
func (userService *UserService) SetUserLoginLog(log system.UserLoginLog) error {
	err := global.GV_DB.Table(model.LoginLogTableName).Create(&log).Error
	if err != nil {
		global.GV_LOG.Error("SetUserLoginLog   error:", zap.Error(err))
		return err
	}
	return nil
}

//设置用户数据
func (userService *UserService) SetUserStatisInfo(statis system.UserStatisInfo) error {
	err := global.GV_DB.Table(model.UserStatisticsInfoTableName).Create(&statis).Error
	if err != nil {
		global.GV_LOG.Error("SetUserStatisInfo error:", zap.Error(err))
		return err
	}
	return nil
}

//用户登录次数或者完成订单总数+1
func (userService *UserService) IncUserStatisCount(uid int64, incType int) error {
	var statis system.UserStatisInfo
	err := global.GV_DB.Table(model.UserStatisticsInfoTableName).Where("uid = ?", uid).First(&statis).Error
	if err != nil {
		global.GV_LOG.Error("IncUserLoginCount error:", zap.Error(err))
		return err
	}
	if incType == model.IncLoginCount {
		statis.LoginCount++
	} else if incType == model.IncFinishOrderCount {
		statis.OrderCount++
	}
	err = global.GV_DB.Table(model.UserStatisticsInfoTableName).Updates(&statis).Error
	if err != nil {
		global.GV_LOG.Error("IncUserLoginCount error:", zap.Error(err))
		return err
	}
	return nil
}

//获取用户数据
func (userService *UserService) GetUserStatisInfo(uid int64) (*system.UserStatisInfo, error) {
	var statis system.UserStatisInfo
	err := global.GV_DB.Table(model.UserStatisticsInfoTableName).Where("uid = ?", uid).First(&statis).Error
	if err != nil {
		global.GV_LOG.Error("SetUserStatisInfo error:", zap.Error(err))
		return nil, err
	}
	return &statis, nil
}

//获取联系人id

func (userService *UserService) GetFriendIds(uid int64) ([]int64, error) {
	var ids []int64
	err := global.GV_DB.Table(model.MessageTableName).Select("to_uid").Where("from_uid = ?", uid).Find(&ids).Error
	if err != nil {
		global.GV_LOG.Error("GetFriendIds error:", zap.Error(err))
		return nil, err
	}
	return ids, nil
}
