{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"gorm.io/gorm"
)

var gProvider DbProvider

type DbProvider interface {
	Db() *gorm.DB
}

func SetDbProvider(provider DbProvider) {
	gProvider = provider
}

func HasDbProvider() bool {
	return gProvider != nil
}

// Db returns a database connection.
func Db() *gorm.DB {
	if !HasDbProvider() {
		panic("entity: database not connected")
	}

	return gProvider.Db()
}

// UnscopedDb returns an unscoped database connection.
func UnscopedDb() *gorm.DB {
	return Db().Unscoped()
}
