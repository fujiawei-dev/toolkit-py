{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"strings"

	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
)

func OperationLogs(f form.SearchPager) (results entity.OperationLogs, totalRows int64, err error) {
	query := Db().Model(&entity.OperationLog{})

	if f.LikeQ != "" {
		for _, q := range strings.Split(f.LikeQ, form.Or) {
			query = query.Where("long_text_field LIKE ?", "%"+q+"%")
		}
	}

	if f.MustQ != "" {
		query = query.Where("uid IN ?", strings.Split(f.MustQ, form.Or))
	}

	if f.NotQ != "" {
		query = query.Not("id IN ?", strings.Split(f.NotQ, form.Or))
	}

	if f.TimeBegin != "" && f.TimeEnd != "" {
		query.Where("created_at BETWEEN ? AND ?", f.TimeBegin, f.TimeEnd)
	} else if f.TimeBegin != "" {
		query.Where("created_at > ?", f.TimeBegin)
	} else if f.TimeEnd != "" {
		query.Where("created_at < ?", f.TimeEnd)
	}

	if err = query.Count(&totalRows).Error; err != nil {
		return
	}

	if f.OrderByField != "" {
		query = query.Order(f.OrderByField)
	} else if f.Order == 1 {
		query = query.Order("id DESC")
	}

	err = query.Offset(f.Offset()).Limit(f.PageSize).Find(&results).Error

	return
}
