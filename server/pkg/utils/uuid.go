package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//工具，生成唯一字符串
func CreateUUID() string {
	id := uuid.NewV4()
	return id.String()
}

//生成四位数字验证码
func CreateVerificationCode() (string, error) {
	//codeHub:="123456789qwertyuipasdfghjklzxcvbnmQWERTYUIPASDFGHJKLZXCVBNM"
	codeHub := "1234567890"
	rand.Seed(time.Now().Unix())
	var code strings.Builder
	for i := 0; i < 4; i++ {
		index := rand.Intn(len(codeHub))
		_, err := fmt.Fprintf(&code, "%c", codeHub[index])
		if err != nil {
			return "", err
		}
	}
	fmt.Println(code.String())
	//global.GV_LOG.Info("VerificationCode:%s",zap.String("code",code.String()))
	return code.String(), nil
}
func OperationIDGenerator() string {
	return strconv.FormatInt(time.Now().UnixNano()+int64(rand.Uint32()), 10)
}
