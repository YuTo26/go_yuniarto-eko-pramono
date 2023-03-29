package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	e := echo.New()

	e.GET("/posts/:id", func(c echo.Context) error {
		// Get id from parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid post ID")
		}

		// Get data from API
		response, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(id))
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Post not found")
		}
		defer response.Body.Close()

		// Parse response body
		var post Post
		if err := json.NewDecoder(response.Body).Decode(&post); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Post not found")
		}

		// Return response
		return c.JSON(http.StatusOK, post)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
