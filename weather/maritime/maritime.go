package main

import (
	"bufio"
	"compress/bzip2"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	seaWeatherContentUrl = "https://opendata.dwd.de/weather/maritime/content.log.bz2"
)

func main() {
	content, _ := DownloadFile(seaWeatherContentUrl)
	for _, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {
		fmt.Println(line)
	}
}

func DownloadFile(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	br := bufio.NewReader(resp.Body)
	cr := bzip2.NewReader(br)

	bodyBytes, err2 := ioutil.ReadAll(cr)
	bodyString := string(bodyBytes)

	if err2 != nil {
		return "", err2
	}
	return bodyString, nil
}
