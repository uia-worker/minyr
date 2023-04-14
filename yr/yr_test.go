package yr

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strings"
	"testing"
)

// antall linjer i filen er 16756
func TestFileLineCount(t *testing.T) {
	filename := ("kjevik-temp-celsius-20220318-20230318.csv")
	expectedLines := 16756

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Feilet å åpne fil %s: %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("Feilet å skanne fil %s: %v", filename, err)
	}

	if lineCount != expectedLines {
		t.Errorf("uforusett linje antall i fil %s: forventet %d, fikk %d", filename, expectedLines, lineCount)
	}
}

// gitt "Kjevik;SN39040;18.03.2022 01:50;6" ønsker å få (want) "Kjevik;SN39040;18.03.2022 01:50;42,8"
func TestConversion8(t *testing.T) {
	// åpnee csv filen
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Errorf("Feilet å åpne fil: %v", err)
	}
	defer file.Close()

	// Lager en ny csv lesere for å lese csv filen
	reader := csv.NewReader(file)

	// Loop gjennom hver linje i csv filen
	for {
		// Leser en linje fra csv filen
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				t.Errorf("Feilet å lese filen: %v", err)
				return
			}
		}

		// Sjekker om linjen matcher den spesifiserte linjen
		if line[0] == "Kjevik" && line[1] == "SN39040" && line[2] == "18.03.2022 01:50" {
			// Sjekker om temperaturen er riktig konvertert
			want := "42.8"
			got := line[3]
			if got != want {
				t.Errorf("Konvertering feil. Fikk %v, forventer %v", got, want)
			}
			return
		}
	}
	t.Errorf("Linje ikke funnet.")
}

// gitt "Kjevik;SN39040;07.03.2023 18:20;0" ønsker å få (want) "Kjevik;SN39040;07.03.2023 18:20;32"
func TestConversion32(t *testing.T) {
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Errorf("Feilet å åpne fil: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				t.Errorf("Feilet å lese filen: %v", err)
				return
			}
		}

		if line[0] == "Kjevik" && line[1] == "SN39040" && line[2] == "07.03.2023 18:20" {
			want := "32.0"
			got := line[3]
			if got != want {
				t.Errorf("Konvertering feilet. Fikk %v, forventer %v", got, want)
			}
			return
		}
	}

	t.Errorf("Linje ikke funnet i filen.")
}

// gitt "Kjevik;SN39040;08.03.2023 02:20;-11" ønsker å få (want) "Kjevik;SN39040;08.03.2023 02:20;12.2"
func TestConversion2(t *testing.T) {
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Errorf("Feilet å åpne fil: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				t.Errorf("Feilet å lese filen: %v", err)
				return
			}
		}

		if line[0] == "Kjevik" && line[1] == "SN39040" && line[2] == "08.03.2023 02:20" {
			want := "12.2"
			got := line[3]
			if got != want {
				t.Errorf("Konvertering feilet. Fikk %v, forventer %v", got, want)
			}
			return
		}
	}

	t.Errorf("Linje ikke funnet i filen.")
}

/*
gitt "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;" ønsker å få (want)
"Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av
Christian Eieland Ålykkja", hvor Christian er navn på studenten som leverer besvarelsen
*/
func TestLastLineOfFile(t *testing.T) {
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Skanner gjennom filen linje for linje, og holder styr på den siste linjen som ble lest
	var lastLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}

	// Sjekker at den siste linjen inneholder den forventede teksten
	expectedText := "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Christian Eieland Ålykkja"
	if !strings.Contains(lastLine, expectedText) {
		t.Errorf("Siste linje i filen matcher ikke forventet resultat. fikk: %q, forventet tekst: %q", lastLine, expectedText)
	}
}

func TestAverageCelsius(t *testing.T) {
	expected := 8.55897099200191

	// Endrer arbeidskatalogen til katalogen der CSV-filen befinner seg
	err := os.Chdir("..")
	if err != nil {
		t.Fatalf("Feil: %v", err)
	}

	// Endrer arbeidskatalogen tilbake til katalogen der testfilen befinner seg når testen er ferdig
	defer func() {
		err = os.Chdir("yr")
		if err != nil {
			t.Fatalf("Feil: %v", err)
		}
	}()

	avg, err := Average("c")
	if err != nil {
		t.Fatalf("Feil: %v", err)
	}

	if avg != expected {
		t.Fatalf("Gjennomsnittet er %v, men forventet %v", avg, expected)
	}
}
