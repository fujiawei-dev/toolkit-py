{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"strings"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
)

type OperationLogResult struct {
	ID uint `json:"id" example:"1"` // 记录ID

	User UserResult `json:"user"`

	Resource string `json:"resource" example:"操作资源: 比如 users"`
	Action   string `json:"action" example:"操作行为: 比如 login、delete、create"`
	Allow    bool   `json:"allow"` // 操作是否被允许

	CreatedAt JSONTime `json:"created_at" example:"2022-03-21 08:57:19"` // 创建时间
}

func OperationLogs(f form.SearchPager) (results []OperationLogResult, totalRows int64, err error) {
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

	var items entity.OperationLogs

	if err = query.Offset(f.Offset()).Limit(f.PageSize).Preload("User").Find(&items).Error; err != nil {
		return
	}

	err = copier.Copy(&results, items)

	return
}

func DeleteAllOperationLogs() (err error) {
	return Db().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity.OperationLog{}).Error
}
