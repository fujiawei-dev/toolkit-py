{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

// https://www.iris-go.com/docs/#/?id=bind-any

// Pager 分页处理
type Pager struct {
	Page      int   `json:"page" form:"page" url:"page" binding:"gte=1"`                         // 页码
	PageSize  int   `json:"page_size" form:"page_size" url:"page_size" binding:"gte=10,lte=100"` // 每页数量
	Order     int   `json:"order" form:"order" binding:"oneof=0 1 2 3"`                          // 排序顺序
	TotalRows int64 `json:"total_rows"`                                                          // 总行数
}

func (p Pager) Offset() int {
	return (p.Page - 1) * p.PageSize
}
