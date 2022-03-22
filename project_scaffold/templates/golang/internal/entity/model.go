{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

    "{{GOLANG_MODULE}}/pkg/rnd"
)

func init() {
	// AddEntity(BasicModel{})
}

// BasicModel basic minimal model example
type BasicModel struct {
	ID uint `gorm:"primary_key" json:"id"`

	// ID easy to traverse, so UID is necessary when get/put/delete
	UID string `gorm:"column:uid" json:"uid" example:"UID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (BasicModel) TableName() string {
	return "basic_models"
}

type BasicModels []BasicModel

func (m *BasicModel) CopyFrom(src interface{}) error {
	return copier.Copy(m, src)
}

func (m *BasicModel) CopyTo(dst interface{}) error {
	return copier.Copy(dst, m)
}

// Create

// Create inserts a new row to the database.
func (m *BasicModel) Create() error {
	return Db().Create(m).Error
}

// BeforeCreate creates a random UID if needed before inserting a new row to the database.
func (m *BasicModel) BeforeCreate(tx *gorm.DB) (err error) {
	if rnd.IsUID(m.UID, 'e') {
		return nil
	}

	m.UID = rnd.PPID('e')

	return
}

// Query

// First get the first record ordered by primary key
func (m *BasicModel) First() (err error) {
	err = Db().First(m).Error
	return
}

// Last get the last record ordered by primary key
func (m *BasicModel) Last() (err error) {
	err = Db().Last(m).Error
	return
}

// Take get one record, no specified order
func (m *BasicModel) Take() (err error) {
	err = Db().Take(m).Error
	return
}

func (ms *BasicModels) FindAll() (err error) {
	err = Db().Find(ms).Error
	return
}

func (m *BasicModel) FindByID(id uint) (err error) {
	return Db().First(m, id).Error
}

func (m *BasicModel) FindBySelf() (err error) {
	err = Db().Where(m).First(m).Error
	return
}

func (m *BasicModel) Exists() (bool, error) {
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
func (m *BasicModel) Pluck(field string, dst interface{}) (err error) {
	err = Db().Model(m).Pluck(field, &dst).Error
	return
}

// Update
// https://gorm.io/docs/update.html

// Save inserts a new row to the database or updates a row if the primary key already exists.
func (m *BasicModel) Save() error {
	// Save will save all fields when performing the Updating SQL
	return Db().Save(m).Error
}

// Update single column
// When updating a single column with Update, it needs to have any conditions,
// or it will raise error ErrMissingWhereClause
func (m *BasicModel) Update(column string, value interface{}) (err error) {
	err = Db().Where(m).Update(column, value).Error
	return
}

func (m *BasicModel) Updates(values interface{}) (err error) {
	err = Db().Where(m).Updates(values).Error
	return
}

// UpdateColumn Without Hooks/Time Tracking
func (m *BasicModel) UpdateColumn(column string, value interface{}) (err error) {
	err = Db().Where(m).UpdateColumn(column, value).Error
	return
}

// UpdateColumns Update multiple columns without Hooks/Time Tracking
func (m *BasicModel) UpdateColumns(values interface{}) (err error) {
	err = Db().Where(m).UpdateColumns(values).Error
	return
}

// Delete marks the entity as deleted.
func (m *BasicModel) Delete() error {
	return Db().Delete(m).Error
}

func (m *BasicModel) DeletePermanently() error {
	return UnscopedDb().Delete(m).Error
}
