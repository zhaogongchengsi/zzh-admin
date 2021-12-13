package response

import (
	"github/gin-react-admin/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

type State struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type Response struct {
	State State       `json:"state"`
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Err   interface{} `json:"err"`
}

// Send 成功时 发送数据
func (r Response) Send(c *gin.Context) {
	r.State.Code = 200
	r.State.Message = "成功"
	r.State.Success = true
	c.JSON(http.StatusOK, gin.H{"state": r.State, "code": r.Code, "msg": r.Msg, "data": r.Data})
}

// SendError 失败时 发送错误码以及发送错误消息
func (r Response) SendError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"state": r.State, "err": r.Err})
}

// SendInputParameterError 入参错误时 发送错误码以及发送错误消息
func (r Response) SendInputParameterError(c *gin.Context) {
	r.State.Code = errcode.InvalidParams
	r.State.Message = "缺少参数"
	r.State.Success = false
	c.JSON(http.StatusOK, gin.H{"state": r.State, "err": r.Err})
}

//
//// CreateOk 资源创建成功
//func (r Response) CreateOk(c *gin.Context) {
//	c.JSON(http.StatusCreated, gin.H{"Message": r.Msg, "data": r.Data})
//}
//
//// CreateFail 资源创建失败
//func (r Response) CreateFail(c *gin.Context) {
//	c.JSON(http.StatusAccepted, gin.H{"Message": r.Msg, "error": r.Err})
//}
