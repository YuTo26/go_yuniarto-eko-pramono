package database

import (
	"project/config"
	"project/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(id int) (models.User, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user *models.User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(user *models.User) error {
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
