package fileService

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/pkg/response"
	"github/gin-react-admin/utils"
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
	filePath := utils.MkDir(file.Filename)
	// 上传文件到指定的 dst 。
	_filePath, upper := utils.SaveUploadedFile(file, filePath, false)

	//upper := c.SaveUploadedFile(file, upPath + "/" + file.Filename)
	if upper != nil {
		reuse := response.Response{
			Err: upper,
			State: response.State{
				Message: "文件保存失败",
			},
		}
		reuse.SendError(c)
		return
	}
	resok := response.Response{ Data:_filePath }
	resok.Send(c)
}