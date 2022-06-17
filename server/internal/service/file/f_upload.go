package file

import (
	"errors"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/model"
	"github.com/tmnhs/fginx/server/internal/model/file"
	"github.com/tmnhs/fginx/server/pkg/upload"
	"go.uber.org/zap"
	"mime/multipart"
)

type FileUploadAndDownloadService struct{}

func (e *FileUploadAndDownloadService) Upload(file file.FileUploadAndDownload) error {
	return global.GV_DB.Table(model.FileTableName).Create(&file).Error
}

func (e *FileUploadAndDownloadService) FindFile(id uint) (error, file.FileUploadAndDownload) {
	var f file.FileUploadAndDownload
	err := global.GV_DB.Table(model.FileTableName).Where("id = ?", id).First(&f).Error
	return err, f
}

func (e *FileUploadAndDownloadService) FindFileByUrl(url string) (error, file.FileUploadAndDownload) {
	var f file.FileUploadAndDownload
	err := global.GV_DB.Table(model.FileTableName).Where(" `url` = ?", url).First(&f).Error
	return err, f
}

func (e *FileUploadAndDownloadService) DeleteFile(f file.FileUploadAndDownload) (err error) {
	var fileFromDb file.FileUploadAndDownload
	err, fileFromDb = e.FindFileByUrl(f.Url)
	if err != nil {
		global.GV_LOG.Error("DeleteFile--FindFileByUrl error", zap.Error(err))
		return err
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GV_DB.Table(model.FileTableName).Where("`url` = ?", f.Url).Unscoped().Delete(&f).Error
	return err
}

func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (err error, f file.FileUploadAndDownload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(uploadErr)
	}
	if noSave == "0" {
		f := file.FileUploadAndDownload{
			Url: filePath,
			Key: key,
		}
		return e.Upload(f), f
	}
	return
}
