package request

type GetGormById struct {
	ID float64 `json:"id" form:"id"` // 主键ID
}

type LimitOffset struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

func NewLimit(Limit int, offset int) *LimitOffset {
	return &LimitOffset{
		Limit:Limit,
		Offset: offset,
	}
}