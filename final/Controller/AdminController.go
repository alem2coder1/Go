package controller

import (
	"errors"
	"final/Common"
	"final/Model"
	"final/dbhelper"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
		err := dbhelper.CloseConnection(connection)
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
		err := dbhelper.CloseConnection(connection)
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
	if err := Common.ValidateUserInput(item); err != nil {
		return "error", err
	}

	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return "error", err
	}
	defer func() {
		err := dbhelper.CloseConnection(connection)
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
		err := dbhelper.CloseConnection(connection)
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
	existingUser.Role = item.Role
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
		err := dbhelper.CloseConnection(connection)
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

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // Redis 地址
})

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

const maxLoginAttempts = 5        // 最大登录尝试次数
const lockoutDuration = time.Hour // 锁定时长（1小时）

func Login(c *gin.Context, item *Model.Users) (*LoginResponse, error) {
	startTime := time.Now()

	log.Printf("Login attempt for email: %s", item.Email)

	if item.Email == "" || item.Password == "" {
		return nil, errors.New("email and password are required")
	}

	cacheKey := "login_attempts:" + item.Email

	attempts, err := rdb.Get(context.Background(), cacheKey).Int()
	if err == redis.Nil {
		attempts = 0
	} else if err != nil {
		log.Printf("Redis error: %v", err)
		return nil, err
	}

	if attempts >= maxLoginAttempts {
		log.Printf("Too many login attempts for email: %s", item.Email)
		return nil, errors.New("too many login attempts, please try again later")
	}

	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		log.Printf("Error opening DB connection: %v", err)
		return nil, err
	}
	defer func() {
		if err := dbhelper.CloseConnection(connection); err != nil {
			log.Printf("Error closing DB connection: %v", err)
		}
	}()

	var foundUser Model.Users
	res := connection.Where("status = ? AND email = ?", 0, item.Email).First(&foundUser)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			rdb.Incr(context.Background(), cacheKey)
			rdb.Expire(context.Background(), cacheKey, lockoutDuration)
			log.Printf("Invalid email or password: %s", item.Email)
			return nil, errors.New("invalid email or password")
		}
		log.Printf("Database query error: %v", res.Error)
		return nil, res.Error
	}

	if err := verifyPassword(foundUser.Password, item.Password); err != nil {
		rdb.Incr(context.Background(), cacheKey)
		rdb.Expire(context.Background(), cacheKey, lockoutDuration)
		log.Printf("Invalid password attempt for email: %s", item.Email)
		return nil, errors.New("invalid email or password")
	}

	rdb.Del(context.Background(), cacheKey)

	tokenString, err := generateJWT(foundUser.Email, foundUser.Role)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	c.SetCookie("token", tokenString, int(24*time.Hour.Seconds()), "/", "localhost", false, true)

	log.Printf("Login successful for email: %s", item.Email)

	elapsedTime := time.Since(startTime)
	log.Printf("Login process completed in %s", elapsedTime)

	return &LoginResponse{
		Token:   tokenString,
		Message: "Success",
	}, nil
}
