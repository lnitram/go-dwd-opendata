package main

import (
	"fmt"
	"github.com/lnitram/go-dwd-opendata/weather/poi"
)

func main() {
	poi.LoadDB("../../res/stations.json")
	_, weather := poi.GetWeatherByName("Hamburg")
	fmt.Println(".")
	if weather != nil {
	}
	//fmt.Println(weather)
}
