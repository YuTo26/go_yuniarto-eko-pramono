package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	"project/models"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) GetUsersController(c echo.Context) error {
	var users []models.User

	if err := uc.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func (uc *UserController) GetUserByIdController(c echo.Context) error {
	id := c.Param("id")

	var user models.User
	if err := uc.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func (uc *UserController) CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := uc.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func (uc *UserController) UpdateUserController(c echo.Context) error {
	id := c.Param("id")

	var user models.User
	if err := uc.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&user)

	if err := uc.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func (uc *UserController) DeleteUserController(c echo.Context) error {
	id := c.Param("id")

	var user models.User
	if err := uc.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := uc.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}
