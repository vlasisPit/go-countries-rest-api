# Golang tutorial. Build a simple REST API
A simple CRUD API written in Golang (just for training purposes). The user can add, retrieve or delete countries from this REST API. All countries are stored in-memory. As I said before, this is a project just for learning the basics features of Golang. For a more professional approach a database must be used instead (a Postgres or a MongoDB database maybe). To use a database instead, the developer must use the `Actions` interface and implement the methods inside this interface.   

### Things done
*  `GET /countries` returns list of countries as JSON
*  `GET /countries/{id}` returns some details of a specific country as JSON
*  `POST /countries` accepts a new country to be added
*  `POST /countries` returns status 415 if content is not `application/json`
*  `GET /countries/random` redirects (Status 302) to a random country
*  `DELETE /countries/{id}` delete a specific country

### Curl samples

```
GET /countries
----
curl --request GET \
  --url http://localhost:8080/countries
```

```
GET /countries/{id}
----
curl --request GET \
  --url http://localhost:8080/countries/greece
```

```
POST /countries
----
curl --request POST \
  --url http://localhost:8080/countries \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Greece",
	"alpha2Code": "GR",
	"capital": "Athens",
	"currencies": [
		{
			"code": "EUR",
			"name": "Euro",
			"symbol": "E"
		}
	]
}'
```

```
GET /countries/{id}
----
curl --request GET \
  --url http://localhost:8080/countries/random
```

```
DELETE /countries/random
----
curl --request DELETE \
  --url http://localhost:8080/countries/spain
```

### MakeFile
*  `test_all` run all tests with coverage
*  `docker_build` build application's docker image
*  `docker_run` run application as a docker container
*  `go_run` run Golang application
