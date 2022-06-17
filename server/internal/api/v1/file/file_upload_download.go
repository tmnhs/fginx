package file

import (
	"github.com/gin-gonic/gin"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/model/common/response"
	"github.com/tmnhs/fginx/server/internal/model/file"
	fileRes "github.com/tmnhs/fginx/server/internal/model/file/response"
	"github.com/tmnhs/fginx/server/pkg/errcode"
	"go.uber.org/zap"
)

type FileUploadAndDownloadApi struct{}

func (fileUploadAndDownloadApi *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GV_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage(errcode.ErrorFile, "接收文件失败", c)
		return
	}
	err, f := fileUploadAndDownloadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GV_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage(errcode.ErrorFile, "修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(fileRes.FileResponse{File: f}, "上传成功", c)
}

func (fileUploadAndDownloadApi *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var req file.FileUploadAndDownload
	_ = c.ShouldBindJSON(&req)
	if err := fileUploadAndDownloadService.DeleteFile(req); err != nil {
		global.GV_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(errcode.ErrorFileDelete, "删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
