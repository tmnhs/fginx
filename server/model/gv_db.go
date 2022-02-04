package model

import (
	config "github.com/flipped-aurora/gin-vue-admin/server/config/cmap"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// DB 初始化一个sql对象
func DB(globalDb string) *gorm.DB {
	if globalDb == "GV_DB" {
		return global.GV_DB
	} else {
		return global.MustGetGlobalDBByDBName(globalDb)
	}
}

//InitDB 初始化数据库
func InitDB(tableAlias string, rw string, suffix string) *gorm.DB {
	TableRegConfigStruct, err := config.InitTables(tableAlias)
	if err != nil {
		return nil
	}
	var r *gorm.DB
	//读库/写库
	if rw == "w" {
		r = DB(TableRegConfigStruct.GlobalDbWrite)
	} else {
		r = DB(TableRegConfigStruct.GlobalDbRead)
	}
	//前缀
	var tableName string
	if TableRegConfigStruct.Prefix != "" {
		//覆盖前缀
		tableName = TableRegConfigStruct.Prefix
	} else {
		//配置yaml前缀
		tableName = global.GV_DBList[TableRegConfigStruct.GlobalDbWrite].TablePre
	}
	//前缀组合表名
	tableName = tableName + TableRegConfigStruct.TableName
	//后缀
	return r.Table(tableName + suffix)
}
