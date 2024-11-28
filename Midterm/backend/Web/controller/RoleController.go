package controller

import (
	"backend/Model/user"
	"backend/Web/dbhelper"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func AllRole() ([]user.Role, error) {
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return nil, err
	}
	var role []user.Role
	res := connection.Where("status = 0").Find(&role)
	if res.Error != nil {
		return nil, errors.New("error retrieving users: " + res.Error.Error())
	}
	if len(role) == 0 {
		return nil, errors.New("no users found")
	}
	return role, nil
}

var jwtKey = []byte("gofristproject-alem")

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func generateJWT(email, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 设置过期时间为24小时后

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
