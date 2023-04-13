package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Lordkissa97/minyr/yr"
)

func main() {
	// Venter på at brukeren skal skrive inn "minyr"
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Skriv inn 'minyr' for å starte programmet: ")
	text, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(text)) != "minyr" {
		fmt.Println("Ugyldig verdi.")
		return
	}

	// Viser brukeren en meny med valg
	fmt.Println("Valg:")
	fmt.Println("  - 'convert' for å konvertere tempraturen fra Celsius til Fahrenheit")
	fmt.Println("  - 'average' for å begregne gjennomsnitt tempratur for perioden")
	fmt.Println("Skriv 'q' eller 'quit' for å avbryte.")
	for {
		fmt.Print("Velg: ")
		option, _ := reader.ReadString('\n')
		option = strings.ToLower(strings.TrimSpace(option))

		if option == "convert" {
			err := yr.Convert()
			if err != nil {
				fmt.Println("Feil med begregning av gjennomsnitt tempratur:", err)
				return
			}
			fmt.Println("Gjennomsnitt begregning fullført.")
			break
		}

		if option == "average" {
			fmt.Print("Velg enhet for begregning ('c' for Celsius eller 'f' for Fahrenheit): ")
			unit, _ := reader.ReadString('\n')
			unit = strings.ToLower(strings.TrimSpace(unit))

			avg, err := yr.Average(unit)
			if err != nil {
				fmt.Println("Feil under kalkulasjon:", err)
				return
			}
			fmt.Printf("Gjennomsnittlig tempratur: %.2f %s\n", avg, unit)
			break
		}

		if option == "q" || option == "quit" {
			fmt.Println("Avbryt program.")
			return
		}

		fmt.Println("Ugylid verdi, prøv igjen.")
	}
}
