package utils

import (
	"encoding/base64"
	"github.com/tmnhs/fginx/server/global"
	"go.uber.org/zap"
	"golang.org/x/crypto/scrypt"
)

// 使用scrypt密码加密
func ScryptPaassword(password string) string {
	const KeyLen = 32
	salt := make([]byte, 8)
	salt = []byte{12, 32, 14, 6, 66, 22, 43, 11}
	//加密
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		global.GV_LOG.Error("ScryptPaassword err:%s", zap.Error(err))
	}
	//将加密后的密码转化为字符串
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
