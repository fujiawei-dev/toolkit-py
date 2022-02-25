{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"database/sql"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type EmbeddedExample struct {
	Email string
}

type AnonymousEmbeddedExample struct {
	Anonymous bool
}

// Example https://gorm.io/docs/models.html
// https://gorm.io/docs/models.html#Fields-Tags
type Example struct {
	Id uint `gorm:"primary_key" json:"id"`

	// Index
	IndexField int `gorm:"index"`
	// UniqueField      int `gorm:"unique"`// SQLite doesn't support
	UniqueIndexField int `gorm:"uniqueIndex"`

	// Basic DataType
	ShortStringField     string  `gorm:"column:short_string_field;size:8"`
	LongStringField      string  `gorm:"column:long_string_field;size:256"` // specifies column data size/length
	LongTextField        string  `gorm:"column:long_text_field;type:text"`
	IntegerField         int     `gorm:"type:int"`
	UnsignedIntegerField uint    `gorm:"type:uint"`
	Float64Field         float64 `gorm:"type:float"`
	Float32Field         float32 `gorm:"type:float"`
	BinaryField          []byte  `gorm:"type:bytes"`

	DefaultField string       `gorm:"default:value"`
	NotNullField sql.NullBool `gorm:"not null" json:"-"`
	CheckField   int          `gorm:"check:integer_field > 5"` // https://gorm.io/docs/constraints.html
	CommentField string       `gorm:"comment"`                 // add comment for field when migration

	// Embedded Struct
	// For anonymous fields, GORM will include its fields into its parent struct
	AnonymousEmbeddedExample

	// For a normal struct field, you can embed it with the tag embedded, for example:
	EmbeddedExample EmbeddedExample `gorm:"embedded"`

	// And you can use tag embeddedPrefix to add prefix to embedded fields’ db name, for example:
	EmbeddedExamplePrefix EmbeddedExample `gorm:"embedded;embeddedPrefix:embedded_"`

	// Field-Level Permission
	AllowReadAndCreate   string `gorm:"<-:create"`          // allow read and create
	AllowReadAndUpdate   string `gorm:"<-:update"`          // allow read and update
	AllowCreateAndUpdate string `gorm:"<-"`                 // allow read and write (create and update)
	ReadOnly             string `gorm:"->"`                 // readonly (disable write permission unless it configured )
	CreateOnly           string `gorm:"->:false;<-:create"` // create only (disabled read from db)
	IgnoreWriteAndRead   string `gorm:"-"`                  // ignore this field when write and read with struct
	IgnoreMigration      string `gorm:"migration"`          // // ignore this field when migration

	// Creating/Updating Time/Unix (Milli/Nano) Seconds Tracking
	CreatedAt             time.Time      `json:"created_at"`           // Set to current time if it is zero on creating
	UpdatedAt             int            `json:"updated_at"`           // Set to current unix seconds on updating or if it is zero on creating
	UpdatedAtNanoSeconds  int64          `gorm:"autoUpdateTime:nano"`  // Use unix nano seconds as updating time
	UpdatedAtMilliSeconds int64          `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
	CreatedAtSeconds      int64          `gorm:"autoCreateTime"`       // Use unix seconds as creating time
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

// Create inserts a new row to the database.
func (m *Example) Create() error {
	return Db().Create(m).Error
}

// Save inserts a new row to the database or updates a row if the primary key already exists.
func (m *Example) Save() error {
	return Db().Save(m).Error
}

// Delete marks the entity as deleted.
func (m *Example) Delete() error {
	return Db().Delete(m).Error
}

func (m *Example) FindById(id uint) (err error) {
	err = Db().First(m, id).Error
	return
}

func (m *Example) FindByShortStringField(s string) (err error) {
	err = Db().First(m, "short_string_field = ?", s).Error
	return
}
