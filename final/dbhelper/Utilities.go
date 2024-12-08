package dbhelper

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"strings"
	"time"
)

var DB *gorm.DB

type CustomNamingStrategy struct {
	schema.NamingStrategy
}

func (c CustomNamingStrategy) TableName(model string) string {
	return strings.ToLower(c.NamingStrategy.TableName(model))
}

func GetOpenConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "go_db"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "go_user"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "go_password"
	}
	fmt.Printf("DB_HOST=%s, DB_PORT=%s, DB_NAME=%s, DB_USER=%s, DB_PASSWORD=%s\n", host, port, dbName, user, password)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: CustomNamingStrategy{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func CloseConnection(db *gorm.DB) error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
