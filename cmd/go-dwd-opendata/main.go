package main

import (
	"fmt"
	"github.com/lnitram/go-dwd-opendata/weather/poi"
)

func main() {
	poi.GenerateStationFile()


	poi.LoadDB("../../res/stations.json")
	station, weather := poi.GetWeatherByName("Hamburg")
	fmt.Println("Station              :", station.Name)
	fmt.Println("Date                 :", weather["surface observations"].Value)
	fmt.Println("Time                 :", weather["Parameter description"].Value)
	fmt.Println("Pressure             :", weather["pressure_reduced_to_mean_sea_level"].Value)
	fmt.Println("Wind Direction (°)   :", weather["mean_wind_direction_during_last_10 min_at_10_meters_above_ground"].Value)
	fmt.Println("Wind Speed (km/h)    :", weather["mean_wind_speed_during last_10_min_at_10_meters_above_ground"].Value)
	fmt.Println("Max Wind Speed (km/h):", weather["maximum_wind_speed_last_hour"].Value)
	fmt.Println("Temp (°C)            :", weather["dry_bulb_temperature_at_2_meter_above_ground"].Value)
	fmt.Println("Visibility (km)      :", weather["horizontal_visibility"].Value)
	fmt.Println("Cloud coverage (%)   :", weather["cloud_cover_total"].Value)
	fmt.Println("Humidity             :", weather["relative_humidity"].Value)
	fmt.Println("Cloud height         :", weather["height_of_base_of_lowest_cloud_above_station"].Value)
}

/*
maximum_wind_speed_as_10_minutes_mean_during_last_hour Maximalwind (letzte Stunde) km/h 18
past_weather_2 vergangenes Wetter 2 CODE_TABLE ---
past_weather_1 vergangenes Wetter 1 CODE_TABLE ---
present_weather aktuelles Wetter CODE_TABLE 1

*/
