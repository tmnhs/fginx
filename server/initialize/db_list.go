package initialize

import "github.com/tmnhs/fginx/server/global"

const sys = "system"

func DBList() {
	dbMap := make(map[string]global.DbStruct)
	for _, info := range global.GV_CONFIG.DBList {
		if info.Disable {
			continue
		}
		dbMap[info.GlobalDb] = global.DbStruct{
			DB:       GormMysqlByConfig(info),
			TablePre: info.TablePre,
		}
	}
	global.GV_DBList = dbMap
}
