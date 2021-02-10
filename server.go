package main

import (
	"encoding/json"
	"net/http"
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
	store map[string]Country
}

/*
Method pointer receiver to countriesHandler
https://tour.golang.org/methods/4
*/
func (h *countriesHandler) get(writer http.ResponseWriter, request *http.Request) {
	countries := make([]Country, len(h.store))

	i := 0
	for _, country := range h.store {
		countries[i] = country
		i++
	}

	jsonBytes, err := json.Marshal(countries)
	if err != nil {
		//TODO
	}

	writer.Write(jsonBytes)
}

/*
Never have a nil map
*/
func newCountriesHandlers() *countriesHandler {
	return &countriesHandler{
		store: map[string]Country{
			"greece": {
				Name:       "Greece",
				Alpha2Code: "GR",
				Capital:    "Athens",
				Currencies: []Currency{{Code: "EUR", Name: "Euro", Symbol: "E"}},
			},
		},
	}
}

func main() {
	countriesHandler := newCountriesHandlers()
	http.HandleFunc("/countries", countriesHandler.get)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
