package service

import (
	"base/models"
	"base/repository"

	"time"
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(userID uint, role string) (string, error) { // добавили role
	secret := strings.TrimSpace(os.Getenv("JWT_SECRET"))
	
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role, // теперь роль настоящая
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func RegisterUser(input models.RegisterInput) (*models.User, error) {
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email:    input.Email,
		Password: hashedPassword,
		FullName: input.FullName,
		Role:	 "student",
	}

	if err := repository.CreateUser(&user); err != nil {
		return nil, errors.New("User already exists")
	}

	return &user, nil
}

func LoginUser(input models.LoginInput) (string, error) {
	user, err := repository.GetUserByEmail(input.Email)
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	if !CheckPassword(input.Password, user.Password) {
		return "", errors.New("Invalid credentials")
	}

	return generateToken(user.ID, user.Role)
}