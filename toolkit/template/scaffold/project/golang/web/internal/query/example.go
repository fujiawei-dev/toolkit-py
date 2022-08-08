package query

import (
	"strings"

	"{{ main_module }}/internal/entity"
	"{{ main_module }}/internal/form"
)

func Examples(f form.SearchPager) (results entity.Examples, totalRows int64, err error) {
	query := Db().Model(&entity.Example{})

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
