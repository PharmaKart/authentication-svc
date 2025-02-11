package services

import (
	"errors"

	"github.com/PharmaKart/authentication-svc/internal/models"
	"github.com/PharmaKart/authentication-svc/internal/repositories"
	"github.com/PharmaKart/authentication-svc/pkg/utils"
)

type AuthService interface {
	Register(username, email, password, firstName, lastName, phone, dob, streetLine1, streetLine2, city, province, postalCode, country string) error
	Login(email, username, password string) (string, string, string, error)
	VerifyToken(token string) (string, string, error)
}

type authService struct {
	userRepo     repositories.UserRepository
	customerRepo repositories.CustomerRepository
	jwtSecret    string
}

func NewAuthService(userRepo repositories.UserRepository, customerRepo repositories.CustomerRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:     userRepo,
		customerRepo: customerRepo,
		jwtSecret:    jwtSecret,
	}
}

func (s *authService) Register(username, email, password, firstName, lastName, phone, dob, streetLine1, streetLine2, city, province, postalCode, country string) error {
	// Check if the user already exists
	_, err := s.userRepo.GetUserByEmail(email)
	if err == nil {
		return errors.New("user already exists")
	}

	_, err = s.userRepo.GetUserByUserName(email)
	if err == nil {
		return errors.New("user already exists")
	}

	// Validate the user input
	if err := utils.ValidateUserInput(username, email, password, firstName, lastName, phone, dob, streetLine1, city, province, postalCode, country); err != nil {
		return err
	}

	// Hash the password
	passwordHash, err := utils.HashPassword(password)

	if err != nil {
		return err
	}

	// Add the user to the database
	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         "customer", // Only customers can register
	}

	userID, err := s.userRepo.CreateUser(user)

	if err != nil {
		return err
	}

	// Add the customer to the database

	// Parse the date of birth
	dobTime, err := utils.ParseDOB(dob)
	if err != nil {
		return err
	}

	customer := &models.Customer{
		UserID:      userID,
		FirstName:   firstName,
		LastName:    lastName,
		Phone:       &phone,
		DateOfBirth: &dobTime,
		StreetLine1: streetLine1,
		StreetLine2: &streetLine2,
		City:        city,
		Province:    province,
		PostalCode:  postalCode,
		Country:     country,
	}

	_, err = s.customerRepo.CreateCustomer(customer)

	if err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(email, username, password string) (string, string, string, error) {
	// Get the user from the database
	var user *models.User
	var err error
	if username != "" {
		user, err = s.userRepo.GetUserByUserName(username)
		if err != nil {
			return "", "", "", errors.New("user not found")
		}
	} else {
		user, err = s.userRepo.GetUserByEmail(email)
		if err != nil {
			return "", "", "", errors.New("user not found")
		}
	}

	// Check if the password is correct
	err = utils.CheckPasswordHash(password, user.PasswordHash)
	if err != nil {
		return "", "", "", errors.New("incorrect password")
	}

	// Generate a JWT
	token, err := utils.GenerateJWT(user.ID.String(), user.Role, s.jwtSecret)
	if err != nil {
		return "", "", "", err
	}

	// Log
	utils.Info("User logged in", map[string]interface{}{
		"userID":   user.ID.String(),
		"username": user.Username,
		"role":     user.Role,
	})

	return token, user.ID.String(), user.Role, nil
}

func (s authService) VerifyToken(token string) (string, string, error) {
	userID, role, err := utils.ValidateJWT(token, s.jwtSecret)
	if err != nil {
		return "", "", err
	}

	return userID, role, nil
}
