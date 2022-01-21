package fileService

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/service"
	"github/gin-react-admin/pkg/response"
)


// @Tags Article
// @Summary 上传文件
// @Produce  multipart/form-data
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /api/v1/upload_download/upload [post]
func UpLoad(c *gin.Context)  {
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

type StringFile struct {
	FileName string `json:"file_name"`
	TextContext string `json:"text_context"`
}


// @Tags Article
// @Summary 上传文件
// @Produce  application/json
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /api/v1/upload_download/upload_string [post]
func UpStringWritFile (c *gin.Context) {
	var sf StringFile
	if err := c.ShouldBindJSON(&sf); err != nil {
		res := response.Response{Err: err}
		res.SendInputParameterError(c)
		return
	}

	smod , ferr := service.SaveString(sf.TextContext, sf.FileName, false)

	if ferr != nil {
		reuse := response.Response{
			Err: ferr,
			State: response.State{
				Message: "文件保存失败",
			},
		}
		reuse.SendError(c)
		return
	}

	resok := response.Response{ Data:smod}
	resok.Send(c)
}


