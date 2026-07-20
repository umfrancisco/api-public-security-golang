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

//encore:api public path=/api/:city
func GetData(ctx context.Context, city string) (*CrimeResponse, error) {
	var crimeRecord []Crime
	var err error
	if city == "sao_paulo" {
		crimeRecord, err = ReadCSV("https://raw.githubusercontent.com/umfrancisco/api-public-security-golang/refs/heads/master/data/TaxaDelito-S%C3%A3o%20Paulo_20260704_164249.csv", "São Paulo")	
	}
	if city == "campinas" {
		crimeRecord, err = ReadCSV("https://raw.githubusercontent.com/umfrancisco/api-public-security-golang/refs/heads/master/data/TaxaDelito-Campinas_20260704_194100.csv", "Campinas")
	}
	if city == "guarulhos" {
		crimeRecord, err = ReadCSV("https://raw.githubusercontent.com/umfrancisco/api-public-security-golang/refs/heads/master/data/TaxaDelito-Guarulhos_20260707_161733.csv", "Guarulhos")
	}
	if city == "osasco" {
		crimeRecord, err = ReadCSV("https://raw.githubusercontent.com/umfrancisco/api-public-security-golang/refs/heads/master/data/TaxaDelito-Osasco_20260707_161819.csv", "Osasco")
	}
	if city == "ribeirao_preto" {
		crimeRecord, err = ReadCSV("https://raw.githubusercontent.com/umfrancisco/api-public-security-golang/refs/heads/master/data/TaxaDelito-Ribeir%C3%A3o%20Preto_20260707_161850.csv", "Ribeirão Preto")
	}
	if city == "santo_andre" {
		crimeRecord, err = ReadCSV("https://raw.githubusercontent.com/umfrancisco/api-public-security-golang/refs/heads/master/data/TaxaDelito-Santo%20Andr%C3%A9_20260707_161804.csv", "Santo André")
	}
	if city == "santos" {
		crimeRecord, err = ReadCSV("https://raw.githubusercontent.com/umfrancisco/api-public-security-golang/refs/heads/master/data/TaxaDelito-Santos_20260707_161840.csv", "Santos")
	}
	if city == "sorocaba" {
		crimeRecord, err = ReadCSV("https://raw.githubusercontent.com/umfrancisco/api-public-security-golang/refs/heads/master/data/TaxaDelito-Sorocaba_20260707_161831.csv", "Sorocaba")
	}
	if city == "sao_bernardo" {
		crimeRecord, err = ReadCSV("https://raw.githubusercontent.com/umfrancisco/api-public-security-golang/refs/heads/master/data/TaxaDelito-S%C3%A3o%20Bernardo%20do%20Campo_20260707_161751.csv", "São Bernardo")
	}
	if city == "sao_jose_dos_campos" {
		crimeRecord, err = ReadCSV("https://raw.githubusercontent.com/umfrancisco/api-public-security-golang/refs/heads/master/data/TaxaDelito-S%C3%A3o%20Jos%C3%A9%20dos%20Campos_20260707_161856.csv", "São José dos Campos")
	}
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

