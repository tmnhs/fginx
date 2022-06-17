package file

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/tmnhs/fginx/server/internal/api/v1"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := Router.Group("file")
	exaFileUploadAndDownloadApi := v1.ApiGroupApp.FileApiGroup.FileUploadAndDownloadApi
	{
		fileUploadAndDownloadRouter.POST("upload", exaFileUploadAndDownloadApi.UploadFile) // 上传文件
		fileUploadAndDownloadRouter.POST("delete", exaFileUploadAndDownloadApi.DeleteFile) // 删除指定文件
	}
}
