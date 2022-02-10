{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

// https://www.iris-go.com/docs/#/?id=bind-any

// Pager 分页处理
type Pager struct {
	Page      int   `json:"page" form:"page,default=1" url:"page" validate:"gte=1"`                          // 页码
	PageSize  int   `json:"page_size" form:"page_size,default=50" url:"page_size" validate:"gte=10,lte=100"` // 每页数量
	TotalRows int64 `json:"total_rows"`                                                                      // 总行数
}

func (p Pager) Offset() int {
	return (p.Page - 1) * p.PageSize
}
