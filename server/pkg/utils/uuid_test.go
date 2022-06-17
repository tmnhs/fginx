package utils_test

import (
	"github.com/tmnhs/fginx/server/pkg/utils"
	"testing"
)

func TestCreateVerificationCode(t *testing.T) {
	code, err := utils.CreateVerificationCode()
	if err != nil {
		t.Error("err:", err)
		return
	}
	t.Log("code:", code)
}

func TestCreateUUID(t *testing.T) {
	uuid := utils.CreateUUID()
	t.Log("uuid:", uuid)

}
