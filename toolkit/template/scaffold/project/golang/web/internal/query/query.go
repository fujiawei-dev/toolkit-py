package query

import (
	"gorm.io/gorm"

	"{{ main_module }}/internal/config"
	"{{ main_module }}/internal/entity"
	"{{ main_module }}/internal/event"
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
