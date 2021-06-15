package controllers

import (
	"errors"
	"go-countries-rest-api/api/models"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type CountriesStorage struct {
	sync.Mutex
	store map[string]models.Country
}

func NewCountriesStorage() *CountriesStorage {
	return &CountriesStorage{
		store: map[string]models.Country{},
	}
}

func (storage *CountriesStorage) AddCountry(country models.Country) (*models.Country, error) {
	storage.Lock()
	storage.store[strings.ToLower(country.Name)] = country
	defer storage.Unlock()
	return &country,nil
}

func (storage *CountriesStorage) DeleteCountry(countryId string) error {
	storage.Lock()
	delete(storage.store, strings.ToLower(countryId))
	defer storage.Unlock()
	return nil
}

func (storage *CountriesStorage) GetAllCountries() (*[]models.Country, error) {
	countries := make([]models.Country, len(storage.store))

	storage.Lock()
	i := 0
	for _, country := range storage.store {
		countries[i] = country
		i++
	}
	storage.Unlock()
	return &countries,nil
}

func (storage *CountriesStorage) GetCountryById(countryId string) (*models.Country, error) {
	storage.Lock()
	country, ok := storage.store[strings.ToLower(countryId)]
	if !ok {
		storage.Unlock()
		return nil,errors.New("Country not found.")
	}
	storage.Unlock()
	return &country,nil
}

func (storage *CountriesStorage) GetRandomCountryId() (*string, error) {
	ids := make([]string, len(storage.store))
	storage.Lock()
	i := 0
	for id := range storage.store {
		ids[i] = id
		i++
	}
	storage.Unlock()

	var target string
	if len(ids) == 0 {
		return nil,errors.New("No countries available to choose randomly.")
	} else if len(ids) == 1 {
		target = ids[0]
	} else {
		rand.Seed(time.Now().UnixNano())
		target = ids[rand.Intn(len(ids))]
	}
	return &target,nil
}
