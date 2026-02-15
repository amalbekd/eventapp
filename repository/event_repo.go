package repository

import (
	"base/models"
)

func CreateEvent(event *models.Event) error {
	return DB.Create(event).Error
}

func GetAllEvents() ([]models.Event, error) {
	var events []models.Event
	err := DB.Find(&events).Error
	return events, err
}

func GetEventByID(id uint) (*models.Event, error) {
	var event models.Event
	err := DB.First(&event, id).Error
	return &event, err
}

func UpdateEvent(event *models.Event) error {
	return DB.Save(event).Error
}

func DeleteEvent(id uint) error {
	return DB.Delete(&models.Event{}, id).Error
}