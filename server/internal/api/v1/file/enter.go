package file

import "github.com/tmnhs/fginx/server/internal/service"

type ApiGroup struct {
	FileUploadAndDownloadApi
}

var (
	fileUploadAndDownloadService = service.ServiceGroupApp.FileServiceGroup.FileUploadAndDownloadService
)
