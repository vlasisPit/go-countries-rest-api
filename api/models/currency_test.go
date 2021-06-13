package models

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestCurrencyDeserialization(t *testing.T) {
	currencyBytes, err := ioutil.ReadFile("currency.json")
	var currency Currency
	err = json.Unmarshal(currencyBytes, &currency)
	if err != nil {
		t.Errorf("Can not deserialize json to struct")
	}

	assert.Equal(t, "Euro", currency.Name)
	assert.Equal(t, "EUR", currency.Code)
	assert.Equal(t, "E", currency.Symbol)
}