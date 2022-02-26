{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"strings"
	"time"

	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
)

func Examples(f form.Pager) (results entity.Examples, totalRows int64, err error) {
	query := Db().Model(&entity.Example{})

	if err = query.Count(&totalRows).Error; err != nil {
		return
	}

	err = query.Offset(f.Offset()).Limit(f.PageSize).Find(&results).Error

	return
}

func FindExamplesBySearch(s Search, f form.Pager) (results entity.Examples, totalRows int64, err error) {
	query := Db().Model(&entity.Example{})

	if s.LikeQ != "" {
		for _, q := range strings.Split(s.LikeQ, Or) {
			query = query.Where("long_text_field LIKE ?", "%"+q+"%")
		}
	}

	if s.MustQ != "" {
		query = query.Where("uid IN ?", strings.Split(s.MustQ, Or))
	}

	if s.NotQ != "" {
		query = query.Not("id IN ?", strings.Split(s.NotQ, Or))
	}

	if s.TimeBegin != "" && s.TimeEnd != "" {
		query.Where("created_at BETWEEN ? AND ?", s.TimeBegin, s.TimeEnd)
	} else if s.TimeBegin != "" {
		query.Where("created_at > ?", s.TimeBegin)
	} else if s.TimeEnd != "" {
		query.Where("created_at < ?", s.TimeEnd)
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

type ResultGroupBySearch struct {
	Search

	Date  time.Time
	Total int
}

func ExamplesGroupBySearch(s Search) (result ResultGroupBySearch, err error) {
	query := Db().Model(&entity.Example{})

	// https://gorm.io/docs/query.html#Group-By-amp-Having

	if s.LikeQ != "" {
		query = query.Select("long_text_field AS like_q, SUM(integer_field) AS total")

		for _, q := range strings.Split(s.LikeQ, Or) {
			query = query.Where("long_text_field LIKE ?", "%"+q+"%")
		}

		err = query.Group("long_text_field").First(&result).Error

		return
	}

	return
}
