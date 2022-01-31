package file

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type FileApiGroup struct {
	FileUploadAndDownloadApi
}

var (
	fileUploadAndDownloadService = service.ServiceGroupApp.FileServiceGruop.FileUploadAndDownloadService
)