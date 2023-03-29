package main

import (
	"bytes"
	"encoding/json"
	"net/http"

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

	e.POST("/posts", func(c echo.Context) error {
		// Get post data from request body
		var post Post
		if err := c.Bind(&post); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		// Marshal post data to JSON
		jsonData, err := json.Marshal(post)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}

		// Create new request to POST data to API
		request, err := http.NewRequest(http.MethodPost, "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(jsonData))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		request.Header.Set("Content-Type", "application/json")

		// Send request to API
		client := http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		defer response.Body.Close()

		// Parse response body
		var result Post
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}

		// Return response
		return c.JSON(http.StatusCreated, result)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
