{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"gorm.io/gorm"

	"{{GOLANG_MODULE}}/internal/config"
	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/event"
)

var log = event.Log
var conf = config.Conf()

const (
	MySQL  = "mysql"
	SQLite = "sqlite3"
	Or     = "|"
)

// MaxResults Max result limit for queries.
const MaxResults = 1000

// SearchRadius About 1km ('good enough' for now)
const SearchRadius = 0.009

// Query searches given an originals path and a db instance.
type Query struct {
	db *gorm.DB
}

// SearchCount is the total number of search hits.
type SearchCount struct {
	Total int
}

// New returns a new Query type with a given path and db instance.
func New(db *gorm.DB) *Query {
	q := &Query{
		db: db,
	}

	return q
}

// Db returns a database connection instance.
func Db() *gorm.DB {
	return entity.Db()
}

// UnscopedDb returns an unscoped database connection instance.
func UnscopedDb() *gorm.DB {
	return entity.Db().Unscoped()
}
