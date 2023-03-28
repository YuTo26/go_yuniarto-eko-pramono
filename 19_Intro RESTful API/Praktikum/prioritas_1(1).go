package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func fetchPosts() ([]Post, error) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var posts []Post
	err = json.NewDecoder(response.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func handleGetPosts(c echo.Context) error {
	posts, err := fetchPosts()
	if err != nil {
		log.Printf("Failed to fetch posts: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to fetch posts")
	}

	return c.JSON(http.StatusOK, posts)
}

func main() {
	e := echo.New()

	e.GET("/", handleGetPosts)

	if err := e.Start(":8080"); err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}
