package model

type Country struct {
	Name       Name `json:"name"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Timezones []string `json:"timezones"`
	Flags     Flag     `json:"flags"`
}

type Countries []Country

type CountryInfo struct {
	Name       Name `json:"name"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Timezone string `json:"timezone"`
}

type Name struct {
	Common string `json:"common"`
}

type Flag struct {
	PNG string `json:"png"`
}
