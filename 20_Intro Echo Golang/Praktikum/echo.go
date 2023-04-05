package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid user id",
		})
	}

	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user by id",
				"user":     user,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid user data",
		})
	}

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid user id",
		})
	}

	for i, user := range users {
		if user.Id == id {
			updatedUser := User{}
			if err := c.Bind(&updatedUser); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"messages": "invalid user data",
				})
			}

			updatedUser.Id = user.Id
			users[i] = updatedUser

			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update user",
				"user":     updatedUser,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid user id",
		})
	}

	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)

			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete user",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

func main() {
	e := echo.New()

	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":1010"))
}
