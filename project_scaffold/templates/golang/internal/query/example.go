{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
)

func Examples(f form.Pager) (results entity.Examples, totalRows int64, err error) {
	query := Db().Model(&entity.Example{})
	err = query.Count(&totalRows).Error
	err = query.Offset(f.Offset()).Limit(f.PageSize).Find(&results).Error
	return
}
