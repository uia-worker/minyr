package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Åpner inputfilen og leser inn data
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Konverterer temperaturene fra Celsius til Fahrenheit og lagrer resultatene i en ny liste
	var convertedData [][]string
	for i, row := range data {
		if i == 0 { // overskriftsraden skal ikke endres
			convertedData = append(convertedData, row)
			continue
		}
		celsius, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			panic(err)
		}
		fahrenheit := (celsius * 9 / 5) + 32
		convertedRow := []string{row[0], fmt.Sprintf("%.1f", fahrenheit)}
		convertedData = append(convertedData, convertedRow)
	}

	// Skriver resultatene til outputfilen
	outputFile, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	if err := writer.WriteAll(convertedData); err != nil {
		panic(err)
	}

	// Skriver en avslutningsmelding til konsollen
	fmt.Println("Konvertering fullført og lagret i kjevik-temp-fahr-20220318-20230318.csv")
	fmt.Println("Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av STUDENTENS_NAVN")
}
