package Common

import (
	"errors"
	"final/Model"
	"regexp"
)

func isTrueEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
func isValidPassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	re := regexp.MustCompile(`^(?=.*[a-zA-Z])(?=.*\d)[a-zA-Z\d]+$`)
	if !re.MatchString(password) {
		return errors.New("password must contain both letters and numbers")
	}
	return nil
}

func ValidateUserInput(item *Model.Users) error {
	if item.Name == "" || item.Surname == "" || item.Email == "" || item.Password == "" || item.Role == "" {
		return errors.New("all fields (name, surname, email, password, role) are required")
	}

	if !isTrueEmail(item.Email) {
		return errors.New("invalid email format")
	}

	if err := isValidPassword(item.Password); err != nil {
		return err
	}

	return nil
}
