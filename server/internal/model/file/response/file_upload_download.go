package response

import (
	"github.com/tmnhs/fginx/server/internal/model/file"
)

type FileResponse struct {
	File file.FileUploadAndDownload `json:"file"`
}
