package main

import "go-countries-rest-api/api/controllers"

func main() {
	app := controllers.App{port: ":8080"}
	app.Run()
}
