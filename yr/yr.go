package yr

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	"github.com/uia-worker/misc/conv"
)

func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon 
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperaturaaling i grader celsius
func CelsiusToFahrenheitLine(line string) (string, error) {
	elementsInLine := strings.Split(line, ";")
	var err error
	if (len(elementsInLine) == 4) {
		elementsInLine[3], err = CelsiusToFahrenheitString(elementsInLine[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(elementsInLine, ";"), nil
}

