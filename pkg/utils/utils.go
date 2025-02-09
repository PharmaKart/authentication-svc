package utils

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func ValidateUserInput(username, email, password, firstName, lastName, phone, dob, billing1, city, province, postalCode, country string) error {
	if strings.TrimSpace(username) == "" {
		return errors.New("username is required")
	}

	if strings.TrimSpace(email) == "" {
		return errors.New("email is required")
	}
	if !isValidEmail(email) {
		return errors.New("invalid email format")
	}

	if strings.TrimSpace(password) == "" {
		return errors.New("password is required")
	}
	if !isValidPassword(password) {
		return errors.New("password must be at least 8 characters long and contain an uppercase letter, a lowercase letter, a digit, and a special character")
	}

	if strings.TrimSpace(firstName) == "" {
		return errors.New("first name is required")
	}

	if strings.TrimSpace(lastName) == "" {
		return errors.New("last name is required")
	}

	if strings.TrimSpace(phone) == "" {
		return errors.New("phone number is required")
	}
	if !isValidPhone(phone) {
		return errors.New("invalid phone number format")
	}

	if strings.TrimSpace(dob) == "" {
		return errors.New("date of birth is required")
	}
	if !isValidDOB(dob) {
		return errors.New("invalid date of birth or must be in the past")
	}

	if strings.TrimSpace(billing1) == "" {
		return errors.New("billing address line 1 is required")
	}

	if strings.TrimSpace(city) == "" {
		return errors.New("city is required")
	}

	if strings.TrimSpace(province) == "" {
		return errors.New("province is required")
	}

	if strings.TrimSpace(postalCode) == "" {
		return errors.New("postal code is required")
	}

	if strings.TrimSpace(country) == "" {
		return errors.New("country is required")
	}

	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	re := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	return re.MatchString(password)
}

func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(phone)
}

func isValidDOB(dob string) bool {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return false
	}
	return parsedDOB.Before(time.Now())
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ParseDOB(dob string) (time.Time, error) {
	return time.Parse("2006-01-02", dob)
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateJWT(userID, role, secret string) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["userid"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with the secret
	return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenString, secret string) (string, string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", "", err
	}

	// Check if the token is valid
	if !token.Valid {
		return "", "", errors.New("invalid token")
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("invalid token")
	}

	userID, ok := claims["userid"].(string)
	if !ok {
		return "", "", errors.New("invalid token")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", "", errors.New("invalid token")
	}

	return userID, role, nil
}
