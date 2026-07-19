package app

import (
	"context"
	"encoding/csv"
	"net/http"
	"fmt"
	"strconv"
)

type Crime struct {
	ID int `json:"id"`
	City string `json:"city"`
	Year int `json:"year"`
	Homicides float64 `json:"homicides"`
	Thefts float64 `json:"thefts"`
	Robberies float64 `json:"robberies"`
}

type CrimeResponse struct {
	Content []Crime
}

//encore:api public path=/data
func GetData(ctx context.Context) (*CrimeResponse, error) {
	crimeRecord, err := ReadCSV("https://raw.githubusercontent.com/umfrancisco/api-public-security-golang/refs/heads/master/data/TaxaDelito-S%C3%A3o%20Paulo_20260704_164249.csv", "São Paulo")
	crimeResponse := &CrimeResponse{Content: crimeRecord}
	if err != nil {
		return nil, err
	}
	return crimeResponse, nil
}

func ReadCSV(url string, city string) ([]Crime, error) {
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch CSV: %w", err)
	}
	defer resp.Body.Close()
	
	reader := csv.NewReader(resp.Body)
	
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	
	var records []Crime
	
	for i, row := range rows {
		if i == 0 {
			continue
		}
		
		year, _ := strconv.Atoi(row[0])
		homicides, _ := strconv.ParseFloat(row[1], 64)
		thefts, _ := strconv.ParseFloat(row[2], 64)
		robberies, _ := strconv.ParseFloat(row[3], 64)
		
		records = append(records, Crime{
			ID: i,
			City: city,
			Year: year,
			Homicides: homicides,
			Thefts: thefts,
			Robberies: robberies,
		})
	}
	
	return records, nil
}

