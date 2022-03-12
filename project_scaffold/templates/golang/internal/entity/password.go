{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	AddEntity(Password{})
}

// Password represents a password hash.
type Password struct {
	UserID uint   `gorm:"primaryKey" json:"user_id"`
	Hash   string `copier:"-" gorm:"type:VARBINARY(255);" json:"hash"`

	// https://stackoverflow.com/questions/64806478/save-is-trying-to-update-created-at-column
	CreatedAt time.Time `gorm:"<-:create" copier:"-" json:"CreatedAt"`
	UpdatedAt time.Time `copier:"-" json:"UpdatedAt"`
}

func (Password) TableName() string {
	return "passwords"
}

// NewPassword creates a new password instance.
func NewPassword(userID uint, password string) Password {
	if userID == 0 {
		panic("auth: can't set password without user_id.")
	}

	m := Password{UserID: userID}

	if password != "" {
		if err := m.SetPassword(password); err != nil {
			log.Printf("auth: failed setting password for %s", userID)
		}
	}

	return m
}

// SetPassword sets a new password stored as hash.
func (m *Password) SetPassword(password string) error {
	// https://stackoverflow.com/questions/69567892/bcrypt-takes-a-lot-of-time-in-go
	if bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8); err != nil {
		return err
	} else {
		m.Hash = string(bytes)
		return nil
	}
}

// InvalidPassword returns true if the given password does not match the hash.
func (m *Password) InvalidPassword(password string) bool {
	if m.Hash == "" && password == "" {
		return false
	}
	// https://stackoverflow.com/questions/49437359/why-bcrypt-library-comparehashandpassword-method-is-slow
	err := bcrypt.CompareHashAndPassword([]byte(m.Hash), []byte(password))
	return err != nil
}

// Create inserts a new row to the database.
func (m *Password) Create() error {
	return Db().Create(m).Error
}

// Save inserts a new row to the database or updates a row if the primary key already exists.
func (m *Password) Save() error {
	return Db().Save(m).Error
}

// FindPassword returns an entity pointer if exists.
func FindPassword(userID uint) *Password {
	result := Password{}

	if err := Db().Where("user_id = ?", userID).First(&result).Error; err == nil {
		return &result
	}

	return nil
}

// String returns the password hash.
func (m *Password) String() string {
	return m.Hash
}

// Unknown returns true if the password is an empty string.
func (m *Password) Unknown() bool {
	return m.Hash == ""
}
