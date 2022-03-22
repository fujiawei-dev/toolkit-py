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
	AddEntity(EntityTemplateEmbedded{})
	AddEntity(EntityTemplate{})
}

type EntityTemplateEmbedded struct {
	gorm.Model
	Email string `json:"email" example:"who@gmail.com"`
}

func (EntityTemplateEmbedded) TableName() string {
	return "entity_templates_embedded"
}

// EntityTemplate basic minimal entity example
type EntityTemplate struct {
	ID uint `gorm:"primary_key" json:"id"`

	// ID easy to traverse, so UID is necessary when get/put/delete
	UID string `gorm:"column:uid" json:"uid" example:"UID"`

	When  string `gorm:"size:128;" json:"when"`
	Where string `gorm:"size:128;" json:"where"`
	Who   string `gorm:"size:128;" json:"who"`
	What  string `gorm:"size:128;" json:"what"`
	How   string `gorm:"size:128;" json:"how"`

	EntityTemplateEmbeddedID uint                   `json:"-"`
	EntityTemplateEmbedded   EntityTemplateEmbedded `gorm:"foreignkey:EntityTemplateEmbeddedID;association_foreignkey:id" json:"-"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (EntityTemplate) TableName() string {
	return "entity_templates"
}

type EntityTemplates []EntityTemplate

func (m *EntityTemplate) CopyFrom(src interface{}) error {
	return copier.Copy(m, src)
}

func (m *EntityTemplate) CopyTo(dst interface{}) error {
	return copier.Copy(dst, m)
}

// Create

// Create inserts a new row to the database.
func (m *EntityTemplate) Create() error {
	return Db().Create(m).Error
}

// BeforeCreate creates a random UID if needed before inserting a new row to the database.
func (m *EntityTemplate) BeforeCreate(tx *gorm.DB) (err error) {
	if rnd.IsUID(m.UID, 'e') {
		return nil
	}

	m.UID = rnd.PPID('e')

	return
}

// Query

// First get the first record ordered by primary key
func (m *EntityTemplate) First() (err error) {
	err = Db().First(m).Error
	return
}

// Last get the last record ordered by primary key
func (m *EntityTemplate) Last() (err error) {
	err = Db().Last(m).Error
	return
}

// Take get one record, no specified order
func (m *EntityTemplate) Take() (err error) {
	err = Db().Take(m).Error
	return
}

func (ms *EntityTemplates) FindAll() (err error) {
	err = Db().Find(ms).Error
	return
}

func (m *EntityTemplate) FindByID(id uint) (err error) {
	return Db().First(m, id).Error
}

func (m *EntityTemplate) FindBySelf() (err error) {
	err = Db().Where(m).First(m).Error
	return
}

func (m *EntityTemplate) Exists() (bool, error) {
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
func (m *EntityTemplate) Pluck(field string, dst interface{}) (err error) {
	err = Db().Model(m).Pluck(field, &dst).Error
	return
}

// Update
// https://gorm.io/docs/update.html

// Save inserts a new row to the database or updates a row if the primary key already exists.
func (m *EntityTemplate) Save() error {
	// Save will save all fields when performing the Updating SQL
	return Db().Save(m).Error
}

// Update single column
// When updating a single column with Update, it needs to have any conditions,
// or it will raise error ErrMissingWhereClause
func (m *EntityTemplate) Update(column string, value interface{}) (err error) {
	err = Db().Where(m).Update(column, value).Error
	return
}

func (m *EntityTemplate) Updates(values interface{}) (err error) {
	err = Db().Where(m).Updates(values).Error
	return
}

// UpdateColumn Without Hooks/Time Tracking
func (m *EntityTemplate) UpdateColumn(column string, value interface{}) (err error) {
	err = Db().Where(m).UpdateColumn(column, value).Error
	return
}

// UpdateColumns Update multiple columns without Hooks/Time Tracking
func (m *EntityTemplate) UpdateColumns(values interface{}) (err error) {
	err = Db().Where(m).UpdateColumns(values).Error
	return
}

// Delete marks the entity as deleted.
func (m *EntityTemplate) Delete() error {
	return Db().Delete(m).Error
}

func (m *EntityTemplate) DeletePermanently() error {
	return UnscopedDb().Delete(m).Error
}
