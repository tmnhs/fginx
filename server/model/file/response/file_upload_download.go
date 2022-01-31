package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
)

type FileResponse struct {
	File file.FileUploadAndDownload `json:"file"`
}
