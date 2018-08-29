package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

// AppPortal "Object
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	Group     string    `json:"group"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	u, _ := uuid.NewV4()
	scope.SetColumn("ID", u.String())
	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
