package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Format 		string `json:"format"`
	City 	  	string `json:"city"`
	Date 	  	time.Time `json:"date"`
	OrganizerID uint   `json:"organizer_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}