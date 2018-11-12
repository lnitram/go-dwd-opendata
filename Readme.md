Playing with Go to get some weatherdata from the DWD opendata server. Not much useful so far. And very bad error handlin ;-)


Dependencies:
```
go get github.com/extrame/xls
```

Packages:
```
weather/weather_reports/poi
- GenerateStationFile()
    - Reads both excel files with station information from DWD server
    - Stores them as ha.xls and na.xls in current directory
    - Reads both excel files and generates JSON with all stations
    - Example usage in cmd/generate-poi-stations/main.go

- LoadDB(station-file)
    - Loads the station db from the stations file to memory

- FindStationByName(name)
    - Searches the station file for a name (substring search). First match wins

- GetWeatherByName(name)
    - Downloads the weather data for the station that first matches the name and returns a map with the newest data


- Example usage in cmd/poi-weather/main.go


```
