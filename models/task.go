package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

// Task "Object
type Task struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Completed bool      `json:"completed"`
}

func (task *Task) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	u, _ := uuid.NewV4()
	scope.SetColumn("ID", u.String())
	return nil
}

func (task *Task) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
