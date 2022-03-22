{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"strings"

	"github.com/jinzhu/copier"

	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
)

type EntityTemplateEmbeddedResult struct {
	ID    uint   `json:"id" example:"1"` // 记录ID
	Email string `json:"email" example:"who@gmail.com"`
}

type EntityTemplateResult struct {
	ID uint `json:"id" example:"1"` // 记录ID

	EntityTemplateEmbedded EntityTemplateEmbeddedResult `json:"user"`

	When  string `json:"when" example:"时间"`
	Where string `json:"where" example:"地点"`
	Who   string `json:"who" example:"人物"`
	What  string `json:"what" example:"事件"`
	How   string `json:"how" example:"过程"`

	CreatedAt JSONTime `json:"created_at" example:"2022-03-21 08:57:19"` // 创建时间
}

func EntityTemplates(f form.SearchPager) (results []EntityTemplateResult, totalRows int64, err error) {
	query := Db().Model(&entity.EntityTemplate{})

	if f.LikeQ != "" {
		for _, q := range strings.Split(f.LikeQ, form.Or) {
			query = query.Where("who LIKE ?", "%"+q+"%")
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

	var items entity.EntityTemplates

	if err = query.Offset(f.Offset()).Limit(f.PageSize).
		Preload("EntityTemplateEmbedded").
		Find(&items).Error; err != nil {
		return
	}

	err = copier.Copy(&results, items)

	return
}
