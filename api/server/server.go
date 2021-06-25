package server

import (
	"encoding/json"
	"fmt"
	"go-countries-rest-api/api/store"
	model "go-countries-rest-api/api/models"
	utils "go-countries-rest-api/api/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
According to https://www.alexedwards.net/blog/a-recap-of-request-handling
you should not use "http.HandleFunc" because of a security vulnerability issue.
Use "mux := http.NewServeMux()" instead
So as a rule of thumb it's a good idea to avoid the DefaultServeMux, and instead
use your own locally-scoped ServeMux, like we have been so far.
Check section "The DefaultServeMux" on article.
*/
func (s *Server) Initialize(port string) {
	s.initializeRoutes()
	err := http.ListenAndServe(port, s.Mux)
	if err != nil {
		panic(err)
	}
}

type Server struct {
	Mux     *http.ServeMux
	Actions store.Actions
}

/*
Method pointer receiver to countriesHandler
https://tour.golang.org/methods/4
*/
func (s *Server) get(writer http.ResponseWriter, request *http.Request) {
	countries, _ := s.Actions.GetAllCountries()

	jsonBytes, err := json.Marshal(countries)
	if err != nil {
		utils.ConstructErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ConstructSuccessfulResponse(writer, http.StatusOK, jsonBytes)
}

func (s *Server) getRandomCountry(writer http.ResponseWriter, request *http.Request) {
	target, err := s.Actions.GetRandomCountryId()
	if err!=nil {
		utils.ConstructErrorResponse(writer, "No countries available to choose randomly", http.StatusNotFound)
		return
	}

	//redirect
	writer.Header().Add("location", fmt.Sprintf("/countries/%s", *target))
	writer.WriteHeader(http.StatusFound)
}

/**
Handle requests with path "/countries/{id}" like
GET /countries/{id}
 */
func (s *Server) getCountry(writer http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.String(), "/")
	if len(parts) != 3 {
		utils.ConstructErrorResponse(writer, "Wrong number of parts on URL path", http.StatusNotFound)
		return
	}

	if parts[2] == "random" {
		s.Actions.GetRandomCountryId()
		return
	}

	country, notFoundError := s.Actions.GetCountryById(parts[2])
	if notFoundError!=nil {
		utils.ConstructErrorResponse(writer, "Country not found", http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(country)
	if err != nil {
		utils.ConstructErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ConstructSuccessfulResponse(writer, http.StatusOK, jsonBytes)
}

func (s *Server) post(writer http.ResponseWriter, request *http.Request) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		utils.ConstructErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	ct := request.Header.Get("content-type")
	if ct != "application/json" {
		utils.ConstructErrorResponse(writer, fmt.Sprintf("need content-type 'application/json', but got '%s'", ct), http.StatusUnsupportedMediaType)
		return
	}

	var country model.Country
	err = json.Unmarshal(bodyBytes, &country)
	if err != nil {
		utils.ConstructErrorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}

	s.Actions.AddCountry(country)
}

/**
Handle (delete) requests with path "/countries/{id}" like
DELETE /countries/{id}
*/
func (s *Server) deleteCountry(writer http.ResponseWriter, request *http.Request) {
	parts := strings.Split(request.URL.String(), "/")
	if len(parts) != 3 {
		utils.ConstructErrorResponse(writer, "Wrong number of parts on URL path", http.StatusNotFound)
		return
	}

	s.Actions.DeleteCountry(parts[2])

	utils.ConstructSuccessfulResponse(writer, http.StatusOK, nil)
}

/**
Handle requests with path "/countries" like
GET /countries
POST /countries
 */
func (s *Server) countries(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		s.get(writer, request)
		return
	case "POST":
		s.post(writer, request)
		return
	default:
		utils.ConstructErrorResponse(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

/**
Handle requests with path "/countries" like
GET /countries
POST /countries
*/
func (s *Server) countryById(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		s.getCountry(writer, request)
		return
	case "DELETE":
		s.deleteCountry(writer, request)
		return
	default:
		utils.ConstructErrorResponse(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
