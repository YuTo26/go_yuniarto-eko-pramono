package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Handler untuk menghapus postingan dengan ID tertentu
	e.DELETE("/posts/:id", func(c echo.Context) error {
		postID := c.Param("id")

		request, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/posts/"+postID, nil)
		if err != nil {
			return err
		}

		client := http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			return c.String(http.StatusOK, "Post with ID "+postID+" has been deleted")
		} else {
			return c.String(http.StatusInternalServerError, "Failed to delete post with ID "+postID)
		}
	})

	e.Logger.Fatal(e.Start(":8080"))
}
