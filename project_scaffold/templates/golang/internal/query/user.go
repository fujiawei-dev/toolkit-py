{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"strings"

	"github.com/jinzhu/copier"

	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
)

type UserResult struct {
	ID       uint   `json:"id" example:"1"` // 记录ID
	Username string `json:"username" example:"用户名"`
	Enable   bool   `json:"enable"` // 是否启用                                                // 是否启用

	CreatedAt JSONTime `json:"created_at" example:"2022-03-21 08:57:19"` // 创建时间
}

func Users(f form.SearchPager) (results []UserResult, totalRows int64, err error) {
	query := Db().Model(&entity.User{})

	if f.LikeQ != "" {
		for _, q := range strings.Split(f.LikeQ, form.Or) {
			query = query.Where("username LIKE ?", "%"+q+"%")
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

	var items entity.Users

	if err = query.Offset(f.Offset()).Limit(f.PageSize).Find(&items).Error; err != nil {
		return
	}

	err = copier.Copy(&results, items)

	return
}
