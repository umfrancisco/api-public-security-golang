package app

import (
	"context"
	"encoding/csv"
	"os"
	"strconv"
)

type Crime struct {
	ID string `json:"id"`
	City string `json:"city"`
	Year int `json:"year"`
	Homicides float64 `json:"homicides"`
	Thefts float64 `json:"thefts"`
	Robberies float64 `json:"robberies"`
}

type CrimeRecord struct {
	City string `json:"city"`
	Year int `json:"year"`
	Value float64 `json:"value"`
}

type CrimeResponse struct {
	Content []Crime
}

//encore:api public path=/data
func GetData(ctx context.Context) (*CrimeResponse, error) {
	crimeRecord, err := ReadCSV("data/TaxaDelito-São Paulo_20260704_164249.csv", "São Paulo")
	crimeResponse := &CrimeResponse{Content: crimeRecord}
	if err != nil {
		return nil, err
	}
	return crimeResponse, nil
}

func ReadCSV(path string, city string) ([]Crime, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	
	defer file.Close()
	
	reader := csv.NewReader(file)
	
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	
	var records []Crime
	
	for i, row := range rows {
		if i == 0 {
			continue
		}
		
		year, _ := strconv.Atoi(row[1])
		homicides, _ := strconv.ParseFloat(row[2], 64)
		thefts, _ := strconv.ParseFloat(row[3], 64)
		robberies, _ := strconv.ParseFloat(row[4], 64)
		
		records = append(records, Crime{
			City: city,
			Year: year,
			Homicides: homicides,
			Thefts: thefts,
			Robberies: robberies,
		})
	}
	
	return records, nil
}

