test_all:
	go test -v -cover ./...

docker_build:
	docker build -t go-countries-rest-api .

docker_run:
	docker run -p 8080:8080 go-countries-rest-api

go_run:
	go run main.go

