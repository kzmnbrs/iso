# iso
Package contains ISO country/currency enums. Based on [github.com/countries/countries](https://github.com/countries/countries).

## Install
```bash
go get -u github.com/kzmnbrs/iso/go
```

Usage:
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"

	iso "github.com/kzmnbrs/iso/go"
)

type data struct {
	Country        iso.Country      `json:"country"`
	Currency       iso.Currency     `json:"currency"`
	SecondCurrency iso.NullCurrency `json:"second_currency"`
}

func main() {
	mustUnmarshal := func(s string, dst interface{}) {
		if err := json.Unmarshal([]byte(s), dst); err != nil {
			log.Panic(err)
		}
	}
	assert := func(x bool) {
		_, _, line, _ := runtime.Caller(1)
		if !x {
			log.Panicf("assert failed: line=%d", line)
		}
	}

	gbr := data{}
	mustUnmarshal(`{"country": "GBR", "currency": "GBP"}`, &gbr)
	assert(gbr.Country == iso.CountryGB)
	assert(gbr.Country == iso.CountryGBR)
	assert(gbr.Currency == iso.CurrencyGBP)
	assert(gbr.Country.Currency() == iso.CurrencyGBP)
	assert(!gbr.SecondCurrency.Valid)

	// Might be nil.
	assert(gbr.Country.PostalCodeRe().MatchString("BA21 5BT"))
	// Might be empty.
	assert(gbr.Country.PhonePrefix() == "44")

	gbr1 := data{}
	mustUnmarshal(`{"country": "GB", "currency": "GBP", "second_currency": null}`, &gbr1)
	assert(gbr == gbr1)

	pa := data{}
	mustUnmarshal(`{"country": "PA", "currency": "PAB", "second_currency": "USD"}`, &pa)
	assert(pa.Country == iso.CountryPAN)
	assert(pa.Country == iso.CountryPA)
	assert(pa.Currency == iso.CurrencyPAB)
	assert(pa.Country.Currency() == iso.CurrencyPAB)
	assert(pa.SecondCurrency.Valid && pa.SecondCurrency.Currency == iso.CurrencyUSD)

	allCountries := iso.AppendCountries(nil)
	assert(len(allCountries) == 249)

	allCurrencies := iso.AppendCurrencies(nil)
	assert(len(allCurrencies) == 152)

	ua := iso.CountryFromCode("UA")
	ukr := iso.CountryFromCode("UKR")
	// Optimized from code.
	ukr1 := iso.CountryFromAlpha3("UKR")
	assert(ua == ukr && ua == ukr1)

	// SQL stuff.
	fmt.Println(ua.Value())
	assert(ua.Scan != nil)
}
```
