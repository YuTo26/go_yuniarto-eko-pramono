package controllers

import (
	"net/http"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{DB: db}
}

func (controller *BookController) GetBooksController(c echo.Context) error {
	var books []models.Book
	if err := controller.DB.Find(&books).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, books)
}

func (controller *BookController) GetBookByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid parameter"})
	}
	var book models.Book
	if err := controller.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Book not found"})
	}
	return c.JSON(http.StatusOK, book)
}

func (controller *BookController) CreateBookController(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	if err := controller.DB.Create(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, book)
}

func (controller *BookController) UpdateBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid parameter"})
	}
	var book models.Book
	if err := controller.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Book not found"})
	}
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	if err := controller.DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, book)
}

func (controller *BookController) DeleteBookByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid parameter"})
	}

	var book models.Book
	if err := controller.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Book not found"})
	}

	if err := controller.DB.Delete(&book, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, book)
}
