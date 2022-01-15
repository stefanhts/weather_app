package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Weather struct {
	Temp  int
	Weath string
	Day   string
}

type PageData struct {
	Weaths []Weather
}

func weather(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Weaths: []Weather{
			{Temp: 69, Weath: "rainy", Day: "Monday"},
			{Temp: 32, Weath: "snowy", Day: "Tuesday"},
		},
	}
	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	mux.HandleFunc("/weather", weather)
	log.Fatal(http.ListenAndServe(":5000", mux))
}
