package utils_test

import (
	"fmt"
	"github.com/tmnhs/fginx/server/pkg/utils"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	content, err := utils.ReadFile(utils.VerificationCodeFile)
	if err != nil {
		panic(err)
	}
	content = strings.Replace(content, "userSlot", "tmnhs", 1)
	content = strings.Replace(content, "emailSlot", "1685290935@qq.com", 1)
	newContent := strings.Replace(content, "codeSlot", "7777", 1)

	//fmt.Println("content:",content)
	fmt.Println("newContent:", newContent)
}
