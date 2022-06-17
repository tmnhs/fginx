package timer

import "time"

// */10 * * * * ? 每隔10秒执行一次
//
//0 */5 * * * ? 每隔5分钟执行一次
//
//0 2,22,32 * * * ? 在2分、22分、32分执行一次
//
//0 0 4-8 * * ? 每天4-8点整点执行一次
//
//0 0 2 * * ? 每天凌晨2点执行一次
//
//0 0 2 1 * ? 每月1号凌晨2点执行一次
const (
	TimeFormatSecond                 = "2006-01-02 15:04:05"
	TimeFormatMinute                 = "2006-01-02 15:04"
	TimeFormatDateV1                 = "2006-01-02"
	TimeFormatDateV2                 = "2006_01_02"
	TimeFormatDateV3                 = "20060102150405"
	TimeFormatDateV4                 = "2006/01/02 - 15:04:05.000"
	CodeExpirationTime time.Duration = 10 * time.Minute //验证码过期时间为10分钟
	//UserOnlineExpireTime = 180 * time.Second //心跳机制，设置过期时间
	UserOnlineExpireTime    = 180 * time.Hour      //心跳机制，设置过期时间
	UserTokenExpireTime     = 24 * 150 * time.Hour //设置token过期时间
	HeartbeatExpirationTime = 60 * 6               //六分钟
	ClearTimeoutConnections = "*/5 * * * * ?"
)

/**字符串->时间对象*/
func Str2Time(formatTimeStr string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime

}

/**字符串->时间戳*/
func Str2Stamp(formatTimeStr string) int64 {
	timeStruct := Str2Time(formatTimeStr)
	millisecond := timeStruct.UnixNano() / 1e6
	return millisecond
}

/**时间对象->字符串*/
func Time2Str(t time.Time) string {
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(TimeFormatSecond)
	return str
}

/*时间对象->时间戳*/
func Time2Stamp(t time.Time) int64 {
	millisecond := t.UnixNano() / 1e6
	return millisecond
}

/*时间戳->字符串*/
func Stamp2Str(stamp int64) string {
	str := time.Unix(stamp/1000, 0).Format(TimeFormatSecond)
	return str
}

/*时间戳->时间对象*/
func Stamp2Time(stamp int64) time.Time {
	stampStr := Stamp2Str(stamp)
	timer := Str2Time(stampStr)
	return timer
}

/*获取当前时间的时间戳*/
func GetNowUnix() int64 {
	return time.Now().Unix()
}
