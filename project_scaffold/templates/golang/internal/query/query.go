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

const (
	Or = "|"
)

type Search struct {
	LikeQ string `json:"like_q" form:"like_q" url:"like_q" example:"Fuzzy query words, multiple query words are separated by |, golang|cpp|rust"`
	MustQ string `json:"must_q" form:"must_q" url:"must_q" example:"Precise query words, multiple query words are separated by |, golang|cpp|rust"`
	NotQ  string `json:"not_q" form:"not_q" url:"not_q" example:"Not query words, multiple query words are separated by |, golang|cpp|rust"`

	TimeBegin string `json:"time_begin" form:"time_begin" url:"time_begin" example:"2022-01-01"`
	TimeEnd   string `json:"time_end" form:"time_end" url:"time_end" example:"2022-12-31"`
}

// New returns a new Query type with a given path and db instance.

// Db returns a database connection instance.
func Db() *gorm.DB {
	return entity.Db()
}

// UnscopedDb returns an unscoped database connection instance.
func UnscopedDb() *gorm.DB {
	return entity.Db().Unscoped()
}
