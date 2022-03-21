{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

const (
	Or = "|"
)

// https://www.iris-go.com/docs/#/?id=bind-any

type Pager struct {
	Page         int    `json:"page" form:"page" query:"page" binding:"gte=1"`                    // 页码
	PageSize     int    `json:"page_size" form:"page_size" query:"page" binding:"gte=10,lte=100"` // 每页数量
	Order        int    `json:"-" form:"order" query:"page" binding:"oneof=0 1 2 3"`          // 已定义字段排序
	OrderByField string `json:"-" form:"order_by_field" query:"page"`                // 自定义字段排序
	TotalRows    int64  `json:"total_rows"`                                                       // 总行数
}

func (p Pager) Offset() int {
	return (p.Page - 1) * p.PageSize
}

type Search struct {
	LikeQ string `json:"like_q" form:"like_q" url:"like_q" example:"Fuzzy query words, multiple query words are separated by |, golang|cpp|rust"`
	MustQ string `json:"must_q" form:"must_q" url:"must_q" example:"Precise query words, multiple query words are separated by |, golang|cpp|rust"`
	NotQ  string `json:"not_q" form:"not_q" url:"not_q" example:"Not query words, multiple query words are separated by |, golang|cpp|rust"`

	TimeBegin string `json:"time_begin" form:"time_begin" url:"time_begin" example:"2022-01-01"`
	TimeEnd   string `json:"time_end" form:"time_end" url:"time_end" example:"2022-12-31"`
}

type SearchPager struct {
	Search
	Pager
}
