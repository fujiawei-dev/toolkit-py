/*
 * @Date: 2022.03.18 15:37
 * @Description: Omit
 * @LastEditors: Rustle Karl
 * @LastEditTime: 2022.03.18 15:37
 */

package entity

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"{{GOLANG_MODULE}}/internal/acl"
)

func init() {
	AddEntity(OperationLog{})
}

type OperationLog struct {
	ID uint `gorm:"primary_key" json:"id"`

	UserID uint `json:"-"`
	User   User `gorm:"foreignkey:UserID;association_foreignkey:id" json:"-"`

	Resource acl.Resource `gorm:"size:32" json:"resource"` // 操作资源
	Action   acl.Action   `gorm:"size:32" json:"action"`   // 操作行为
	Allow    bool         `gorm:"size:32" json:"allow"`    // 操作是否被允许

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (OperationLog) TableName() string {
	return "operation_logs"
}

type OperationLogs []OperationLog

func NewOperationLog(userID uint, resource acl.Resource, action acl.Action, allow bool) OperationLog {
	return OperationLog{UserID: userID, Resource: resource, Action: action, Allow: allow}
}

func (m *OperationLog) CopyFrom(src interface{}) error {
	return copier.Copy(m, src)
}

func (m *OperationLog) CopyTo(dst interface{}) error {
	return copier.Copy(dst, m)
}

// Create

// Create inserts a new row to the database.
func (m *OperationLog) Create() error {
	return Db().Create(m).Error
}

// Query

// First get the first record ordered by primary key
func (m *OperationLog) First() (err error) {
	err = Db().First(m).Error
	return
}

// Last get the last record ordered by primary key
func (m *OperationLog) Last() (err error) {
	err = Db().Last(m).Error
	return
}

// Take get one record, no specified order
func (m *OperationLog) Take() (err error) {
	err = Db().Take(m).Error
	return
}

func (ms *OperationLogs) FindAll() (err error) {
	err = Db().Find(ms).Error
	return
}

func (m *OperationLog) FindByID(id uint) (err error) {
	return Db().First(m, id).Error
}

func (m *OperationLog) FindBySelf() (err error) {
	err = Db().Where(m).First(m).Error
	return
}

func (m *OperationLog) Exists() (bool, error) {
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
func (m *OperationLog) Pluck(field string, dst interface{}) (err error) {
	err = Db().Model(m).Pluck(field, &dst).Error
	return
}

// Update
// https://gorm.io/docs/update.html

// Save inserts a new row to the database or updates a row if the primary key already exists.
func (m *OperationLog) Save() error {
	// Save will save all fields when performing the Updating SQL
	return Db().Save(m).Error
}

// Update single column
// When updating a single column with Update, it needs to have any conditions,
// or it will raise error ErrMissingWhereClause
func (m *OperationLog) Update(column string, value interface{}) (err error) {
	err = Db().Where(m).Update(column, value).Error
	return
}

func (m *OperationLog) Updates(values interface{}) (err error) {
	err = Db().Where(m).Updates(values).Error
	return
}

// UpdateColumn Without Hooks/Time Tracking
func (m *OperationLog) UpdateColumn(column string, value interface{}) (err error) {
	err = Db().Where(m).UpdateColumn(column, value).Error
	return
}

// UpdateColumns Update multiple columns without Hooks/Time Tracking
func (m *OperationLog) UpdateColumns(values interface{}) (err error) {
	err = Db().Where(m).UpdateColumns(values).Error
	return
}

// Delete marks the entity as deleted.
func (m *OperationLog) Delete() error {
	return Db().Delete(m).Error
}

func (m *OperationLog) DeletePermanently() error {
	return UnscopedDb().Delete(m).Error
}
