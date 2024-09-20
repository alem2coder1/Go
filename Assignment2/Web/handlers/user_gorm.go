package handlers

import (
	"Assignment2/Model/user"

	"gorm.io/gorm"
)

// 全局数据库连接实例
var dbGorm *gorm.DB

// 初始化 GORM 数据库连接
func InitDBGorm(gormDB *gorm.DB) {
	dbGorm = gormDB
}

// 使用 GORM 获取用户列表
func GetUsersGORM() ([]user.Users, error) {
	var users []user.Users
	result := dbGorm.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
