package request

type CosTempKeyRequest struct {
	PlaceOfOwnership string `json:"place_of_ownership" binding:"required"` // 归属地
	BucketName string `json:"bucket_name" binding:"required"` // 储存桶名称
	Action string `json:"action"` //操作名称
}
