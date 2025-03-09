package utils

import (
	"errors"
	"time"

	"github.com/PharmaKart/authentication-svc/internal/proto"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
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

func ConvertMapToKeyValuePairs(m map[string]string) []*proto.KeyValuePair {
	if m == nil {
		return nil
	}

	result := make([]*proto.KeyValuePair, 0, len(m))
	for k, v := range m {
		result = append(result, &proto.KeyValuePair{
			Key:   k,
			Value: v,
		})
	}
	return result
}
