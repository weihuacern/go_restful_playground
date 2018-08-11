package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

// AppPortal "Object
type AppPortal struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IP        string    `json:"ip"`
	AppType   string    `json:"app_type"`
	AppName   string    `json:"app_name"`
	TimeStamp time.Time `json:"timestamp"`
	URI       string    `json:"uri"`
}

func (appportal *AppPortal) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	u, _ := uuid.NewV4()
	scope.SetColumn("ID", u.String())
	return nil
}

func (appportal *AppPortal) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
