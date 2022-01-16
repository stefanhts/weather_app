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
const CITY = "Bethesda"
const API_KEY = API.API_KEY

func weather(w http.ResponseWriter, r *http.Request) {
	data := structs.PageData{
		Weaths: []structs.Weather{
			{Temp: 69, Weath: "rainy", Day: "Monday"},
			{Temp: 32, Weath: "snowy", Day: "Tuesday"},
		},
	}
	tmpl.Execute(w, data)
}

func city_search(city string) string {
	var store []structs.WeatherResp
	uri := fmt.Sprintf("http://dataservice.accuweather.com/locations/v1/cities/search?apikey=%s&q=%s", API_KEY, city)
	resp, err := http.Get(uri)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &store)
	print("test\n")
	if err != nil {
		print(err.Error())
	}
	sb := string(body)
	//log.Printf(sb)
	fmt.Println(store[0])

	return sb

	//return uri
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	city_search("Bethesda")
	mux.HandleFunc("/weather", weather)
	//print(outp)
	log.Fatal(http.ListenAndServe(":5000", mux))
}
