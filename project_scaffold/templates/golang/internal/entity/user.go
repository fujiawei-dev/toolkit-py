{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/form"
)

func init() {
	AddEntity(User{})
}

// User represents a person that may optionally log in as user.
type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Username string `gorm:"column:username;size:128;" json:"username"`
	Password string `gorm:"-" json:"-"`

	Role acl.Role `gorm:"size:32;default:default;" json:"-"`

	Disabled bool `json:"disabled"`

	LoginAttempts int       `json:"-" yaml:"-"` // 登录密码尝试次数
	LoginAt       time.Time `json:"-"`          // 最近一次登录时间

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (User) TableName() string {
	return "users"
}

type Users []User

// Admin is the default admin user.
var Admin = User{
	ID:       1,
	Username: "admin",
	Password: "admin",
	Role:     acl.RoleAdmin,
	Disabled: false,
	LoginAt:  time.Now(),
}

// Guest is the default guest user.
var Guest = User{
	ID:       2,
	Username: "guest",
	Password: "guest",
	Role:     acl.RoleGuest,
	Disabled: false,
	LoginAt:  time.Now(),
}

func (m User) Invalid() bool {
	return m.ID == 0 || m.Username == "" || m.Disabled
}

// CreateDefaultUsers initializes the database with default user accounts.
func CreateDefaultUsers() {
	if user := FirstOrCreateUser(&Admin); user != nil {
		Admin = *user
		Admin.InitPassword(Admin.Password)
	}

	if user := FirstOrCreateUser(&Guest); user != nil {
		Guest = *user
		Guest.InitPassword(Guest.Password)
	}
}

func (m *User) CopyFrom(src interface{}) error {
	return copier.Copy(m, src)
}

func (m *User) CopyTo(dst interface{}) error {
	return copier.Copy(dst, m)
}

// Create new entity in the database.
func (m *User) Create() error {
	return Db().Create(m).Error
}

// Save entity properties.
func (m *User) Save() error {
	return Db().Save(m).Error
}

// CreateWithPassword Creates User with Password in db transaction.
func CreateWithPassword(f form.User) error {
	u := &User{Username: f.Username, LoginAt: time.Now()}

	if len(f.Password) < 4 {
		return fmt.Errorf("user: new password for %s must be at least 4 characters", u.Username)
	}

	if err := u.Validate(); err != nil {
		return err
	}

	return Db().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(u).Error; err != nil {
			return err
		}

		pw := NewPassword(u.ID, f.Password)
		if err := tx.Create(&pw).Error; err != nil {
			return err
		}

		log.Printf("user: created user %s with id %d", u.Username, u.ID)

		return nil
	})
}

// FirstOrCreateUser returns an existing row, inserts a new row, or nil in case of errors.
func FirstOrCreateUser(m *User) *User {
	if err := Db().Where("id = ?", m.ID).Attrs(m).FirstOrCreate(m).Error; err != nil {
		log.Printf("%s: %v", m.TableName(), err)
		return nil
	}
	return m
}

// FindUserByUsername returns an existing user or nil if not found.
func FindUserByUsername(username string) (result User, err error) {
	err = Db().Where("username = ?", username).First(&result).Error
	return
}

// FindUserByID returns an existing user or nil if not found.
func FindUserByID(id uint) (result User, err error) {
	err = Db().Where("id = ?", id).First(&result).Error
	return
}

// Delete marks the entity as deleted.
func (m *User) Delete() error {
	if m.ID <= 3 {
		return fmt.Errorf("%s: can't delete default entity", m.TableName())
	}

	return Db().Delete(m).Error
}

// Deleted tests if the entity is marked as deleted.
func (m *User) Deleted() bool {
	return m.DeletedAt.Valid
}

// SetPassword sets a new password stored as hash.
func (m *User) SetPassword(password string) error {
	if len(password) < 4 {
		return fmt.Errorf("%s: new password for %s must be at least 4 characters", m.TableName(), m.Username)
	}

	pw := NewPassword(m.ID, password)

	return pw.Save()
}

// InitPassword sets the initial user password stored as hash.
func (m *User) InitPassword(password string) {
	if password == "" {
		return
	}

	if FindPassword(m.ID) != nil {
		return
	}

	pw := NewPassword(m.ID, password)

	if err := pw.Save(); err != nil {
		log.Printf("%s: %v", pw.TableName(), err)
	}
}

// InvalidPassword returns true if the given password does not match the hash.
func (m *User) InvalidPassword(password string) bool {
	if password == "" {
		return true
	}

	pw := FindPassword(m.ID)
	if pw == nil {
		return false
	}

	if pw.InvalidPassword(password) {
		if err := Db().Model(m).UpdateColumn("login_attempts",
			gorm.Expr("login_attempts + ?", 1)).Error; err != nil {
			log.Printf("%s: %s (update login attempts)", m.TableName(), err)
		}

		return true
	}

	if err := Db().Model(m).Updates(map[string]interface{}{
		"login_attempts": 0, "login_at": time.Now(),
	}).Error; err != nil {
		log.Printf("%s: %s (update last login)", m.TableName(), err)
	}

	return false
}

// Validate Makes sure username is unique and meet requirements. Returns error if any property is invalid
func (m *User) Validate() error {
	if m.Username == "" {
		return fmt.Errorf("%s: username must not be empty", m.TableName())
	}

	if len(m.Username) < 4 {
		return fmt.Errorf("%s: username must be at least 4 characters", m.TableName())
	}

	var err error
	var resultName = User{}

	if err = Db().Where("username = ? AND id <> ?", m.Username, m.ID).First(&resultName).Error; err == nil {
		return fmt.Errorf("%s: username already exists", m.TableName())
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
