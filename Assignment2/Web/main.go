package main

import (
	"Assignment2/Web/services"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 设置 MySQL 连接字符串
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接到 MySQL 数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// 初始化 GORM 数据库连接
	services.InitDBGorm(db)

	// 获取用户数据
	users, err := services.GetUsersGORM()
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return
	}

	fmt.Println("Users:", users)
}
