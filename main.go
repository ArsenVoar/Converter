package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Json struct {
	JS string `json:"JS"`
	CS string `json:"CS"`
	RK int64  `json:"RK"`
}

func vet(src, dn string) error {
	File, err := os.Open(src)
	if err != nil {
		return err
	}
	defer File.Close()

	var ranking []Json
	if err := json.NewDecoder(File).Decode(&ranking); err != nil {
		return err
	}

	outputFile, err := os.Create(dn)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"JS", "CS", "RK"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range ranking {
		var csvRow []string
		csvRow = append(csvRow, r.JS, r.CS, fmt.Sprint(r.RK))
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := vet("json.json", "data.csv"); err != nil {
		log.Fatal(err)
	}
}
