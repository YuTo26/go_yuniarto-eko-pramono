package controllers

import (
	"net/http"

	"project_1/lib/database"
	"project_1/lib/utils"
	"project_1/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

	if err := uc.DB.Create(&user).Error; err != nil {
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
		"status": "success delete users",
	})
}

func (uc *UserController) LoginController(c echo.Context) error {
	loginData := models.User{}
	c.Bind(&loginData)

	// Perform authentication logic here
	user, err := database.AuthenticateUser(loginData.Email, loginData.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	// Generate token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"token":  token,
	})
}

func (uc *UserController) LogoutController(c echo.Context) error {
	// Perform logout logic here

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Logged out successfully",
	})
}

//func (uc *UserController) authenticateUser(username, password string) (*models.User, error) {
//var user models.User

// Perform authentication logic using username and password
// Example:
// if err := uc.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
// 	return nil, err
// }

//	return &user, nil
//}

//func (uc *UserController) generateToken(user *models.User) (string, error) {
// Generate token logic here
// Example:
// claims := jwt.MapClaims{
// 	"authorized": true,
// 	"user_id":    user.ID,
// 	"exp":        time.Now().Add(time.Hour * 24).Unix(),
// }
// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// return token.SignedString([]byte(constants.JWT_SECRET))

//return "", nil
//}
