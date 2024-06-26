package routes

import (
	"project/controllers"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()

	//instance users
	userController := controllers.NewUserController(db)
	//menampilkan semua user
	e.GET("/users", userController.GetUsersController)
	//menampilkan user by id
	e.GET("/users/:id", userController.GetUserByIdController)
	//menambahkan user s
	e.POST("/users", userController.CreateUserController)
	//mengedit user by id
	e.PUT("/users/:id", userController.UpdateUserController)
	//menghapus user by id
	e.DELETE("/users/:id", userController.DeleteUserController)

	//instance book
	bookController := controllers.NewBookController(db)
	//menampilkan semua buku
	e.GET("/books", bookController.GetBooksController)
	//menampilkan buku by id
	e.GET("/books/:id", bookController.GetBookByIdController)
	//menambahkan buku
	e.POST("/books", bookController.CreateBookController)
	//mengedit buku by id
	e.PUT("/books/:id", bookController.UpdateBookController)
	//menghapus buku by id
	e.DELETE("/books/:id", bookController.DeleteBookByIdController)

	//instance blog
	blogController := controllers.NewBlogsController(db)
	//menampilkan semua blog
	e.GET("/blogs", blogController.GetBlogsController)
	//menampilkan blog by id
	e.GET("/blogs/:id", blogController.GetBlogsByIDController)
	//menambahkan blog
	e.POST("/blogs", blogController.CreateBlogsController)
	//mengedit blog by id
	e.PUT("/blogs/:id", blogController.UpdateBlogsController)
	//menghapus blog by id
	e.DELETE("/blogs/:id", blogController.DeleteBlogsController)

	return e
}
