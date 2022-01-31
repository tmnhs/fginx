package file

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"mime/multipart"
	"strings"
)

type FileUploadAndDownloadService struct{}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *FileUploadAndDownloadService) Upload(file file.FileUploadAndDownload) error {
	return global.GV_DB.Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFile
//@description: 删除文件切片记录
//@param: id uint
//@return: error, model.ExaFileUploadAndDownload

func (e *FileUploadAndDownloadService) FindFile(id uint) (error, file.FileUploadAndDownload) {
	var file file.FileUploadAndDownload
	err := global.GV_DB.Where("id = ?", id).First(&file).Error
	return err, file
}
//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.ExaFileUploadAndDownload
//@return: err error

func (e *FileUploadAndDownloadService) DeleteFile(f file.FileUploadAndDownload) (err error) {
	var fileFromDb file.FileUploadAndDownload
	err, fileFromDb = e.FindFile(f.ID)
	if err != nil {
		return
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GV_DB.Where("id = ?", f.ID).Unscoped().Delete(&f).Error
	return err
}
//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func  (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (err error, f file.FileUploadAndDownload)  {
	oss:=upload.NewOss()
	filePath,key,uploadErr:=oss.UploadFile(header)
	if uploadErr!=nil{
		panic(uploadErr)
	}
	if noSave=="0"{
		s:=strings.Split(header.Filename,".")
		f:=file.FileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return e.Upload(f),f
	}
	return
}