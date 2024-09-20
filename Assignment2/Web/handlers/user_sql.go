package handlers

import (
	"Assignment2/Model/user"
	"database/sql"
)

var db *sql.DB

func GetUsersSQL() ([]user.Users, error) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.Users
	for rows.Next() {
		var u user.Users
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
