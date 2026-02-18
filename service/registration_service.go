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

func GetParticipants(userID uint, eventID uint) ([]models.Registration, error) {
	event, err := repository.GetEventByID(eventID)
	if err != nil || event.OrganizerID != userID {
		return nil, errors.New("access denied: you are not the organizer")
	}

	return repository.GetRegistrationsByEvent(eventID)
}


func UpdateApplicationStatus(userID uint, regID uint, newStatus string) (*models.Registration, error) {
	reg, err := repository.GetRegistrationByID(regID)
	if err != nil {
		return nil, errors.New("registration not found")
	}

	if reg.Event.OrganizerID != userID {
		return nil, errors.New("access denied: you are not the organizer")
	}

	reg.Status = newStatus
	if err := repository.UpdateRegistration(reg); err != nil {
		return nil, err
	}

	return reg, nil
}