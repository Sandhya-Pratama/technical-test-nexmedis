package service

import (
	"errors"

	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

// HashPassword hashes a plain text password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// VerifyPassword compares a hashed password with a plain text password
func verifyPassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func (s *UserService) RegisterUser(user *models.User) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("username, email, and password are required")
	}

	existingUser, err := s.UserRepo.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("username already exists")
	}

	// Hash the password before storing
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	return s.UserRepo.CreateUser(user)
}

func (s *UserService) LoginUser(username, password string) (*models.User, error) {
	user, err := s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Verify the password
	if err := verifyPassword(user.Password, password); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
