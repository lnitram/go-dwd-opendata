package poi

import (
        "io/ioutil"
        "log"
        "fmt"
        "net/http"
        "strings"
)

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

