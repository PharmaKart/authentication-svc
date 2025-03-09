package utils

import (
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/PharmaKart/authentication-svc/pkg/errors"
)

func ValidateUserInput(username, email, password, firstName, lastName, phone, dob, billing1, city, province, postalCode, country string) error {
	validationErrors := make(map[string]string)

	if strings.TrimSpace(username) == "" {
		validationErrors["username"] = "Username is required"
	}

	if strings.TrimSpace(email) == "" {
		validationErrors["email"] = "Email is required"
	} else if !isValidEmail(email) {
		validationErrors["email"] = "Invalid email format"
	}

	if strings.TrimSpace(password) == "" {
		validationErrors["password"] = "Password is required"
	} else if !isValidPassword(password) {
		validationErrors["password"] = "Password must be at least 8 characters long and contain an uppercase letter, a lowercase letter, a digit, and a special character"
	}

	if strings.TrimSpace(firstName) == "" {
		validationErrors["firstName"] = "First name is required"
	}

	if strings.TrimSpace(lastName) == "" {
		validationErrors["lastName"] = "Last name is required"
	}

	if strings.TrimSpace(phone) == "" {
		validationErrors["phone"] = "Phone number is required"
	} else if !isValidPhone(phone) {
		validationErrors["phone"] = "Invalid phone number format"
	}

	if strings.TrimSpace(dob) == "" {
		validationErrors["dateOfBirth"] = "Date of birth is required"
	} else if !isValidDOB(dob) {
		validationErrors["dateOfBirth"] = "Invalid date of birth or must be in the past"
	}

	if strings.TrimSpace(billing1) == "" {
		validationErrors["streetLine1"] = "Billing address line 1 is required"
	}

	if strings.TrimSpace(city) == "" {
		validationErrors["city"] = "City is required"
	}

	if strings.TrimSpace(province) == "" {
		validationErrors["province"] = "Province is required"
	}

	if strings.TrimSpace(postalCode) == "" {
		validationErrors["postalCode"] = "Postal code is required"
	} else if !regexp.MustCompile(`^[A-Z][0-9][A-Z] [0-9][A-Z][0-9]$`).MatchString(postalCode) {
		validationErrors["postalCode"] = "Invalid postal code format"
	}

	if strings.TrimSpace(country) == "" {
		validationErrors["country"] = "Country is required"
	}

	if len(validationErrors) > 0 {
		return errors.NewValidationErrors(validationErrors)
	}

	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	specialChars := regexp.MustCompile(`[@$!%*?&]`)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case specialChars.MatchString(string(char)):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}

func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^\+1\s\(\d{3}\)\s\d{3}\-\d{4}$`)
	return re.MatchString(phone)
}

func isValidDOB(dob string) bool {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return false
	}
	return parsedDOB.Before(time.Now())
}

func ParseDOB(dob string) (time.Time, error) {
	return time.Parse("2006-01-02", dob)
}
