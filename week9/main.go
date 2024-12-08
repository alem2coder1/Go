package week9

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"time"
	"week9/dbhelper"
)

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

func Login(item *Model.Users) (*LoginResponse, error) {
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return nil, err
	}
	if item.Email == "" {
		return nil, errors.New("email is required")
	}
	if item.Password == "" {
		return nil, errors.New("password is required")
	}

	var foundUser Model.Users
	if err := connection.Where("status = ? AND email = ? AND password = ?", 0, item.Email, item.Password).First(&foundUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	var userRole string
	if err := connection.Model(&Model.Users{}).Where("email = ?", item.Email).Select("role").Scan(&userRole).Error; err != nil {
		return nil, err
	}

	if userRole == "" {
		userRole = "user"
	}
	tokenString, err := generateJWT(foundUser.Email, userRole)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token:   tokenString,
		Message: "Success",
	}, nil
}

var jwtKey = []byte("gofristproject-alem")

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func generateJWT(email, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
