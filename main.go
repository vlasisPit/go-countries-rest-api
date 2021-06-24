package main

import (
	"go-countries-rest-api/api/server"
)

func main() {
	app := server.App{Port: ":8080"}
	app.Run()
}
