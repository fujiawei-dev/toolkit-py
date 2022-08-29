package entity

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"{{ main_module }}/pkg/rnd"
)

func init() {
	AddEntity({{ factory.entity_slug.pascal_case }}Embedded{})
	AddEntity({{ factory.entity_slug.pascal_case }}{})
}

type {{ factory.entity_slug.pascal_case }}Embedded struct {
	gorm.Model
	Email string `gorm:"type:varchar(255)" json:"email" example:"who@gmail.com"`
}

func ({{ factory.entity_slug.pascal_case }}Embedded) TableName() string {
	return "{{ factory.entity }}s_embedded"
}

type {{ factory.entity_slug.pascal_case }} struct {
	ID uint `gorm:"primary_key" json:"id"`

	// ID easy to traverse, so UID is necessary when get/put/delete
	UID string `gorm:"column:uid" json:"uid" example:"UID"`

	When  string `gorm:"type:varchar(255)" json:"when"`
	Where string `gorm:"size:128;" json:"where"`
	Who   string `gorm:"size:128;" json:"who"`
	What  string `gorm:"size:128;" json:"what"`
	How   string `gorm:"size:128;" json:"how"`

	{{ factory.entity_slug.pascal_case }}EmbeddedID uint                   `json:"-"`
	{{ factory.entity_slug.pascal_case }}Embedded   {{ factory.entity_slug.pascal_case }}Embedded `gorm:"foreignkey:{{ factory.entity_slug.pascal_case }}EmbeddedID;association_foreignkey:id" json:"-"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func ({{ factory.entity_slug.pascal_case }}) TableName() string {
	return "{{ factory.entity }}s"
}

type {{ factory.entity_slug.pascal_case }}s []{{ factory.entity_slug.pascal_case }}

func (m *{{ factory.entity_slug.pascal_case }}) CopyFrom(src interface{}) error {
	return copier.Copy(m, src)
}

func (m *{{ factory.entity_slug.pascal_case }}) CopyTo(dst interface{}) error {
	return copier.Copy(dst, m)
}

// Create

// Create inserts a new row to the database.
func (m *{{ factory.entity_slug.pascal_case }}) Create() error {
	return Db().Create(m).Error
}

// BeforeCreate creates a random UID if needed before inserting a new row to the database.
func (m *{{ factory.entity_slug.pascal_case }}) BeforeCreate(tx *gorm.DB) (err error) {
	if rnd.IsUID(m.UID, 'e') {
		return nil
	}

	m.UID = rnd.PPID('e')

	return
}

// Query

// First get the first record ordered by primary key
func (m *{{ factory.entity_slug.pascal_case }}) First() (err error) {
	err = Db().First(m).Error
	return
}

// Last get the last record ordered by primary key
func (m *{{ factory.entity_slug.pascal_case }}) Last() (err error) {
	err = Db().Last(m).Error
	return
}

// Take get one record, no specified order
func (m *{{ factory.entity_slug.pascal_case }}) Take() (err error) {
	err = Db().Take(m).Error
	return
}

func (ms *{{ factory.entity_slug.pascal_case }}s) FindAll() (err error) {
	err = Db().Find(ms).Error
	return
}

func (m *{{ factory.entity_slug.pascal_case }}) FindByID(id uint) (err error) {
	return Db().First(m, id).Error
}

func (m *{{ factory.entity_slug.pascal_case }}) FindBySelf() (err error) {
	err = Db().Where(m).First(m).Error
	return
}

func (m *{{ factory.entity_slug.pascal_case }}) Exists() (bool, error) {
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
func (m *{{ factory.entity_slug.pascal_case }}) Pluck(field string, dst interface{}) (err error) {
	err = Db().Model(m).Pluck(field, &dst).Error
	return
}

// Update
// https://gorm.io/docs/update.html

// Save inserts a new row to the database or updates a row if the primary key already exists.
func (m *{{ factory.entity_slug.pascal_case }}) Save() error {
	// Save will save all fields when performing the Updating SQL
	return Db().Save(m).Error
}

// Update single column
// When updating a single column with Update, it needs to have any conditions,
// or it will raise error ErrMissingWhereClause
func (m *{{ factory.entity_slug.pascal_case }}) Update(column string, value interface{}) (err error) {
	err = Db().Model(m).Where(m).Update(column, value).Error
	return
}

func (m *{{ factory.entity_slug.pascal_case }}) Updates(values interface{}) (err error) {
	err = Db().Model(m).Where(m).Updates(values).Error
	return
}

// UpdateColumn Without Hooks/Time Tracking
func (m *{{ factory.entity_slug.pascal_case }}) UpdateColumn(column string, value interface{}) (err error) {
	err = Db().Model(m).Where(m).UpdateColumn(column, value).Error
	return
}

// UpdateColumns Update multiple columns without Hooks/Time Tracking
func (m *{{ factory.entity_slug.pascal_case }}) UpdateColumns(values interface{}) (err error) {
	err = Db().Model(m).Where(m).UpdateColumns(values).Error
	return
}

// Delete marks the factory.entity as deleted.
func (m *{{ factory.entity_slug.pascal_case }}) Delete() error {
	return Db().Delete(m).Error
}

func (m *{{ factory.entity_slug.pascal_case }}) DeletePermanently() error {
	return UnscopedDb().Delete(m).Error
}
