package upload

import (
	"errors"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/pkg/timer"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type Local struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	//读取文件后缀
	ext := path.Ext(file.Filename)
	name := strings.TrimSuffix(file.Filename, ext)
	//拼接新文件名
	filename := name + "_" + time.Now().Format(timer.TimeFormatDateV3) + ext
	//尝试创建路径
	mkdirErr := os.MkdirAll(global.GV_CONFIG.Local.Path, os.ModePerm)
	if mkdirErr != nil {
		global.GV_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	//拼接路径和文件名
	p := global.GV_CONFIG.Local.Path + "/" + filename

	f, openError := file.Open() //读取文件
	if openError != nil {
		global.GV_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		global.GV_LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.GV_LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*Local) DeleteFile(key string) error {
	p := global.GV_CONFIG.Local.Path + "/" + key
	if strings.Contains(p, global.GV_CONFIG.Local.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败，err:" + err.Error())
		}
	}
	return nil
}
