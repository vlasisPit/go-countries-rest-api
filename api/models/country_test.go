package models

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestCountryDeserialization(t *testing.T) {
	countryBytes, err := ioutil.ReadFile("country.json")
	var country Country
	err = json.Unmarshal(countryBytes, &country)
	if err != nil {
		t.Errorf("Can not deserialize json to struct")
	}

	assert.Equal(t, "Greece", country.Name)
	assert.Equal(t, "GR", country.Alpha2Code)
	assert.Equal(t, "Athens", country.Capital)
	assert.Equal(t, 1, len(country.Currencies))
	assert.Equal(t, "Euro", country.Currencies[0].Name)
	assert.Equal(t, "EUR", country.Currencies[0].Code)
}