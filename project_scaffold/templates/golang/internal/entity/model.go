{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// Model 基本模型
type Model struct {
	ID uint `gorm:"primary_key"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (m *Model) CopyFrom(src interface{}) error {
	return copier.Copy(m, src)
}

func (m *Model) CopyTo(dst interface{}) error {
	return copier.Copy(dst, m)
}

// Create

// Create inserts a new row to the database.
func (m *Model) Create() error {
	return Db().Create(m).Error
}

// Query

// First get the first record ordered by primary key
func (m *Model) First() (err error) {
	err = Db().First(m).Error
	return
}

// Last get the last record ordered by primary key
func (m *Model) Last() (err error) {
	err = Db().Last(m).Error
	return
}

func (m *Model) FindByID(id uint) (err error) {
	return Db().First(m, id).Error
}

func (m *Model) FindBySelf() (err error) {
	err = Db().Where(m).First(m).Error
	return
}

func (m *Model) Exists() (bool, error) {
	err := m.FindBySelf()

	if err == nil {
		return true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, err
}

// Pluck query single column from database and scan into a slice
func (m *Model) Pluck(field string, dst interface{}) (err error) {
	err = Db().Model(m).Pluck(field, &dst).Error
	return
}

// Update
// https://gorm.io/docs/update.html

// Save inserts a new row to the database or updates a row if the primary key already exists.
func (m *Model) Save() error {
	// Save will save all fields when performing the Updating SQL
	return Db().Save(m).Error
}

// Update single column
// When updating a single column with Update, it needs to have any conditions,
// or it will raise error ErrMissingWhereClause
func (m *Model) Update(column string, value interface{}) (err error) {
	err = Db().Where(m).Update(column, value).Error
	return
}

func (m *Model) Updates(n Model) (err error) {
	err = Db().Where(m).Updates(n).Error
	return
}

// UpdateColumn Without Hooks/Time Tracking
func (m *Model) UpdateColumn(column string, value interface{}) (err error) {
	err = Db().Where(m).UpdateColumn(column, value).Error
	return
}

// UpdateColumns Update multiple columns without Hooks/Time Tracking
func (m *Model) UpdateColumns(n Model) (err error) {
	err = Db().Where(m).UpdateColumns(n).Error
	return
}

// Delete marks the entity as deleted.
func (m *Model) Delete() error {
	return Db().Delete(m).Error
}

func (m *Model) DeletePermanently() error {
	return UnscopedDb().Delete(m).Error
}
