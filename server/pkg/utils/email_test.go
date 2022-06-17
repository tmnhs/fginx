package utils_test

import (
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/config"
	"github.com/tmnhs/fginx/server/pkg/utils"
	"strings"
	"testing"
)

func TestSendEmail(t *testing.T) {
	//定义收件人
	mailTo := []string{
		"mananhai@highlight.mobi",
		//"1685290935@qq.com",
		//"2268080550@qq.com",
	}
	/*//邮件主题为"Hello"
	subject := "Hello by golang gomail from hust_mall"
	// 邮件正文
	body := "Hello,by gomail sent,this is a test email"*/
	subject := "hust_mall客服"
	content, err := utils.ReadFile("../../../docs/code.html")
	content = strings.Replace(content, "codeSlot", "7777", 1)
	content = strings.Replace(content, "emailSlot", "2268080550@qq.com", 1)
	body := strings.Replace(content, "userSlot", "yihao", 1)
	global.GV_CONFIG.Email = config.Email{
		Port:   465,
		From:   "1685290935@qq.com",
		Host:   "smtp.qq.com",
		Secret: "otjkvfgiguswciga",
		IsSSL:  false,
	}
	err = utils.SendMail(mailTo, subject, body)
	if err != nil {
		t.Logf("send fail")
		t.Error(err)
		return
	}
	t.Logf("send successfully")
}
