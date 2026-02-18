package repository

import "base/models"

func CreateRegistration(reg *models.Registration) error {
	return DB.Create(reg).Error
}

func GetRegistration(userID, eventID uint) (*models.Registration, error) {
	var reg models.Registration
	err := DB.Where("user_id = ? AND event_id = ?", userID, eventID).First(&reg).Error
	return &reg, err
}

func GetUserRegistrations(userID uint) ([]models.Registration, error) {
	var regs []models.Registration
	err := DB.Preload("Event").Where("user_id = ?", userID).Find(&regs).Error
	return regs, err
}

func GetRegistrationsByEvent(eventID uint) ([]models.Registration, error) {
	var regs []models.Registration
	err := DB.Preload("User").Where("event_id = ?", eventID).Find(&regs).Error
	return regs, err
}

func GetRegistrationByID(regID uint) (*models.Registration, error) {
	var reg models.Registration
	err := DB.Preload("Event").First(&reg, regID).Error
	return &reg, err
}

func UpdateRegistration(reg *models.Registration) error {
	return DB.Save(reg).Error
}