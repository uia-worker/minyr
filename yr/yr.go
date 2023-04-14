package yr

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// konverterer temperatur fra Celsius til Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func Convert() error {
	// Sjekker om output filen eksisterer
	if _, err := os.Stat("yr/kjevik-temp-fahr-20220318-20230318.csv"); !os.IsNotExist(err) {
		var regenerate string
		fmt.Print("Output filen eksisterer fra før av, ønsker du å regenerere den? (y/n): ")
		fmt.Scanln(&regenerate)
		if regenerate != "y" && regenerate != "Y" {
			fmt.Println("Avbryt uten å generere ny fil.")
			return nil
		}
	}

	// Åpner input csv filen
	inputFile, err := os.Open("yr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		fmt.Println("Feil med åpning av input fil:", err)
	}
	defer inputFile.Close()

	// Lager en ny scanner for å lese input csv filen
	inputScanner := bufio.NewScanner(inputFile)

	// lag en ny csv writer for å skrive til output csv filen
	outputFile, err := os.Create("yr/kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		fmt.Println("Feil under generering av output fil:", err)
	}
	defer outputFile.Close()

	outputWriter := csv.NewWriter(outputFile)
	defer outputWriter.Flush()

	// Skriver ut første linje i input csv filen
	if inputScanner.Scan() {
		firstLine := inputScanner.Text()
		if err = outputWriter.Write(strings.Split(firstLine, ";")); err != nil {
			fmt.Println("Feil under skriving av første linje:", err)
		}
	}

	// Loop gjennom hver linje i input csv filen
	lineNo := 2 // Starter på linje 2 siden linje 1 allerede er skrevet
	for inputScanner.Scan() {
		// Sjekker om linje nummer overstiger 16755 og bryter ut av loopen hvis det gjør det
		if lineNo > 16755 {
			break
		}

		// Splitter linjen i felt
		fields := strings.Split(inputScanner.Text(), ";")

		// Sjekker at feltet har minst 4 elementer
		if len(fields) != 4 {
			fmt.Printf("Error on line %d: Invalid input format.\n", lineNo)
			continue
		}

		// Henter ut siste siffer fra fjerde kolonne
		temperatureField := fields[3]
		if temperatureField == "" {
			fmt.Printf("Feil på linje %d: Tempraturfeilt er tomt.\n", lineNo)
			continue
		}
		temperature, err := strconv.ParseFloat(temperatureField, 64)
		if err != nil {
			fmt.Printf("Feil på linje %d: %v\n", lineNo, err)
			continue
		}
		if math.IsNaN(temperature) {
			fmt.Printf("Feil på linje %d: Tempratur er ikke gyldig float64 verdi.\n", lineNo)
			continue
		}

		lastDigit := temperature

		// Konverterer Celsius til Fahrenheit
		// fahrenheit := conv.CelsiusToFarenheit(lastDigit) //
		fahrenheit := celsiusToFahrenheit(lastDigit)

		// Skriver output til CSV filen
		fields[3] = fmt.Sprintf("%.1f", fahrenheit)
		if err = outputWriter.Write(fields); err != nil {
			fmt.Println("Feilet under skriving av output fil:", err)
			return err
		}

		lineNo++
	}

	dataText := "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Christian Eieland,,,"
	err = outputWriter.Write([]string{dataText})
	if err != nil {
		fmt.Println("Error writing data text to output file:", err)

	}

	return nil
}

func Average(unit string) (float64, error) {
	var filename string
	var tempColumn int
	var delimiter rune

	if unit == "c" {
		filename = "yr/kjevik-temp-celsius-20220318-20230318.csv"
		tempColumn = 3
		delimiter = ';'
	} else if unit == "f" {
		filename = "yr/kjevik-temp-fahr-20220318-20230318.csv"
		tempColumn = 3
		delimiter = ','
	} else {
		return 0, fmt.Errorf("Ugyldig verdi: %s", unit)
	}

	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = delimiter

	var total float64
	var count int

	// Looper gjennom hver linje i CSV filen
	for i := 1; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		if i < 1 || i > 1675 {
			// hopper over linjer utenfor rangen
			continue
		}

		if len(record) <= tempColumn {
			return 0, fmt.Errorf("Ugyldig data i filen %s", filename)
		}

		temp, err := strconv.ParseFloat(record[tempColumn], 64)
		if err != nil {
			return 0, err
		}

		total += temp
		count++

	}

	if count == 0 {
		return 0, fmt.Errorf("Ingen tempratur ble funnet i filen %s", filename)
	}

	return total / float64(count), nil
}
