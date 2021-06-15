package main

/**
For each, `newCountriesHandlers` is called. So a new empty internal storage is created
*/

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	storage "go-countries-rest-api/api/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	model "go-countries-rest-api/api/models"
)

func TestGetAllCountriesWithEmptyMemory(t *testing.T) {
	mux := initializeHandlers()
	getAllReq, _ := http.NewRequest("GET", "/countries", nil)
	getAllReqRecorder := newRequestRecorder(getAllReq, mux)
	assert.Equal(t, http.StatusOK, getAllReqRecorder.Code)
	assert.Equal(t, "[]", getAllReqRecorder.Body.String())
}

func TestAddOneCountryAndGetAllCountries(t *testing.T) {
	mux := initializeHandlers()

	body := "{\"name\": \"Greece\",\"alpha2Code\": \"GR\",\"capital\": \"Athens\",\"currencies\": [{\"code\": \"EUR\",\"name\": \"Euro\",\"symbol\": \"E\"}]}"
	addReq, _ := http.NewRequest("POST", "/countries", strings.NewReader(body))
	addReq.Header.Add("Content-Type", "application/json")
	addReqRecorder := newRequestRecorder(addReq, mux)
	assert.Equal(t, http.StatusOK, addReqRecorder.Code)
	assert.Equal(t, "", addReqRecorder.Body.String())

	getAllReq, _ := http.NewRequest("GET", "/countries", nil)
	getAllReqRecorder := newRequestRecorder(getAllReq, mux)
	actualCountries := constructCountriesFromJson(getAllReqRecorder.Body.String())
	assert.Equal(t, http.StatusOK, getAllReqRecorder.Code)
	assert.Equal(t, 1, len(*actualCountries))
	assert.Equal(t, "Greece", (*actualCountries)[0].Name)
	assert.Equal(t, "GR", (*actualCountries)[0].Alpha2Code)
	assert.Equal(t, "Athens", (*actualCountries)[0].Capital)
	assert.Equal(t, "Euro", (*actualCountries)[0].Currencies[0].Name)
	assert.Equal(t, "EUR", (*actualCountries)[0].Currencies[0].Code)
}

func TestAddTwoCountriesAndGetAllCountries(t *testing.T) {
	mux := initializeHandlers()

	bodyGr := "{\"name\": \"Greece\",\"alpha2Code\": \"GR\",\"capital\": \"Athens\",\"currencies\": [{\"code\": \"EUR\",\"name\": \"Euro\",\"symbol\": \"E\"}]}"
	addReqGr, _ := http.NewRequest("POST", "/countries", strings.NewReader(bodyGr))
	addReqGr.Header.Add("Content-Type", "application/json")
	addReqRecorderGr := newRequestRecorder(addReqGr, mux)
	assert.Equal(t, http.StatusOK, addReqRecorderGr.Code)
	assert.Equal(t, "", addReqRecorderGr.Body.String())

	bodySp := "{\"name\": \"Spain\",\"alpha2Code\": \"ES\",\"capital\": \"Madrid\",\"currencies\": [{\"code\": \"EUR\",\"name\": \"Euro\",\"symbol\": \"E\"}]}"
	addReqSp, _ := http.NewRequest("POST", "/countries", strings.NewReader(bodySp))
	addReqSp.Header.Add("Content-Type", "application/json")
	addReqRecorderSp := newRequestRecorder(addReqSp, mux)
	assert.Equal(t, http.StatusOK, addReqRecorderSp.Code)
	assert.Equal(t, "", addReqRecorderSp.Body.String())

	getAllReq, _ := http.NewRequest("GET", "/countries", nil)
	getAllReqRecorder := newRequestRecorder(getAllReq, mux)
	actualCountries := constructCountriesFromJson(getAllReqRecorder.Body.String())
	assert.Equal(t, http.StatusOK, getAllReqRecorder.Code)
	assert.Equal(t, 2, len(*actualCountries))
	assert.Equal(t, "Greece", (*actualCountries)[0].Name)
	assert.Equal(t, "GR", (*actualCountries)[0].Alpha2Code)
	assert.Equal(t, "Athens", (*actualCountries)[0].Capital)
	assert.Equal(t, "Euro", (*actualCountries)[0].Currencies[0].Name)
	assert.Equal(t, "EUR", (*actualCountries)[0].Currencies[0].Code)
	assert.Equal(t, "Spain", (*actualCountries)[1].Name)
	assert.Equal(t, "ES", (*actualCountries)[1].Alpha2Code)
	assert.Equal(t, "Madrid", (*actualCountries)[1].Capital)
	assert.Equal(t, "Euro", (*actualCountries)[1].Currencies[0].Name)
}

func TestAddTwoCountriesAndGetSpecificCountry(t *testing.T) {
	mux := initializeHandlers()

	bodyGr := "{\"name\": \"Greece\",\"alpha2Code\": \"GR\",\"capital\": \"Athens\",\"currencies\": [{\"code\": \"EUR\",\"name\": \"Euro\",\"symbol\": \"E\"}]}"
	addReqGr, _ := http.NewRequest("POST", "/countries", strings.NewReader(bodyGr))
	addReqGr.Header.Add("Content-Type", "application/json")
	addReqRecorderGr := newRequestRecorder(addReqGr, mux)
	assert.Equal(t, http.StatusOK, addReqRecorderGr.Code)
	assert.Equal(t, "", addReqRecorderGr.Body.String())

	bodySp := "{\"name\": \"Spain\",\"alpha2Code\": \"ES\",\"capital\": \"Madrid\",\"currencies\": [{\"code\": \"EUR\",\"name\": \"Euro\",\"symbol\": \"E\"}]}"
	addReqSp, _ := http.NewRequest("POST", "/countries", strings.NewReader(bodySp))
	addReqSp.Header.Add("Content-Type", "application/json")
	addReqRecorderSp := newRequestRecorder(addReqSp, mux)
	assert.Equal(t, http.StatusOK, addReqRecorderSp.Code)
	assert.Equal(t, "", addReqRecorderSp.Body.String())

	getGreeceReq, _ := http.NewRequest("GET", "/countries/greece", nil)
	getAllReqRecorder := newRequestRecorder(getGreeceReq, mux)
	actualCountry := constructCountryFromJson(getAllReqRecorder.Body.String())
	assert.Equal(t, http.StatusOK, getAllReqRecorder.Code)
	assert.Equal(t, "Greece", actualCountry.Name)
	assert.Equal(t, "GR", actualCountry.Alpha2Code)
	assert.Equal(t, "Athens", actualCountry.Capital)
	assert.Equal(t, "Euro", actualCountry.Currencies[0].Name)
	assert.Equal(t, "EUR", actualCountry.Currencies[0].Code)
}

func TestAddTwoCountriesAndDeleteSpecificCountry(t *testing.T) {
	mux := initializeHandlers()

	bodyGr := "{\"name\": \"Greece\",\"alpha2Code\": \"GR\",\"capital\": \"Athens\",\"currencies\": [{\"code\": \"EUR\",\"name\": \"Euro\",\"symbol\": \"E\"}]}"
	addReqGr, _ := http.NewRequest("POST", "/countries", strings.NewReader(bodyGr))
	addReqGr.Header.Add("Content-Type", "application/json")
	addReqRecorderGr := newRequestRecorder(addReqGr, mux)
	assert.Equal(t, http.StatusOK, addReqRecorderGr.Code)
	assert.Equal(t, "", addReqRecorderGr.Body.String())

	bodySp := "{\"name\": \"Spain\",\"alpha2Code\": \"ES\",\"capital\": \"Madrid\",\"currencies\": [{\"code\": \"EUR\",\"name\": \"Euro\",\"symbol\": \"E\"}]}"
	addReqSp, _ := http.NewRequest("POST", "/countries", strings.NewReader(bodySp))
	addReqSp.Header.Add("Content-Type", "application/json")
	addReqRecorderSp := newRequestRecorder(addReqSp, mux)
	assert.Equal(t, http.StatusOK, addReqRecorderSp.Code)
	assert.Equal(t, "", addReqRecorderSp.Body.String())

	getAllReq, _ := http.NewRequest("GET", "/countries", nil)
	getAllReqRecorder := newRequestRecorder(getAllReq, mux)
	actualCountries := constructCountriesFromJson(getAllReqRecorder.Body.String())
	assert.Equal(t, http.StatusOK, getAllReqRecorder.Code)
	assert.Equal(t, 2, len(*actualCountries))

	deleteReq, _ := http.NewRequest("DELETE", "/countries/spain", nil)
	deleteReqRecorder := newRequestRecorder(deleteReq, mux)
	assert.Equal(t, http.StatusOK, deleteReqRecorder.Code)

	getAllReqRecorder2 := newRequestRecorder(getAllReq, mux)
	actualCountriesAfterDelete := constructCountriesFromJson(getAllReqRecorder2.Body.String())
	assert.Equal(t, http.StatusOK, getAllReqRecorder2.Code)
	assert.Equal(t, 1, len(*actualCountriesAfterDelete))

	getSpainReq, _ := http.NewRequest("GET", "/countries/spain", nil)
	getSpainReqRecorder := newRequestRecorder(getSpainReq, mux)
	assert.Equal(t, http.StatusNotFound, getSpainReqRecorder.Code)
}

func constructCountryFromJson(jsonData string) *model.Country {
	country := &model.Country{}
	json.Unmarshal([]byte(jsonData), country)
	return country
}

func constructCountriesFromJson(jsonData string) *[]model.Country {
	countries := &[]model.Country{}
	json.Unmarshal([]byte(jsonData), countries)
	return countries
}

// Mocks a handler and returns a httptest.ResponseRecorder
func newRequestRecorder(req *http.Request, mux *http.ServeMux) *httptest.ResponseRecorder {
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	mux.ServeHTTP(rr, req)
	return rr
}

func initializeHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	countriesStorage := storage.NewCountriesStorage()
	server := Server{
		mux:     mux,
		actions: countriesStorage,
	}
	mux.HandleFunc("/countries", server.countries)
	mux.HandleFunc("/countries/", server.countryById)
	return mux
}