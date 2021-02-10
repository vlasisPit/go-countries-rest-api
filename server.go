package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type Country struct {
	Name       string     `json:"name"`
	Alpha2Code string     `json:"alpha2Code"`
	Capital    string     `json:"capital"`
	Currencies []Currency `json:"currencies"`
}

type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type countriesHandler struct {
	sync.Mutex
	store map[string]Country
}

/*
Method pointer receiver to countriesHandler
https://tour.golang.org/methods/4
*/
func (h *countriesHandler) get(writer http.ResponseWriter, request *http.Request) {
	countries := make([]Country, len(h.store))

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

	var country Country
	err = json.Unmarshal(bodyBytes, &country)
	if err != nil {
		constructErrorResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}

	h.Lock()
	h.store[strings.ToLower(country.Name)] = country
	defer h.Unlock()
}

func constructErrorResponse(writer http.ResponseWriter, errorMessage string, serverError int) {
	writer.WriteHeader(serverError)
	writer.Write([]byte(errorMessage))
}

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

/*
Never have a nil map
*/
func newCountriesHandlers() *countriesHandler {
	return &countriesHandler{
		store: map[string]Country{
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

func main() {
	countriesHandler := newCountriesHandlers()
	http.HandleFunc("/countries", countriesHandler.countries)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
