package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type ForecastResponse struct {
	List []ForecastItem  `json:"list"`
}

type ForecastItem struct {
	Dt   int `json:"dt"`
	Main struct {
		Temp      float64 `json:"temp"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
		Humidity  int     `json:"humidity"`
		TempKf    float64 `json:"temp_kf"`
	} `json:"main"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Snow struct {
		ThreeH float64 `json:"3h"`
	} `json:"snow"`
	Sys struct {
		Pod string `json:"pod"`
	} `json:"sys"`
	DtTxt string `json:"dt_txt"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, there")
}

func tempSeriesHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://api.openweathermap.org/data/2.5/forecast?lat=42.6979&lon=23.3222&appid=d5a332c2fd770645632f720a59006d58&units=metric")

	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var j ForecastResponse;
	json.Unmarshal(bytes, &j)

	for i, item := range j.List {
		fmt.Println(i, item.Weather[0].Main)
	}

	fmt.Fprint(w, j.List[0].Dt)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/weather/temp_series", tempSeriesHandler)
	http.ListenAndServe(":1337", nil)
}
