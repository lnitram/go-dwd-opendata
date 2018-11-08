package main

import (
        "fmt"
        "github.com/lnitram/go-dwd-opendata/weather/poi"
)

func main() {
	poi.LoadDB("../../res/stations.json")
	station, weather := poi.GetWeatherByName("helgoland")
	fmt.Println(weather)
	fmt.Println(station)
//        url := "http://opendata.dwd.de/weather/weather_reports/poi/10007-BEOB.csv"
//        weather := poi.GetWeather(url)
//        for _, v := range weather {
//                fmt.Println(v.Headline)
//        }
}

