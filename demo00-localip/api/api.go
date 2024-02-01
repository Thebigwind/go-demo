// api.go
package api

import (
	"os"

	"github.com/tal-tech/go-zero/rest/httpx"
)

type FileDownloadHandler struct {
}

func (h *FileDownloadHandler) DownloadFile(ctx *httpx.Context) error {
	// Replace the file path with the actual path of the file you want to serve
	filePath := "path/to/your/file.txt"

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Set response headers
	ctx.Response.Header().Set("Content-Disposition", "attachment; filename="+file.Name())
	ctx.Response.Header().Set("Content-Type", "application/octet-stream")

	return httpx.File(ctx.Response, file)
}
