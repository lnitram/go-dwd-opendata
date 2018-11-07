package main

import (
	"encoding/json"
	"fmt"
	"github.com/extrame/xls"
	"io"
	"net/http"
	"os"
	"strings"
	"strconv"
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
	generateHAJson()
	generateNAJson()
}

func checkFormat(row []string, format string) bool {
	expected := strings.Split(format,",")

	if len(expected) != len(row) {
		fmt.Println("unknown format: number of fields does not match")
		return false
	}

	for i:= 0; i < len(expected); i++ {
		if expected[i] != row[i] {
			fmt.Println("unknown format: field does not match:",i,expected[i],row[i])
			return false
		}
	}
	return true
}

func generateHAJson() {
	haFormat := "ID,Stations-Name,WMO-Kennung,BG,BM,BS,LG,LM,LS,GEOGR_BREITE,GEOGR_LAENGE,STATIONSHOEHE,Betreiber,Melde-Grp,Country"
	xls, err := xls.Open("./ha.xls", "utf-8")
	if err != nil {
		fmt.Println(err)
		return
	}

	all := xls.ReadAllCells(1000000)
	headlines := all[0]
	if !checkFormat(headlines, haFormat) {
		panic("Unknown format")
	}
	for i := 1; i < len(all); i++ {
		row := all[i]
		id, _ := strconv.Atoi(row[0])
		name := row[1]
		kennung := row[2]
		lat, _ := strconv.ParseFloat(row[9], 64)
		lon, _ := strconv.ParseFloat(row[10], 64)
		height, _ := strconv.ParseFloat(row[11], 64)
		owner := row[12]
		country := row[14]
		s := Station{id, name, kennung, lat, lon, height, owner, country}
		b, _ := json.Marshal(s)
		fmt.Println(string(b))
	}
}

func generateNAJson() {
	naFormat := "STATIONSKENNUNG,STATIONSNAME,STATIONS_ID,MaxvonGERAETETYP_NAME,MinvonVON_DATUM,GEOGR_BREITE,GEOGR_LAENGE,STATIONSHOEHE,Niederschlag 1 Min,Schnee manuell,Wind 10 Min,Temperatur und Feuchte 2 m 10 Min,Sonne 10 Min,Erdbodentemperaturen Standard 10 Min,HEADING_BUFR1,HEADING_BUFR2,HEADING_BUFR3,HEADING_BUFR4,HEADING_BUFR5,HEADING_BUFR6,HEADING_BUFR7"
	xls, err := xls.Open("./na.xls", "utf-8")
	if err != nil {
		fmt.Println(err)
		return
	}

	all := xls.ReadAllCells(1000000)
	headlines := all[0]
	if !checkFormat(headlines, naFormat) {
		panic("Unknown format")
	}

	for i := 1; i < len(all); i++ {
		row := all[i]
		id, _ := strconv.Atoi(row[2])
		name := row[1]
		kennung := row[0]
		lat, _ := strconv.ParseFloat(row[5], 64)
		lon, _ := strconv.ParseFloat(row[6], 64)
		height, _ := strconv.ParseFloat(row[7], 64)
		owner := ""
		country := ""
		s := Station{id, name, kennung, lat, lon, height, owner, country}
		b,_ := json.Marshal(s)
		fmt.Println(string(b))
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
