package controller

import (
	"backend/Web/dbhelper"
	"errors"
	"final/Model"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func AllRole(offset, limit int) ([]Model.Role, error) {
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return nil, err
	}
	defer func() {
		err := dbhelper.CloseConnection()
		if err != nil {

		}
	}()

	var roles []Model.Role
	query := connection.Where("status = 0")
	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}

	res := query.Find(&roles)
	if res.Error != nil {
		return nil, fmt.Errorf("error retrieving roles: %w", res.Error)
	}

	if len(roles) == 0 {
		return nil, errors.New("no roles found")
	}
	return roles, nil
}

var jwtKey = []byte(getJWTKey())

func getJWTKey() string {
	if key := os.Getenv("JWT_SECRET"); key != "" {
		return key
	}
	return "default-secret-key"
}

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func generateJWT(email, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := Claims{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenString, nil
}
