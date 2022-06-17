package utils_test

import (
	"github.com/tmnhs/fginx/server/pkg/utils"
	"testing"
)

func TestScryptPaassword(t *testing.T) {
	password := "123456"
	newPassword := utils.ScryptPaassword(password)

	t.Log("newPassword:", newPassword)
}
