package Utils

import (
	"encoding/json"
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/model"
	"net/http"
	"regexp"
)

func EmptyString(str string) (isEmpty bool) {
	return str == ""
}

func NotEmailString(email string) (isNotEmail bool) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return !emailRegex.MatchString(email)
}

func NotCountryCode(countryCode string) (isNotCountry bool) {
	url := fmt.Sprintf("https://restcountries.com/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err:", err)
		return true
	}
	decoder := json.NewDecoder(resp.Body)

	var countries model.Countries
	err = decoder.Decode(&countries)
	if err != nil {
		fmt.Println(err)
	}
	return false
}
