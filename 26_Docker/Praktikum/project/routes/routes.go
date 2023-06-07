package routes

import (
	"project_1/constants"
	"project_1/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	jwtMiddleware := middleware.JWT([]byte(constants.JWT_SECRET))

	userController := controllers.NewUserController(db)
	bookController := controllers.NewBookController(db)

	// No auth
	e.POST("/register", userController.CreateUserController)
	e.POST("/login", userController.LoginController)

	// Auth
	authGroup := e.Group("")
	authGroup.Use(jwtMiddleware)

	authGroup.GET("/users", userController.GetUsersController)
	authGroup.GET("/users/:id", userController.GetUserByIdController)
	authGroup.PUT("/users/:id", userController.UpdateUserController)
	authGroup.DELETE("/users/:id", userController.DeleteUserController)

	authGroup.GET("/books", bookController.GetBooksController)
	authGroup.GET("/books/:id", bookController.GetBookByIdController)
	authGroup.POST("/books", bookController.CreateBookController)
	authGroup.PUT("/books/:id", bookController.UpdateBookController)
	authGroup.DELETE("/books/:id", bookController.DeleteBookByIdController)

	return e
}
