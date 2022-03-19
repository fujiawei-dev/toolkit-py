{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"gorm.io/gorm"

	"{{GOLANG_MODULE}}/internal/config"
	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/event"
)

var (
	log  = event.Logger()
	conf = config.Conf()
)

// Db returns a database connection instance.
func Db() *gorm.DB {
	return entity.Db()
}

// UnscopedDb returns an unscoped database connection instance.
func UnscopedDb() *gorm.DB {
	return entity.Db().Unscoped()
}
