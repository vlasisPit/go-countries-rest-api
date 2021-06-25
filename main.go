package main

import (
	"go-countries-rest-api/api"
)

func main() {
	app := api.App{Port: ":8080"}
	app.Run()
}
