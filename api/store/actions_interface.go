package store

import "go-countries-rest-api/api/models"

type Actions interface {
	AddCountry(country models.Country) (*models.Country, error)
	DeleteCountry(countryId string) error
	GetCountryById(countryId string) (*models.Country, error)
	GetAllCountries() (*[]models.Country, error)

	/**
	Get Random Country Id to use it to redirect the call to GetCountryById
	 */
	GetRandomCountryId() (*string, error)
}
