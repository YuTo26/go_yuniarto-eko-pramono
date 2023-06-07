package database

import (
	"errors"
	"project/config"
	"project/models"
)

var ErrInvalidID = errors.New("invalid ID")

func GetBlogs() ([]models.Blogs, error) {
	var blogs []models.Blogs

	err := config.DB.Find(&blogs).Error
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func GetBlogsByID(id int) (models.Blogs, error) {
	var blogs models.Blogs

	err := config.DB.First(&blogs, id).Error
	if err != nil {
		return blogs, err
	}

	return blogs, nil
}

func CreateBlogs(blogs models.Blogs) (models.Blogs, error) {
	err := config.DB.Create(&blogs).Error

	if err != nil {
		return blogs, err
	}

	return blogs, nil
}

func DeleteBlogs(id int) error {
	var blogs models.Blogs

	result := config.DB.Where("id = ?", id).Delete(&blogs)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrInvalidID
	}

	return nil
}

func UpdateBlogs(id int, blogs models.Blogs) (models.Blogs, error) {
	result := config.DB.Model(&models.Blogs{}).Where("id = ?", id).Updates(blogs)
	if result.Error != nil {
		return blogs, result.Error
	}

	if result.RowsAffected == 0 {
		return blogs, ErrInvalidID
	}

	return blogs, nil
}
