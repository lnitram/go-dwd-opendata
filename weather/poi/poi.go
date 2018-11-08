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

var poiDB []Station

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

type WeatherData struct {
	Headline    string
	Description string
	Unit        string
	Value       string
}

func getPresentWeather(code int) string {
	switch code {
	case 1:
		return "wolkenlos"
	case 2:
		return "heiter"
	case 3:
		return "bewoelkt"
	case 4:
		return "bedeckt"
	case 5:
		return "Nebel"
	case 6:
		return "gefrierender Nebel"
	case 7:
		return "leichter Regen"
	case 8:
		return "Regen"
	case 9:
		return "kraeftiger Regen"
	case 10:
		return "gefrierender Regen"
	case 11:
		return "kraeftiger gefrierender Regen"
	case 12:
		return "Schneeregen"
	case 13:
		return "kraeftiger Schneeregen"
	case 14:
		return "leichter Schneefall"
	case 15:
		return "Schneefall"
	case 16:
		return "kraeftiger Schneefall"
	case 17:
		return "Eiskoerner"
	case 18:
		return "Regenschauer"
	case 19:
		return "kraeftiger Regenschauer"
	case 20:
		return "Schneeregenschauer"
	case 21:
		return "kraeftiger Schneeregenschauer"
	case 22:
		return "Schneeschauer"
	case 23:
		return "kraeftiger Schneeschauer"
	case 24:
		return "Graupelschauer"
	case 25:
		return "kraeftiger Graupelschauer"
	case 26:
		return "Gewitter ohne Niederschlag"
	case 27:
		return "Gewitter"
	case 28:
		return "kraeftiges Gewitter"
	case 29:
		return "Gewitter mit Hagel"
	case 30:
		return "kraeftiges Gewitter mit Hagel"
	case 31:
		return "Boen"
	default:
		return "---"
	}
}

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

func (w WeatherData) String() string {
	return fmt.Sprintf("%v (%v): %v %v", w.Headline, w.Description, w.Value, w.Unit)
}
