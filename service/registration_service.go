package service

import (
	"base/models"
	"base/repository"
	"errors"
)

func RegisterToEvent(userID uint, eventID uint) (*models.Registration, error) {
	existing, _ := repository.GetRegistration(userID, eventID)
	if existing.ID != 0 {
		return nil, errors.New("you are already registered for this event")
	}

	reg := models.Registration {
		UserID: userID,
		EventID: eventID,
		Status: "pending",
	}

	if err := repository.CreateRegistration(&reg); err != nil {
		return nil, err
	}

	return &reg, nil
}


func GetMyEvents(userID uint) ([]models.Registration, error) {
	return repository.GetUserRegistrations(userID)
}