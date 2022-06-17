package file

import "github.com/tmnhs/fginx/server/global"

type FileUploadAndDownload struct {
	global.GV_MODEL
	Url string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Key string `json:"key" gorm:"comment:编号"`   // 编号
}
