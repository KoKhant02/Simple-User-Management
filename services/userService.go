package services

import (
	"errors"
	dao "user-crud/daos"
	"user-crud/models"
	"user-crud/utils"
)

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(input UserInput) error {
	hashedPassword, _ := utils.HashPassword(input.Password)
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     input.Role,
	}
	return dao.CreateUser(user)
}

func LoginUser(input LoginInput) (string, error) {
	user, err := dao.GetUserByEmail(input.Email)
	if err != nil || !utils.CheckPasswordHash(input.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}
	return utils.GenerateJWT(user.ID, user.Role), nil
}

func GetUserList() ([]models.User, error) {
	users, err := dao.FetchAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id string) (models.User, error) {
	return dao.GetUserByID(id)
}

func UpdateUser(id string, input UserInput) error {
	user, err := dao.GetUserByID(id)
	if err != nil {
		return err
	}
	user.Name = input.Name
	return dao.UpdateUser(user)
}

func DeleteUser(id string) error {
	return dao.DeleteUser(id)
}
