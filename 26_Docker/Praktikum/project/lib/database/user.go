package database

import (
	"errors"
	"project_1/config"
	"project_1/models"
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

func AuthenticateUser(email, password string) (models.User, error) {
	var user models.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	if user.Password != password {
		return user, errors.New("invalid password")
	}

	return user, nil
}

func LogoutUser(userId int) error {
	// Perform logout logic here
	// For example, you can invalidate the user's session or access token

	return nil
}
