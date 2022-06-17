package utils

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"io/ioutil"
	"net/http"
	"strings"
)

//func GetIP( r *http.Request) string {
//	ip := exnet.ClientPublicIP(r)
//	if ip == "" {
//		ip = exnet.ClientIP(r)
//	}
//
//	return ip
//}

//爬虫获取ip所在城市
func GetIpAddress(ip string) *string {
	if len(ip) == 0 {
		return nil
	}
	var address string
	resp, err := http.Get(fmt.Sprintf("http://www.cip.cc/%s", ip))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	doc := soup.HTMLParse(string(out))
	//爬取数据
	subDocs := doc.Find("pre")
	contents := strings.Split(subDocs.Text(), "\n")
	for _, content := range contents {
		if strings.Contains(content, "地址") == true {
			vvsplit := strings.Split(strings.Split(content, ":")[1], " ")
			for _, vv := range vvsplit {
				if len(vv) != 0 {
					address = address + vv + "/"
				}
			}
		}
	}
	//去除最后一个"/"
	result := address[:len(address)-1]
	return &result
}
