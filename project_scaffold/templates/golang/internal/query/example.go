{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"{{GOLANG_MODULE}}/internal/config"
	"{{GOLANG_MODULE}}/internal/entity"
)

func Examples(f form.Pager) (results entity.Examples, totalRows int64, err error) {
	query := Db().Model(&entity.Example{})
	err = query.Count(&totalRows).Error
	err = query.Offset(f.Offset()).Limit(f.PageSize).Find(&results).Error
	return
}
