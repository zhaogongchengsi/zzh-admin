package request

type CosTempKeyRequest struct {
	PlaceOfOwnership string `json:"Region" binding:"required"` // 归属地
	BucketName string `json:"Bucket" binding:"required"` // 储存桶名称
	Action string `json:"action"` //操作名称
}
