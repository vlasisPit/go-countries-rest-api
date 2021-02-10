# Golang tutorial. Build a simple REST API
Thanks to https://www.youtube.com/watch?v=2v11Ym6Ct9Q

### Things done
*  `GET /countries` returns list of countries as JSON
*  `GET /countries/{id}` returns details of a specific country as JSON
*  `POST /countries` accepts a new country to be added
*  `POST /countries` returns status 415 if content is not `application/json`
*  `GET /countries/{id}` redirects (Status 302) to a random country

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