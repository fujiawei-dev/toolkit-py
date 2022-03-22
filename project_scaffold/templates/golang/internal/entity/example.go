{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

    "{{GOLANG_MODULE}}/pkg/rnd"
)

func init() {
	AddEntity(Example{})
}

type EmbeddedExample struct {
	Email string
}

type AnonymousEmbeddedExample struct {
	Anonymous bool
}

type AssociationExample struct {
	gorm.Model
	Number string
	UserID uint
}

type JoinsPreloadExample struct {
	gorm.Model
	Keyword string
}

// Example https://gorm.io/docs/models.html
// https://gorm.io/docs/models.html#Fields-Tags
type Example struct {
	ID uint `gorm:"primaryKey" json:"id" example:"1024"`

	UID string `gorm:"column:uid" json:"uid" example:"UID"`

	// Index
	IndexField int `gorm:"index" example:"21"`
	// UniqueField      int `gorm:"unique"`// SQLite doesn't support
	UniqueIndexField int `gorm:"uniqueIndex" example:"12"`

	// Basic DataType
	ShortStringField     string  `gorm:"column:short_string_field;size:8" example:"短字符串字段"`
	LongStringField      string  `gorm:"column:long_string_field;size:256" example:"长字符串字段"` // specifies column data size/length
	LongTextField        string  `gorm:"column:long_text_field;type:text" example:"超长文本字段"`
	IntegerField         int     `gorm:"type:int" example:"3"`
	// UnsignedIntegerField uint    `gorm:"type:uint;autoIncrement" example:"4"` // specifies column auto incremental
	Float64Field         float64 `gorm:"type:float" example:"3.1415926535"`
	Float32Field         float32 `gorm:"type:float" example:"3.14159"`
	BinaryField          []byte  `gorm:"type:bytes" example:"255"`

	// DefaultField     string `gorm:"default:golang.org" example:"默认字段"`
	// DefaultUID       string `gorm:"column:default_uid;default:uuid_generate_v3()" json:"default_uid" example:"默认SQL函数字段"`
	// DefaultGenerated string `gorm:"->;type:GENERATED ALWAYS AS (concat(uid,' ',default_uid));default:(-);" example:"默认组合生成字段"`

	NotNullField sql.NullBool `gorm:"not null" json:"-" example:"禁止空值字段"`
	CheckField   string       `gorm:"check:integer_field > 5" example:"验证值字段"` // https://gorm.io/docs/constraints.html
	CommentField string       `gorm:"comment" example:"数据库注释字段"`               // add comment for field when migration

	// Embedded Struct
	// When creating some data with associations, if its associations value is not zero-value,
	// those associations will be upserted, and its Hooks methods will be invoked.
	// AssociationExample AssociationExample
	// For anonymous fields, GORM will include its fields into its parent struct
	AnonymousEmbeddedExample
	// For a normal struct field, you can embed it with the tag embedded, for example:
	EmbeddedExample EmbeddedExample `gorm:"embedded"`
	// And you can use tag embeddedPrefix to add prefix to embedded fields’ db name, for example:
	EmbeddedExamplePrefix EmbeddedExample `gorm:"embedded;embeddedPrefix:embedded_"`

	// Field-Level Permission
	AllowReadAndCreate   string `gorm:"<-:create" example:"允许读和创建"`         // allow read and create
	AllowReadAndUpdate   string `gorm:"<-:update" example:"允许读和更新"`         // allow read and update
	AllowCreateAndUpdate string `gorm:"<-" example:"允许创建和更新"`               // allow read and write (create and update)
	ReadOnly             string `gorm:"->" example:"只读"`                    // readonly (disable write permission unless it configured )
	CreateOnly           string `gorm:"->:false;<-:create" example:"只允许创建"` // create only (disabled read from db)
	IgnoreWriteAndRead   string `gorm:"-" example:"忽略读写"`                   // ignore this field when write and read with struct
	IgnoreMigration      string `gorm:"migration" example:"忽略迁移"`           // // ignore this field when migration

	// Creating/Updating Time/Unix (Milli/Nano) Seconds Tracking
	CreatedAt             time.Time      `json:"created_at" example:"2020-07-14T16:20:00+08:00"` // Set to current time if it is zero on creating
	UpdatedAt             int            `json:"updated_at" example:"1678775400"`                // Set to current unix seconds on updating or if it is zero on creating
	UpdatedAtNanoSeconds  int64          `gorm:"autoUpdateTime:nano" example:"1678775400"`       // Use unix nano seconds as updating time
	UpdatedAtMilliSeconds int64          `gorm:"autoUpdateTime:milli" example:"1678775400"`      // Use unix milli seconds as updating time
	CreatedAtSeconds      int64          `gorm:"autoCreateTime" example:"1678775400"`            // Use unix seconds as creating time
	DeletedAt             gorm.DeletedAt `json:"-"`
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

// Create

// Create inserts a new row to the database.
func (m *Example) Create() error {
	return Db().Create(m).Error
}

// CreateSelectFields create a record and assign a value to the fields specified.
func (m *Example) CreateSelectFields() error {
	return Db().Select(
		"AllowReadAndCreate",
		"AllowCreateAndUpdate",
		"CreateOnly",
	).Create(m).Error
}

// CreateOmitFields create a record and ignore the values for fields passed to omit.
func (m *Example) CreateOmitFields() error {
	return Db().Omit(
		"AssociationExample",
		"UpdatedAtNanoSeconds",
		"UpdatedAtMilliSeconds",
		"CreatedAtSeconds",
	).Create(m).Error
}

// CreateOmitAssociations skip all associations.
func (m *Example) CreateOmitAssociations() error {
	return Db().Omit(clause.Associations).Create(m).Error
}

// Create efficiently insert large number of records
func (ms Examples) Create() error {
	if len(ms) < 5000 {
		return Db().Create(ms).Error
	}

	// specify batch size when creating
	return Db().CreateInBatches(ms, 5000).Error
}

// CreateOnConflict do something on conflict
func (m *Example) CreateOnConflict() (err error) {
	// https://gorm.io/docs/create.html#Upsert-On-Conflict
	return Db().Clauses(clause.OnConflict{
		Columns:      nil,
		Where:        clause.Where{},
		TargetWhere:  clause.Where{},
		OnConstraint: "",
		DoNothing:    false, // Do nothing on conflict
		DoUpdates:    nil,
		UpdateAll:    false,
	}).Create(m).Error
}

// BeforeCreate creates a random UID if needed before inserting a new row to the database.
func (m *Example) BeforeCreate(tx *gorm.DB) (err error) {
	if rnd.IsUID(m.UID, 'e') {
		return nil
	}

	m.UID = rnd.PPID('e')

	return
}

// Query

// First get the first record ordered by primary key
func (m *Example) First() (err error) {
	err = Db().First(m).Error
	return
}

// Last get the last record ordered by primary key
func (m *Example) Last() (err error) {
	err = Db().Last(m).Error
	return
}

// Take get one record, no specified order
func (m *Example) Take() (err error) {
	err = Db().Take(m).Error
	return
}

func (m *Example) FirstByPrimaryKey(id uint) (err error) {
	err = Db().First(m, id).Error
	return
}

func (m *Example) FindByID(id uint) (err error) {
	return m.FirstByPrimaryKey(id)
}

func (ms Examples) FindByPrimaryKeys(ids []uint) (err error) {
	err = Db().Find(&ms, ids).Error
	return
}

func (ms Examples) FindByIDs(ids []uint) (err error) {
	return ms.FindByPrimaryKeys(ids)
}

func (ms Examples) FindByStruct(m Example, fields ...interface{}) (err error) {
	err = Db().Where(&m, fields...).Find(&ms).Error
	return
}

func (ms Examples) FindByMap(where map[string]interface{}) (err error) {
	err = Db().Where(where).Find(&ms).Error
	return
}

func (ms *Examples) FindAll() (err error) {
	err = Db().Find(ms).Error
	return
}

func (m *Example) FindByShortStringField(s string) (err error) {
	err = Db().First(m, "short_string_field = ?", s).Error
	return
}

func (m *Example) FindBySelf() (err error) {
	err = Db().Where(m).First(m).Error
	return
}

func (m *Example) Exists() (bool, error) {
	err := m.FindBySelf()

	if err == nil {
		return true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, err
}

// FirstOrInitialize get first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (m *Example) FirstOrInitialize() (err error) {
	err = Db().FirstOrInit(m, m).Error
	return
}

// FirstOrCreate get first matched record or create a new one with given conditions (only works with struct, map conditions)
func (m *Example) FirstOrCreate() (err error) {
	// https://gorm.io/docs/advanced_query.html#FirstOrCreate
	err = Db().FirstOrCreate(m, m).Error
	return
}

// Pluck query single column from database and scan into a slice
func (m *Example) Pluck(field string, dst interface{}) (err error) {
	err = Db().Model(m).Pluck(field, &dst).Error
	return
}

// Update
// https://gorm.io/docs/update.html

// Save inserts a new row to the database or updates a row if the primary key already exists.
func (m *Example) Save() error {
	// Save will save all fields when performing the Updating SQL
	return Db().Save(m).Error
}

// Update single column
// When updating a single column with Update, it needs to have any conditions,
// or it will raise error ErrMissingWhereClause
func (m *Example) Update(column string, value interface{}) (err error) {
	err = Db().Where(m).Update(column, value).Error
	return
}

func (m *Example) Updates(n Example) (err error) {
	err = Db().Where(m).Updates(n).Error
	return
}

// UpdateColumn Without Hooks/Time Tracking
func (m *Example) UpdateColumn(column string, value interface{}) (err error) {
	err = Db().Where(m).UpdateColumn(column, value).Error
	return
}

// UpdateColumns Update multiple columns without Hooks/Time Tracking
func (m *Example) UpdateColumns(n Example) (err error) {
	err = Db().Where(m).UpdateColumns(n).Error
	return
}

// Delete marks the entity as deleted.
func (m *Example) Delete() error {
	return Db().Delete(m).Error
}

func (m *Example) DeletePermanently() error {
	return UnscopedDb().Delete(m).Error
}
