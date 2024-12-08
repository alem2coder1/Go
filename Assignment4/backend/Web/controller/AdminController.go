package controller

import (
	"backend/Web/dbhelper"
	"errors"
	"final/Model"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func FindUserByID(connection *gorm.DB, id int) (*Model.Users, error) {
	var user Model.Users
	res := connection.Where("status = 0 AND id = ?", id).First(&user)
	if res.Error != nil {
		return nil, fmt.Errorf("user with ID %d not found: %w", id, res.Error)
	}
	return &user, nil
}
func GetUser(id int) (*Model.Users, error) {
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return nil, err
	}
	defer func() {
		err := dbhelper.CloseConnection()
		if err != nil {

		}
	}()

	var retrievedUser Model.Users
	res := connection.Where("status = 0 AND id = ?", id).First(&retrievedUser)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error retrieving user: %w", res.Error)
	}

	return &retrievedUser, nil
}

func AllUsers(offset, limit int) ([]Model.Users, error) {
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return nil, err
	}
	defer func() {
		err := dbhelper.CloseConnection()
		if err != nil {

		}
	}()

	var users []Model.Users
	query := connection.Where("status = 0")
	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}
	res := query.Find(&users)
	if res.Error != nil {
		return nil, fmt.Errorf("error retrieving users: %w", res.Error)
	}

	return users, nil
}

func AddUser(item *Model.Users) (string, error) {
	if item.Name == "" || item.Surname == "" || item.Email == "" || item.Password == "" || item.Role == "" {
		return "error", errors.New("all fields (name, surname, email, password, role) are required")
	}

	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return "error", err
	}
	defer func() {
		err := dbhelper.CloseConnection()
		if err != nil {

		}
	}()

	hashedPassword, err := hashPassword(item.Password)
	if err != nil {
		return "error", fmt.Errorf("error hashing password: %w", err)
	}

	res := connection.Create(&Model.Users{
		Name:     item.Name,
		Surname:  item.Surname,
		Role:     item.Role,
		Age:      item.Age,
		Birthday: item.Birthday,
		Job:      item.Job,
		Email:    item.Email,
		Password: hashedPassword,
		Status:   item.Status,
	})
	if res.Error != nil {
		return "error", res.Error
	}
	return "success", nil
}

func UpdateUser(item *Model.Users) (string, error) {
	if item.ID <= 0 {
		return "error", errors.New("id is required")
	}

	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return "error", err
	}
	defer func() {
		err := dbhelper.CloseConnection()
		if err != nil {

		}
	}()

	var existingUser Model.Users
	res := connection.Where("status = 0 AND id = ?", item.ID).First(&existingUser)
	if res.Error != nil {
		return "error", fmt.Errorf("user with ID %d not found: %w", item.ID, res.Error)
	}

	existingUser.Name = item.Name
	existingUser.Surname = item.Surname
	existingUser.Age = item.Age
	existingModel.Role = item.Role
	existingUser.Birthday = item.Birthday
	existingUser.Job = item.Job
	existingUser.Email = item.Email
	existingUser.Status = item.Status

	res = connection.Save(&existingUser)
	if res.Error != nil {
		return "error", res.Error
	}
	return "success", nil
}

func DeleteUser(item *Model.Users) (string, error) {
	if item.ID <= 0 {
		return "error", errors.New("id is required")
	}

	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return "error", err
	}
	defer func() {
		err := dbhelper.CloseConnection()
		if err != nil {

		}
	}()

	var existingUser Model.Users
	res := connection.Where("status = 0 AND id = ?", item.ID).First(&existingUser)
	if res.Error != nil {
		return "error", fmt.Errorf("user with ID %d not found: %w", item.ID, res.Error)
	}

	existingUser.Status = 1
	res = connection.Save(&existingUser)
	if res.Error != nil {
		return "error", res.Error
	}
	return "success", nil
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

func Login(c *gin.Context, item *Model.Users) (*LoginResponse, error) {
	if item.Email == "" || item.Password == "" {
		return nil, errors.New("email and password are required")
	}

	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return nil, err
	}
	defer func() {
		err := dbhelper.CloseConnection()
		if err != nil {

		}
	}()

	var foundUser Model.Users
	res := connection.Where("status = ? AND email = ?", 0, item.Email).First(&foundUser)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, res.Error
	}

	if err := verifyPassword(foundUser.Password, item.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	tokenString, err := generateJWT(foundUser.Email, foundModel.Role)
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	c.SetCookie("token", tokenString, int(24*time.Hour.Seconds()), "/", "localhost", false, true)

	return &LoginResponse{
		Token:   tokenString,
		Message: "Success",
	}, nil
}
