package fileService

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/service"
	"github/gin-react-admin/pkg/response"
)

func UpLoad(c *gin.Context)  {
	//var upPath string = global.FileService.UploadPath

	file, err := c.FormFile("file")
	if err != nil {
	 	reset := response.Response{
			Err: err,
			State: response.State{
				Message: "文件上传失败",
			},
		}
		reset.SendError(c)
		return
	}

	fileMo, e := service.FileSave(file, false)
	if e != nil {
		reuse := response.Response{
			Err: e,
			State: response.State{
				Message: "文件保存失败",
			},
		}
		reuse.SendError(c)
		return
	}
	resok := response.Response{ Data:fileMo }
	resok.Send(c)
}