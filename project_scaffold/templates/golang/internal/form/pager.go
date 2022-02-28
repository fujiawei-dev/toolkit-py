{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

// https://www.iris-go.com/docs/#/?id=bind-any

// Pager 分页处理
type Pager struct {
	Page         int    `json:"page" form:"page" query:"page" url:"page" binding:"gte=1"`                         // 页码
	PageSize     int    `json:"page_size" form:"page_size" query:"page" url:"page_size" binding:"gte=10,lte=100"` // 每页数量
	Order        int    `json:"order" form:"order" query:"page" url:"order" binding:"oneof=0 1 2 3"`              // 已定义字段排序
	OrderByField string `json:"order_by_field" form:"order_by_field" query:"page" url:"order_by_field"`           // 自定义字段排序
	TotalRows    int64  `json:"total_rows"`                                                          // 总行数
}

func (p Pager) Offset() int {
	return (p.Page - 1) * p.PageSize
}
