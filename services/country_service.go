package services

import (
	"encoding/json"
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/model"
	"net/http"
)

func GetCountryInfo(countryCode string) (countryInfo model.CountryInfo, e error) {
	url := fmt.Sprintf("https://restcountries.com/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(url)
	if err != nil {
		return model.CountryInfo{}, err
	}
	decoder := json.NewDecoder(resp.Body)

	var countries model.Countries

	err = decoder.Decode(&countries)
	if err != nil {
		return model.CountryInfo{}, err
	}

	if len(countries) < 1 {
		return model.CountryInfo{}, fmt.Errorf("country code is not asociated with an existing country")
	}

	country := countries[0]
	return model.CountryInfo{
			Name:       country.Name,
			Currencies: country.Currencies,
			Timezone:   country.Timezones[0],
		},
		nil
}

func NotCountryCode(countryCode string) (isNotCountry bool) {
	_, err := GetCountryInfo(countryCode)
	if err != nil {
		return true
	}
	return false
}
