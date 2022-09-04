package services

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/foxieze/tsundoku-server/config"
	"github.com/foxieze/tsundoku-server/entities"
	"github.com/golang-jwt/jwt/v4"
)

func createToken(userId int, username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userId,
		"username": username,
	})

	token, err := claims.SignedString("secret")

	if err != nil {
		return "", err
	}

	return token, nil
}

func LoginUser(username string, password string) (string, error) {
	// Check if user exists
	var user entities.User
	result := config.Database.Where("username = ?", username).First(&user)
	if result.RowsAffected == 0 {
		return "", errors.New("user not found")
	}

	// Check if password is correct
	passwordCorrect, err := checkPassword(user, password)
	if (!passwordCorrect) || (err != nil) {
		return "", errors.New("invalid password")
	}

	// Create token
	token, err := createToken(user.Id, username)
	if err != nil {
		return "", err
	}

	// Return token
	return token, nil
}

func hashPassword(password string) (string, error) {
	// convert password to sha256
	passwordconverted := sha256.Sum256([]byte(password))

	// bcrypt hash password
	hash, err := bcrypt.GenerateFromPassword(passwordconverted[:], 14)
	if err != nil {
		return "", err
	}

	// convert hash to base64
	hashconverted := base64.StdEncoding.EncodeToString(hash)

	// return hash
	return hashconverted, nil
}

func checkPassword(user entities.User, password string) (bool, error) {
	// convert base64 to hash
	hash, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil {
		return false, err
	}

	// convert password to sha256
	passwordconverted := sha256.Sum256([]byte(password))

	// compare password
	err = bcrypt.CompareHashAndPassword(hash, passwordconverted[:])
	if err != nil {
		return false, err
	}

	return true, nil
}
