package service

import (
	"base/models"
	"base/repository"
	"errors"
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

func UpdateEvent(userID uint, eventID uint, input models.Event) (*models.Event, error) {
	event, err := repository.GetEventByID(eventID)
	if err != nil {
		return nil, errors.New("event not found!!")
	}

	if event.OrganizerID != userID {
		return nil, errors.New("you are not the organizer of this event")
	}
	event.Title = input.Title
	event.Description = input.Description
	event.Category = input.Category
	event.City = input.City
	event.Format = input.Format
	event.Date = input.Date

	if err := repository.UpdateEvent(event); err != nil {
		return nil, err
	}

	return event, nil
}

func DeleteEvent(userID uint, eventID uint) error {
	event, err := repository.GetEventByID(eventID)
	if err != nil {
		return errors.New("event not found!!")
	}

	if event.OrganizerID != userID {
		return errors.New("you are not the organizer of this event")
	}

	return repository.DeleteEvent(eventID)
}