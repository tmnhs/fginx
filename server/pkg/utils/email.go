package utils

import (
	"github.com/tmnhs/fginx/server/global"
	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string) error {

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(global.GV_CONFIG.Email.From, "hust_mall客服")) //这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])

	m.SetHeader("To", mailTo...) //发送给多个用户
	//m.SetHeader("To", mailTo[0])    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(global.GV_CONFIG.Email.Host, global.GV_CONFIG.Email.Port, global.GV_CONFIG.Email.From, global.GV_CONFIG.Email.Secret)

	return d.DialAndSend(m)

}
