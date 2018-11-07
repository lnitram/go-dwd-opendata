package main

import (
        "fmt"
        "github.com/lnitram/go-dwd-opendata/weather/poi"
)

func main() {
        url := "http://opendata.dwd.de/weather/weather_reports/poi/10007-BEOB.csv"
        weather := poi.GetWeather(url)
        for _, v := range weather {
                fmt.Println(v)
        }
}

