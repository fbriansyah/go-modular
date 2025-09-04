package userModel

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/fbriansyah/go-modular/utils"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(id, email, password, firstName, lastName string) (*User, error) {
	user := &User{
		UUID:      id,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Status:    UserStatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}
	return user, nil
}

type User struct {
	UUID      string     `json:"id" db:"id"`
	FirstName string     `json:"first_name" db:"first_name"`
	LastName  string     `json:"last_name" db:"last_name"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password" db:"password"`
	Status    UserStatus `json:"status" db:"status"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	Version   int        `json:"version" db:"version"` // Optimistic locking
}

// CheckPassword verifies if the provided password matches the user's password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) SetPassword(password string) error {
	if err := utils.ValidatePassword(password); err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	u.Password = string(hashedPassword)
	return nil
}

// Validate performs comprehensive validation of the User aggregate
func (u *User) Validate() error {
	if u.UUID == "" {
		return errors.New("user ID cannot be empty")
	}

	if err := u.validateEmail(); err != nil {
		return err
	}

	if err := u.validateName(); err != nil {
		return err
	}

	return nil
}

// validateEmail validates the email format and requirements
func (u *User) validateEmail() error {
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}

	// Basic email regex pattern
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}

	if len(u.Email) > 255 {
		return errors.New("email cannot exceed 255 characters")
	}

	return nil
}

// validateName validates first and last name requirements
func (u *User) validateName() error {
	if u.FirstName == "" {
		return errors.New("first name cannot be empty")
	}

	if u.LastName == "" {
		return errors.New("last name cannot be empty")
	}

	if len(u.FirstName) > 100 {
		return errors.New("first name cannot exceed 100 characters")
	}

	if len(u.LastName) > 100 {
		return errors.New("last name cannot exceed 100 characters")
	}

	// Check for valid characters (letters, spaces, hyphens, apostrophes)
	nameRegex := regexp.MustCompile(`^[a-zA-Z\s\-']+$`)
	if !nameRegex.MatchString(u.FirstName) {
		return errors.New("first name contains invalid characters")
	}

	if !nameRegex.MatchString(u.LastName) {
		return errors.New("last name contains invalid characters")
	}

	return nil
}
