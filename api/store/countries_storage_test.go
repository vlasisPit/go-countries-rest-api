package store

import (
	"github.com/stretchr/testify/assert"
	"go-countries-rest-api/api/models"
	"testing"
)

func TestStorageGetAllCountriesWithEmptyMemory(t *testing.T) {
	storage := NewCountriesStorage()
	actualCountries, getAllCountriesError := storage.GetAllCountries()
	assert.Equal(t, nil, getAllCountriesError)
	assert.Equal(t, 0, len(*actualCountries))
}

func TestStorageAddOneCountryAndGetAllCountries(t *testing.T) {
	storage := NewCountriesStorage()
	country := constructCountryGreece()
	_, addCountryError := storage.AddCountry(country)
	actualCountries, getAllCountriesError := storage.GetAllCountries()
	assert.Nil(t, addCountryError)
	assert.Nil(t, getAllCountriesError)
	assert.Equal(t, 1, len(*actualCountries))
	assert.Equal(t, "Greece", (*actualCountries)[0].Name)
	assert.Equal(t, "GR", (*actualCountries)[0].Alpha2Code)
	assert.Equal(t, "Athens", (*actualCountries)[0].Capital)
	assert.Equal(t, "Euro", (*actualCountries)[0].Currencies[0].Name)
	assert.Equal(t, "EUR", (*actualCountries)[0].Currencies[0].Code)
}

func TestStorageAddTwoCountriesAndGetAllCountries(t *testing.T) {
	storage := NewCountriesStorage()
	greece := constructCountryGreece()
	spain := constructCountrySpain()
	_, addGreeceCountryError := storage.AddCountry(greece)
	_, addSpainCountryError := storage.AddCountry(spain)
	actualCountries, getAllCountriesError := storage.GetAllCountries()
	assert.Nil(t, addGreeceCountryError)
	assert.Nil(t, addSpainCountryError)
	assert.Nil(t, getAllCountriesError)
	assert.Equal(t, 2, len(*actualCountries))
}

func TestStorageAddTwoCountriesAndGetSpecificCountry(t *testing.T) {
	storage := NewCountriesStorage()
	greece := constructCountryGreece()
	spain := constructCountrySpain()
	_, addGreeceCountryError := storage.AddCountry(greece)
	_, addSpainCountryError := storage.AddCountry(spain)
	actualCountries, getAllCountriesError := storage.GetAllCountries()
	assert.Nil(t, addGreeceCountryError)
	assert.Nil(t, addSpainCountryError)
	assert.Nil(t, getAllCountriesError)
	assert.Equal(t, 2, len(*actualCountries))

	actual, addGreeceCountryError := storage.GetCountryById("greece")
	assert.Equal(t, "Greece", actual.Name)
	assert.Equal(t, "GR", actual.Alpha2Code)
	assert.Equal(t, "Athens", actual.Capital)
	assert.Equal(t, "Euro", actual.Currencies[0].Name)
	assert.Equal(t, "EUR", actual.Currencies[0].Code)
}

func TestStorageAddTwoCountriesAndDeleteSpecificCountry(t *testing.T) {
	storage := NewCountriesStorage()
	greece := constructCountryGreece()
	spain := constructCountrySpain()
	_, addGreeceCountryError := storage.AddCountry(greece)
	_, addSpainCountryError := storage.AddCountry(spain)
	actualCountries, getAllCountriesError := storage.GetAllCountries()
	assert.Nil(t, addGreeceCountryError)
	assert.Nil(t, addSpainCountryError)
	assert.Nil(t, getAllCountriesError)
	assert.Equal(t, 2, len(*actualCountries))

	deleteSpainCountryError := storage.DeleteCountry("spain")
	assert.Nil(t, deleteSpainCountryError)

	actualCountriesAfterDeletion, getAllCountriesErrorAfterDeletion := storage.GetAllCountries()
	assert.Nil(t, getAllCountriesErrorAfterDeletion)
	assert.Equal(t, 1, len(*actualCountriesAfterDeletion))

	actual, addGreeceCountryError := storage.GetCountryById("greece")
	assert.Equal(t, "Greece", actual.Name)
	assert.Equal(t, "GR", actual.Alpha2Code)
	assert.Equal(t, "Athens", actual.Capital)
	assert.Equal(t, "Euro", actual.Currencies[0].Name)
	assert.Equal(t, "EUR", actual.Currencies[0].Code)
}

func TestStorageAddTwoCountriesAndGetRandomCountry(t *testing.T) {
	storage := NewCountriesStorage()
	greece := constructCountryGreece()
	spain := constructCountrySpain()
	_, addGreeceCountryError := storage.AddCountry(greece)
	_, addSpainCountryError := storage.AddCountry(spain)
	actualCountries, getAllCountriesError := storage.GetAllCountries()
	assert.Nil(t, addGreeceCountryError)
	assert.Nil(t, addSpainCountryError)
	assert.Nil(t, getAllCountriesError)
	assert.Equal(t, 2, len(*actualCountries))

	actual, randomCountryError := storage.GetRandomCountryId()
	assert.Nil(t, randomCountryError)
	assert.Contains(t, [2]string{"greece", "spain"}, *actual)
}

func TestStorageAddOneCountriesAndGetRandomCountry(t *testing.T) {
	storage := NewCountriesStorage()
	greece := constructCountryGreece()
	_, addGreeceCountryError := storage.AddCountry(greece)
	actualCountries, getAllCountriesError := storage.GetAllCountries()
	assert.Nil(t, addGreeceCountryError)
	assert.Nil(t, getAllCountriesError)
	assert.Equal(t, 1, len(*actualCountries))

	actual, randomCountryError := storage.GetRandomCountryId()
	assert.Nil(t, randomCountryError)
	assert.Equal(t, "greece", *actual)
}

func TestStorageNoCountryAddedAndGetRandomCountry(t *testing.T) {
	storage := NewCountriesStorage()
	actualCountries, getAllCountriesError := storage.GetAllCountries()
	assert.Nil(t, getAllCountriesError)
	assert.Equal(t, 0, len(*actualCountries))

	actual, randomCountryError := storage.GetRandomCountryId()
	assert.Equal(t, "No countries available to choose randomly.", randomCountryError.Error())
	assert.Nil(t, actual)
}

func TestStorageAddTwoCountriesAndGetNotExistingCountry(t *testing.T) {
	storage := NewCountriesStorage()
	greece := constructCountryGreece()
	spain := constructCountrySpain()
	_, addGreeceCountryError := storage.AddCountry(greece)
	_, addSpainCountryError := storage.AddCountry(spain)
	actualCountries, getAllCountriesError := storage.GetAllCountries()
	assert.Nil(t, addGreeceCountryError)
	assert.Nil(t, addSpainCountryError)
	assert.Nil(t, getAllCountriesError)
	assert.Equal(t, 2, len(*actualCountries))

	actual, addFranceCountryError := storage.GetCountryById("france")
	assert.Equal(t, "Country not found.", addFranceCountryError.Error())
	assert.Nil(t, actual)
}

func constructCountryGreece() models.Country {
	return models.Country{
		Name:       "Greece",
		Alpha2Code: "GR",
		Capital:    "Athens",
		Currencies: []models.Currency{{Code: "EUR", Name: "Euro", Symbol: "E"}},
	}
}

func constructCountrySpain() models.Country {
	return models.Country{
		Name:       "Spain",
		Alpha2Code: "ES",
		Capital:    "Madrid",
		Currencies: []models.Currency{{Code: "EUR", Name: "Euro", Symbol: "E"}},
	}
}