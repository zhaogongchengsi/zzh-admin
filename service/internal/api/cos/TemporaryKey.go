package cos

import (
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"github/gin-react-admin/global"
	"github/gin-react-admin/pkg/request"
	"github/gin-react-admin/pkg/response"
	"os"
	"time"
)


// @Tags TempKey
// @Summary 获取cos 临时密匙
// @Produce  application/json
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /cos/get_temporary_key [get]
func GetTempKey (c *gin.Context) {
	var cosKey request.CosTempKeyRequest
	if err := c.ShouldBindJSON(&cosKey); err != nil {
		res := response.Response{Err: err}
		res.SendInputParameterError(c)
		return
	}
	tempoKey, err := IssueCosTempKey(cosKey.BucketName ,cosKey.PlaceOfOwnership,"*")
	if err != nil {
		res := response.Response{Err: err, State: response.State{
			Message: "cos临时密匙获取失败",
		}}
		res.SendError(c)
		return
	}
	res := response.Response{
		Data: tempoKey,
	}
	res.Send(c)
}

// 储存桶名字 归属地名称 权限配置
func IssueCosTempKey (bName string, poName string, resource string) (*sts.CredentialResult, error) {
	appid := global.Cos.AppId
	c := sts.NewClient(
		os.Getenv("COS_SECRETID"),
		os.Getenv("COS_SECRETKEY"),
		nil,
	)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          poName,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					// 密钥的权限列表。简单上传和分片需要以下的权限，其他权限列表请看 https://cloud.tencent.com/document/product/436/31923
					Action: []string{
						// 简单上传
						"name/cos:PostObject",
						"name/cos:PutObject",
						// 分片上传
						"name/cos:InitiateMultipartUpload",
						"name/cos:ListMultipartUploads",
						"name/cos:ListParts",
						"name/cos:UploadPart",
						"name/cos:CompleteMultipartUpload",
					},
					Effect: "allow",
					Resource: []string{
						// 这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						"qcs::cos:"+ poName +":uid/" + appid + ":" + bName + "/" + resource,
					},
				},
			},
		},
	}
	res, err := c.GetCredential(opt)
	return res, err
}