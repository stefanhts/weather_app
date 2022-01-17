package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"weather_app/API"
	"weather_app/structs"
)

var tmpl *template.Template

const US = "US"
const NA = "NAM"
const CITY_CODE = "333558"
const API_KEY = API.API_KEY

func weather(w http.ResponseWriter, r *http.Request) {
	var currentWeather []structs.CurrenntWeather
	uri := fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/%s?apikey=%s", CITY_CODE, API_KEY)
	resp, err := http.Get(uri)
	if err != nil {
		// TODO handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		print(err.Error())
	} else {
		json.Unmarshal([]byte(body), &currentWeather)
	}
	tmpl.Execute(w, currentWeather[0])
}

func GetCityCode(city string) string {
	var store []structs.WeatherResp
	uri := fmt.Sprintf("http://dataservice.accuweather.com/locations/v1/cities/search?apikey=%s&q=%s", API_KEY, city)
	resp, err := http.Get(uri)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &store)
	if err != nil {
		print(err.Error())
	}
	return store[0].Key
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	mux.HandleFunc("/weather", weather)
	log.Fatal(http.ListenAndServe(":5000", mux))
}
