package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseRate(t *testing.T) {
	json := `{
		"success": true,
		"query": {
			"from": "GBP",
			"to": "EUR",
			"amount": 1
		},
		"info": {
			"timestamp": 1664999883,
			"rate": 1.146072        
		},
		"date": "2022-10-05",       
		"result": 1.146072
	}`
	rate := Quote{Result: 1.146072}

	result, err := parseQuote([]byte(json))
	assert.Nil(t, err)
	assert.Equal(t, rate, result)
}

func Test_CompareRate(t *testing.T) {
	averageRate := 1.146072
	currentRate := 1.157183
	pips := computePips(currentRate, averageRate)
	assert.Equal(t, 111, pips)
}
