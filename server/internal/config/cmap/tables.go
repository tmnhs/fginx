package config

import (
	"errors"
	"strings"
)

//  注册结构体
type TableRegConfigStruct struct {
	GlobalDbRead  string //读库
	GlobalDbWrite string //写库
	TableName     string //表名
	Prefix        string //前缀 （注意，公共前缀在yaml配置里，如果这里配置了前缀，公共前缀会被覆盖）
	Suffix        bool   //后缀模式  （hash，时间等分表）
}

//getTablesConfig 获取table-config
func getTablesConfig() map[string]TableRegConfigStruct {
	tableAliasMap := make(map[string]TableRegConfigStruct)
	//注册的时候，请确认数据表和GlobalDbRead/GlobalDbWrite的关系
	tableAliasMap["check_images"] = TableRegConfigStruct{
		GlobalDbRead:  "GVA_DBMatchSns_R",
		GlobalDbWrite: "GVA_DBMatchSns_W",
		TableName:     "check_images",
		Suffix:        true,
	}
	return tableAliasMap
}

//InitTables 初始化tables
func InitTables(tableAlias string) (*TableRegConfigStruct, error) {
	tableAliasMap := getTablesConfig()
	tableAlias = strings.Replace(tableAlias, "m_", "", 1)
	if r, ok := tableAliasMap[tableAlias]; ok {
		return &r, nil
	} else {
		return nil, errors.New(tableAlias + "的tableAlias未注册，无法使用")
	}
}
