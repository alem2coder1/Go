package services

import (
	"Assignment2/Model/user" // 导入 Model 中的 User 结构体
	"database/sql"
)

// 假设你已经初始化了数据库连接 db
var db *sql.DB

func GetUsersSQL() ([]user.Users, error) {
	// 执行 SQL 查询
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 定义一个用户切片来存储结果
	var users []user.Users

	// 遍历查询结果并将其添加到切片中
	for rows.Next() {
		var u user.Users
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	// 检查是否在循环中发生错误
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// 返回结果集
	return users, nil
}
