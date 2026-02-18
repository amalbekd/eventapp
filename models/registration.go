package models

import (
	"time"
	"gorm.io/gorm"
)

type Registration struct {
	ID 			uint `gorm:"primaryKey" json:"id"`
	UserID 		uint `gorm:"not null" json:"user_id"`
	User 		User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	EventID 	uint `gorm:"not null" json:"event_id"`
	Event 		Event `gorm:"foreignKey:EventID" json:"event,omitempty"`
	Status 		string `gorm:"default:'pending'" json:"status"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
	DeleteddAt 	gorm.DeletedAt `gorm:"index" json:"-"`
}