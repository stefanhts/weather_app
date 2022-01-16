package structs

import "time"

type Weather struct {
	Temp  int
	Weath string
	Day   string
}

type PageData struct {
	Weaths []Weather
}
type WeatherResp struct {
	Version           int    `json:"Version"`
	Key               string `json:"Key"`
	Type              string `json:"Type"`
	Rank              int    `json:"Rank"`
	LocalizedName     string `json:"LocalizedName"`
	EnglishName       string `json:"EnglishName"`
	PrimaryPostalCode string `json:"PrimaryPostalCode"`
	Region            struct {
		ID            string `json:"ID"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
	} `json:"Region"`
	Country struct {
		ID            string `json:"ID"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
	} `json:"Country"`
	AdministrativeArea struct {
		ID            string `json:"ID"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
		Level         int    `json:"Level"`
		LocalizedType string `json:"LocalizedType"`
		EnglishType   string `json:"EnglishType"`
		CountryID     string `json:"CountryID"`
	} `json:"AdministrativeArea"`
	TimeZone struct {
		Code             string    `json:"Code"`
		Name             string    `json:"Name"`
		GmtOffset        int       `json:"GmtOffset"`
		IsDaylightSaving bool      `json:"IsDaylightSaving"`
		NextOffsetChange time.Time `json:"NextOffsetChange"`
	} `json:"TimeZone"`
	GeoPosition struct {
		Latitude  float64 `json:"Latitude"`
		Longitude float64 `json:"Longitude"`
		Elevation struct {
			Metric struct {
				Value    int    `json:"Value"`
				Unit     string `json:"Unit"`
				UnitType int    `json:"UnitType"`
			} `json:"Metric"`
			Imperial struct {
				Value    int    `json:"Value"`
				Unit     string `json:"Unit"`
				UnitType int    `json:"UnitType"`
			} `json:"Imperial"`
		} `json:"Elevation"`
	} `json:"GeoPosition"`
}
