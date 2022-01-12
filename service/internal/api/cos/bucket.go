package cos

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github/gin-react-admin/pkg/response"
	"net/http"
	"net/url"
	"os"
)

func GetDuckerList(c *gin.Context)  {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse("https://bloghtml-1301735126.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv("COS_SECRETID"),
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv("COS_SECRETKEY"),
		},
	})

	res2, _, err := client.Service.Get(context.Background())
	if err != nil {
		res := response.Response{Err: err, State: response.State{
			Message: "获取储存桶列表失败",
		}}
		res.SendError(c)
		return
	}
	res := response.Response{
		Data: res2,
	}
	res.Send(c)
}