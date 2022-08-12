package entity

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"{{ main_module }}/pkg/rnd"
)

func init() {
	AddEntity({{ factory.entity|title }}Embedded{})
	AddEntity({{ factory.entity|title }}{})
}

type {{ factory.entity|title }}Embedded struct {
	gorm.Model
	Email string `json:"email" example:"who@gmail.com"`
}

func ({{ factory.entity|title }}Embedded) TableName() string {
	return "{{ factory.entity }}s_embedded"
}

type {{ factory.entity|title }} struct {
	ID uint `gorm:"primary_key" json:"id"`

	// ID easy to traverse, so UID is necessary when get/put/delete
	UID string `gorm:"column:uid" json:"uid" example:"UID"`

	When  string `gorm:"size:128;" json:"when"`
	Where string `gorm:"size:128;" json:"where"`
	Who   string `gorm:"size:128;" json:"who"`
	What  string `gorm:"size:128;" json:"what"`
	How   string `gorm:"size:128;" json:"how"`

	{{ factory.entity|title }}EmbeddedID uint                   `json:"-"`
	{{ factory.entity|title }}Embedded   {{ factory.entity|title }}Embedded `gorm:"foreignkey:{{ factory.entity|title }}EmbeddedID;association_foreignkey:id" json:"-"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (*{{ factory.entity|title }}) TableName() string {
	return "{{ factory.entity }}s"
}

type {{ factory.entity|title }}s []{{ factory.entity|title }}

func (m *{{ factory.entity|title }}) CopyFrom(src interface{}) error {
	return copier.Copy(m, src)
}

func (m *{{ factory.entity|title }}) CopyTo(dst interface{}) error {
	return copier.Copy(dst, m)
}

// Create

// Create inserts a new row to the database.
func (m *{{ factory.entity|title }}) Create() error {
	return Db().Create(m).Error
}

// BeforeCreate creates a random UID if needed before inserting a new row to the database.
func (m *{{ factory.entity|title }}) BeforeCreate(tx *gorm.DB) (err error) {
	if rnd.IsUID(m.UID, 'e') {
		return nil
	}

	m.UID = rnd.PPID('e')

	return
}

// Query

// First get the first record ordered by primary key
func (m *{{ factory.entity|title }}) First() (err error) {
	err = Db().First(m).Error
	return
}

// Last get the last record ordered by primary key
func (m *{{ factory.entity|title }}) Last() (err error) {
	err = Db().Last(m).Error
	return
}

// Take get one record, no specified order
func (m *{{ factory.entity|title }}) Take() (err error) {
	err = Db().Take(m).Error
	return
}

func (ms *{{ factory.entity|title }}s) FindAll() (err error) {
	err = Db().Find(ms).Error
	return
}

func (m *{{ factory.entity|title }}) FindByID(id uint) (err error) {
	return Db().First(m, id).Error
}

func (m *{{ factory.entity|title }}) FindBySelf() (err error) {
	err = Db().Where(m).First(m).Error
	return
}

func (m *{{ factory.entity|title }}) Exists() (bool, error) {
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
func (m *{{ factory.entity|title }}) Pluck(field string, dst interface{}) (err error) {
	err = Db().Model(m).Pluck(field, &dst).Error
	return
}

// Update
// https://gorm.io/docs/update.html

// Save inserts a new row to the database or updates a row if the primary key already exists.
func (m *{{ factory.entity|title }}) Save() error {
	// Save will save all fields when performing the Updating SQL
	return Db().Save(m).Error
}

// Update single column
// When updating a single column with Update, it needs to have any conditions,
// or it will raise error ErrMissingWhereClause
func (m *{{ factory.entity|title }}) Update(column string, value interface{}) (err error) {
	err = Db().Model(m).Where(m).Update(column, value).Error
	return
}

func (m *{{ factory.entity|title }}) Updates(values interface{}) (err error) {
	err = Db().Model(m).Where(m).Updates(values).Error
	return
}

// UpdateColumn Without Hooks/Time Tracking
func (m *{{ factory.entity|title }}) UpdateColumn(column string, value interface{}) (err error) {
	err = Db().Model(m).Where(m).UpdateColumn(column, value).Error
	return
}

// UpdateColumns Update multiple columns without Hooks/Time Tracking
func (m *{{ factory.entity|title }}) UpdateColumns(values interface{}) (err error) {
	err = Db().Model(m).Where(m).UpdateColumns(values).Error
	return
}

// Delete marks the factory.entity as deleted.
func (m *{{ factory.entity|title }}) Delete() error {
	return Db().Delete(m).Error
}

func (m *{{ factory.entity|title }}) DeletePermanently() error {
	return UnscopedDb().Delete(m).Error
}
