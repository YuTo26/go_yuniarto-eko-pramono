package main

import (
	"project_1/config"
	"project_1/routes"
)

func main() {
	config.InitDB()

	e := routes.New(config.DB)
	e.Logger.Fatal(e.Start(":8080"))
}
