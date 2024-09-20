package services

import (
	"Assignment2/Model/user"

	"gorm.io/gorm"
)

var dbGorm *gorm.DB

func InitDBGorm(gormDB *gorm.DB) {
	dbGorm = gormDB
}

func GetUsersGORM() ([]user.Users, error) {
	var users []user.Users
	result := dbGorm.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
