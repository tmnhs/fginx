package utils_test

import (
	"github.com/tmnhs/fginx/server/pkg/utils"
	"testing"
)

func TestGetIpInfo(t *testing.T) {
	//fmt.Println("get ip",utils.GetIP())
	t.Log("ip_info:", *utils.GetIpAddress(`114.55.178.217`))
}
