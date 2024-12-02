package dao

import (
	"errors"
	configs "user-crud/config"
	"user-crud/models"

	"gorm.io/gorm"
)

func CreateUser(user models.User) error {
	return configs.InitDB().Create(&user).Error
}

func FetchAllUsers() ([]models.User, error) {
	var users []models.User
	err := configs.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id string) (models.User, error) {
	var user models.User
	err := configs.InitDB().First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return user, errors.New("user not found")
	}
	return user, err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := configs.InitDB().Where("email = ?", email).First(&user).Error
	return user, err
}

func UpdateUser(user models.User) error {
	return configs.InitDB().Save(&user).Error
}

func DeleteUser(id string) error {
	return configs.InitDB().Delete(&models.User{}, id).Error
}
