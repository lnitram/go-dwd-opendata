package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	downloadFiles()
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
		fmt.Println("File", filepath, "already exists, skipping download ...")
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
