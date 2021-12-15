package request

type GetGormById struct {
	ID float64 `json:"id" form:"id"` // 主键ID
}