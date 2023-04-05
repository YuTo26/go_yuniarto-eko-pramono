package main

import (
	"project/config"
	"project/routes"
)

func main() {
	config.InitDB()

	e := routes.New(config.DB)
	e.Logger.Fatal(e.Start(":8080"))
}
