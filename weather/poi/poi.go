package poi

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Station struct {
	ID      int
	Name    string
	Kennung string
	Lat     float64
	Lon     float64
	Height  float64
	Owner   string
	Country string
}

var poiDB []Station

func FindStationByName(name string) Station {
	for _, v := range poiDB {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(name)) {
			return v
		}
	}
	return Station{}
}

func GetWeatherByName(name string) (Station, map[string]WeatherData) {
	station := FindStationByName(name)
	url := "http://opendata.dwd.de/weather/weather_reports/poi/" + station.Kennung + "-BEOB.csv"
	return station, GetWeather(url)
}

func LoadDB(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		res := Station{}
		json.Unmarshal([]byte(row), &res)
		poiDB = append(poiDB, res)
	}
}

func downloadTextFile(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error:", err)
		return ""
	}

	if resp.StatusCode != 200 {
		log.Println("Download failed:", url, resp.StatusCode)
		return ""
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error:", err)
		return ""
	}
	return string(contents)
}

func GetWeather(url string) map[string]WeatherData {
	csv := downloadTextFile(url)
	lines := strings.Split(csv, "\n")
	if len(lines) < 3 {
		log.Println("Invalid csv file", url)
		return nil
	}
	weather := make(map[string]WeatherData)
	headlines := strings.Split(lines[0], ";")
	units := strings.Split(lines[1], ";")
	descriptions := strings.Split(lines[2], ";")
	values := strings.Split(lines[3], ";")
	for i := 0; i < len(headlines); i++ {
		weather[headlines[i]] = WeatherData{headlines[i], descriptions[i], units[i], values[i]}
	}
	return weather
}

type WeatherData struct {
	Headline    string
	Description string
	Unit        string
	Value       string
}

func (w WeatherData) String() string {
	return fmt.Sprintf("%v (%v): %v %v", w.Headline, w.Description, w.Value, w.Unit)
}
