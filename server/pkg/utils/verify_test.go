package utils_test

import (
	"github.com/tmnhs/fginx/server/pkg/utils"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	var regexpImpl utils.Regexp

	//result:=regexp.MustCompile(regexpImpl.Email()).MatchString("1685290935@qq.com")
	//result:=regexp.MustCompile(regexpImpl.Email()).MatchString("1685290935qq.com")

	//result:=regexp.MustCompile(regexpImpl.Phone()).MatchString("19873347933")
	//result:=regexp.MustCompile(regexpImpl.Phone()).MatchString("19873347988888")

	//result:=regexp.MustCompile(regexpImpl.QQ()).MatchString("1682")
	//result:=regexp.MustCompile(regexpImpl.QQ()).MatchString("1685290935")

	//result:=regexp.MustCompile(regexpImpl.Password()).MatchString("aaaaabaaaaaaaaaa")
	result := regexp.MustCompile(regexpImpl.Password()).MatchString("aaaaabaaaa_")

	t.Log("result:", result)
}
