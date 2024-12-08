package controller

import (
	"Assignment2/Web/dbhelper"
	"errors"
	"gorm.io/gorm"
)

func AllUsers() ([]Model.Users, error) {
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return nil, err
	}
	var users []Model.Users
	res := connection.Where("status = 0").Find(&users)
	if res.Error != nil {
		return nil, errors.New("error retrieving users: " + res.Error.Error())
	}
	if len(users) == 0 {
		return nil, errors.New("no users found")
	}
	return users, nil
}
func AddUser(item *Model.Users) (string, error) {
	if item.Name == "" {
		return "name", errors.New("name is required")
	}
	if item.Surname == "" {
		return "surname", errors.New("surname is required")
	}
	if item.Email == "" {
		return "email", errors.New("email is required")
	}
	if item.Password == "" {
		return "password", errors.New("password is required")
	}
	if item.Role == "" {
		return "Role", errors.New("Role is required")
	}
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return "error", err
	}
	if item.ID <= 0 {
		res := connection.Create(&Model.Users{
			Name:     item.Name,
			Surname:  item.Surname,
			Role:     item.Role,
			Age:      item.Age,
			Birthday: item.Birthday,
			Job:      item.Job,
			Email:    item.Email,
			Password: item.Password,
			Status:   item.Status,
		})
		if res.Error != nil {
			return "error", res.Error
		}

	}
	return "success", nil
}

func UpdateUser(item *Model.Users) (string, error) {

	if item.ID <= 0 {
		return "id", errors.New("id is required")
	}
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return "error", err
	}

	if item.ID >= 0 {
		var user Model.Users
		res := connection.Where("status = 0 AND id = ?", item.ID).First(&user)

		if res.Error != nil {
			return "error", errors.New("user not found")
		} else {
			user.Name = item.Name
			user.Surname = item.Surname
			user.Age = item.Age
			Model.Role = item.Role
			user.Birthday = item.Birthday
			user.Job = item.Job
			user.Email = item.Email
			user.Password = item.Password
			user.Status = item.Status
			res = connection.Save(&user)
			if res.Error != nil {
				return "error", res.Error
			}
		}
	}
	return "success", nil
}

func DeleteUser(item *Model.Users) (string, error) {
	if item.ID <= 0 {
		return "id", errors.New("id is required")
	}
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return "error", err
	}

	if item.ID >= 0 {
		var user Model.Users
		res := connection.Where("status = 0 AND id =? ", item.ID).First(&user)

		if res.Error != nil {
			return "error", errors.New("user not found")
		} else {
			user.Status = 1
			res = connection.Save(&user)
			if res.Error != nil {
				return "error", res.Error
			}
		}
	}
	return "success", nil
}

func Login(item *Model.Users) (string, error) {
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return "", err
	}

	if item.Email == "" {
		return "", errors.New("email is required")
	}
	if item.Password == "" {
		return "", errors.New("password is required")
	}

	var foundUser Model.Users
	if err := connection.Where("status = ? AND email = ? AND password = ?", 0, item.Email, item.Password).First(&foundUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	return "Success", nil
}
