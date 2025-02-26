package vdwl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadVid(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("FAILED TO DOWNLOAD FILE: %s", resp.Status)
	}

	outFile, err := os.Create(filename)
	if err != nil {
		return err
	} 
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}

	log.Println("Download complete: ", filename)
	return nil
}

