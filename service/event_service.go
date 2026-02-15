package service

import (
	"base/models"
	"base/repository"
)

func CreateEvent(event *models.Event) error {
	return repository.CreateEvent(event)
}

func GetEvents() ([]models.Event, error) {
	return repository.GetAllEvents()
}

func GetEventByID(id uint) (*models.Event, error) {
	return repository.GetEventByID(id)
}