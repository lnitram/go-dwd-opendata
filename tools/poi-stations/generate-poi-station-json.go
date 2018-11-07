package main

/*
	Reads the station list from the 2 excel files and generates json to be used by other packages
*/

import (
	"encoding/json"
	"fmt"
	"github.com/extrame/xls"
	"io"
	"net/http"
	"os"
	"strconv"
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

func main() {
	downloadFiles()
	os.Remove("./stations.json")
	generateJson("./ha.xls", "HA")
	generateJson("./na.xls", "NA")
}

func checkFormat(row []string, format string) bool {
	expected := strings.Split(format, ",")

	if len(expected) != len(row) {
		fmt.Println("unknown format: number of fields does not match")
		return false
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != row[i] {
			fmt.Println("unknown format: field does not match:", i, expected[i], row[i])
			return false
		}
	}
	return true
}

func getStation(row []string, format string) Station {
	if format == "HA" {
		return getHAStation(row)
	} else if format == "NA" {
		return getNAStation(row)
	}
	return Station{}
}

func getHAStation(row []string) Station {
	id, _ := strconv.Atoi(row[0])
	name := row[1]
	kennung := row[2]
	lat, _ := strconv.ParseFloat(row[9], 64)
	lon, _ := strconv.ParseFloat(row[10], 64)
	height, _ := strconv.ParseFloat(row[11], 64)
	owner := row[12]
	country := row[14]
	return Station{id, name, kennung, lat, lon, height, owner, country}
}

func getNAStation(row []string) Station {
	id, _ := strconv.Atoi(row[2])
	name := row[1]
	kennung := row[0]
	lat, _ := strconv.ParseFloat(row[5], 64)
	lon, _ := strconv.ParseFloat(row[6], 64)
	height, _ := strconv.ParseFloat(row[7], 64)
	owner := ""
	country := ""
	return Station{id, name, kennung, lat, lon, height, owner, country}

}

func generateJson(filename string, format string) {
	f, _ := os.OpenFile("./stations.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	header := ""
	if format == "HA" {
		header = "ID,Stations-Name,WMO-Kennung,BG,BM,BS,LG,LM,LS,GEOGR_BREITE,GEOGR_LAENGE,STATIONSHOEHE,Betreiber,Melde-Grp,Country"
	} else if format == "NA" {
		header = "STATIONSKENNUNG,STATIONSNAME,STATIONS_ID,MaxvonGERAETETYP_NAME,MinvonVON_DATUM,GEOGR_BREITE,GEOGR_LAENGE,STATIONSHOEHE,Niederschlag 1 Min,Schnee manuell,Wind 10 Min,Temperatur und Feuchte 2 m 10 Min,Sonne 10 Min,Erdbodentemperaturen Standard 10 Min,HEADING_BUFR1,HEADING_BUFR2,HEADING_BUFR3,HEADING_BUFR4,HEADING_BUFR5,HEADING_BUFR6,HEADING_BUFR7"
	}

	xls, err := xls.Open(filename, "utf-8")
	if err != nil {
		fmt.Println(err)
		return
	}

	all := xls.ReadAllCells(1000000)
	headlines := all[0]
	if !checkFormat(headlines, header) {
		panic("Unknown format")
	}

	for i := 1; i < len(all); i++ {
		row := all[i]
		s := getStation(row, format)
		b, _ := json.Marshal(s)
		f.WriteString(string(b) + "\n")
	}
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func downloadFiles() {
	fileUrlHA := "https://www.dwd.de/DE/leistungen/opendata/help/stationen/ha_messnetz.xls?__blob=publicationFile&v=1"
	fileUrlNA := "https://www.dwd.de/DE/leistungen/opendata/help/stationen/na_messnetz.xls?__blob=publicationFile&v=9"
	err := DownloadFile("ha.xls", fileUrlHA)
	if err != nil {
		fmt.Println("Error while downloading", fileUrlHA)
	}
	err = DownloadFile("na.xls", fileUrlNA)
	if err != nil {
		fmt.Println("Error while downloading", fileUrlNA)
	}
}

func DownloadFile(filepath string, url string) error {
	if exists, _ := fileExists(filepath); exists {
		return nil
	}
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
