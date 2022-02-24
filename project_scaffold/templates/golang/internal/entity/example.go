{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Example struct {
	Id uint `gorm:"primary_key" json:"id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (Example) TableName() string {
	return "examples"
}

type Examples []Example

func (m *Example) CopyFrom(src interface{}) error {
	return copier.Copy(m, src)
}

func (m *Example) CopyTo(dst interface{}) error {
	return copier.Copy(dst, m)
}

// Create inserts a new row to the database.
func (m *Example) Create() error {
	return Db().Create(m).Error
}

// Save inserts a new row to the database or updates a row if the primary key already exists.
func (m *Example) Save() error {
	return Db().Save(m).Error
}
