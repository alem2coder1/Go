package dbhelper

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbInstance *gorm.DB

func GetOpenConnection() (*gorm.DB, error) {
	host := "localhost"
	port := "3306"
	dbName := "go_db"
	user := "go_dba"
	password := "87654321"
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
func CloseConnection() error {
	if dbInstance != nil {
		sqlDB, err := dbInstance.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
