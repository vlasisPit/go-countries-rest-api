package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	model "go-countries-rest-api/api/models"
)

type countriesHandler struct {
	sync.Mutex
	store map[string]model.Country
}

/*
Method pointer receiver to countriesHandler
https://tour.golang.org/methods/4
*/
func (h *countriesHandler) get(writer http.ResponseWriter, request *http.Request) {
	countries := make([]model.Country, len(h.store))

	h.Lock()
	i := 0
	for _, country := range h.store {
		countries[i] = country
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(countries)
	if err != nil {
		constructErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonBytes)
}

func (h *countriesHandler) getRandomCountry(writer http.ResponseWriter, request *http.Request) {
	ids := make([]string, len(h.store))
	h.Lock()
	i := 0
	for id := range h.store {
		ids[i] = id
		i++
	}
	h.Unlock()

	var target string
	if len(ids) == 0 {
		constructErrorResponse(writer, "No countries available to choose randomly", http.StatusNotFound)
		return
	} else if len(ids) == 1 {
		target = ids[0]
	} else {
		rand.Seed(time.Now().UnixNano())
		target = ids[rand.Intn(len(ids))]
	}

	//redirect
	writer.Header().Add("location", fmt.Sprintf("/countries/%s", target))
	writer.WriteHeader(http.StatusFound)
}

/**
Handle requests with path "/countries/{id}" like
GET /countries/{id}
 */
func (h *countriesHandler) getCountry(writer http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.String(), "/")
	if len(parts) != 3 {
		constructErrorResponse(writer, "Wrong number of parts on URL path", http.StatusNotFound)
		return
	}

	if parts[2] == "random" {
		h.getRandomCountry(writer, request)
		return
	}

	h.Lock()
	country, ok := h.store[strings.ToLower(parts[2])]
	if !ok {
		constructErrorResponse(writer, "Country not found", http.StatusNotFound)
		h.Unlock()
		return
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(country)
	if err != nil {
		constructErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonBytes)
}

func (h *countriesHandler) post(writer http.ResponseWriter, request *http.Request) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		constructErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	ct := request.Header.Get("content-type")
	if ct != "application/json" {
		constructErrorResponse(writer, fmt.Sprintf("need content-type 'application/json', but got '%s'", ct), http.StatusUnsupportedMediaType)
		return
	}

	var country model.Country
	err = json.Unmarshal(bodyBytes, &country)
	if err != nil {
		constructErrorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}

	h.Lock()
	h.store[strings.ToLower(country.Name)] = country
	defer h.Unlock()
}

/**
Handle (delete) requests with path "/countries/{id}" like
DELETE /countries/{id}
*/
func (h *countriesHandler) deleteCountry(writer http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.String(), "/")
	if len(parts) != 3 {
		constructErrorResponse(writer, "Wrong number of parts on URL path", http.StatusNotFound)
		return
	}

	h.Lock()
	delete(h.store, strings.ToLower(parts[2]))
	defer h.Unlock()

	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func constructErrorResponse(writer http.ResponseWriter, errorMessage string, serverError int) {
	writer.WriteHeader(serverError)
	writer.Write([]byte(errorMessage))
}

/**
Handle requests with path "/countries" like
GET /countries
POST /countries
 */
func (h *countriesHandler) countries(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		h.get(writer, request)
		return
	case "POST":
		h.post(writer, request)
		return
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("method not allowed"))
		return
	}
}

/**
Handle requests with path "/countries" like
GET /countries
POST /countries
*/
func (h *countriesHandler) countryById(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		h.getCountry(writer, request)
		return
	case "DELETE":
		h.deleteCountry(writer, request)
		return
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("method not allowed"))
		return
	}
}

/*
Never have a nil map
*/
func newCountriesHandlers() *countriesHandler {
	return &countriesHandler{
		store: map[string]model.Country{
/*		store: map[string]Country{
			"greece": {
				Name:       "Greece",
				Alpha2Code: "GR",
				Capital:    "Athens",
				Currencies: []Currency{{Code: "EUR", Name: "Euro", Symbol: "E"}},
			},*/
		},
	}
}

/**
According to https://www.alexedwards.net/blog/a-recap-of-request-handling
you should not use "http.HandleFunc" because of a security vulnerability issue.
Use "mux := http.NewServeMux()" instead
So as a rule of thumb it's a good idea to avoid the DefaultServeMux, and instead
use your own locally-scoped ServeMux, like we have been so far.
Check section "The DefaultServeMux" on article.
*/
func initialize(port string, mux *http.ServeMux, countriesHandler *countriesHandler) {
	mux.HandleFunc("/countries", countriesHandler.countries)
	mux.HandleFunc("/countries/", countriesHandler.countryById)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		panic(err)
	}
}

/**
https://dev.to/bmf_san/introduction-to-url-router-from-scratch-with-golang-3p8j
https://github.com/gsingharoy/httprouter-tutorial/tree/master/part4
Check this about routing
*/
type App struct {
}

func (a *App) Run(port string) {
	mux := http.NewServeMux()
	countriesHandler := newCountriesHandlers()
	initialize(port, mux, countriesHandler)
}