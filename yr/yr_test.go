package yr

import (
	"testing"
)

func TestCelsiusToFahrenheitString(t *testing.T) {
     type test struct {
	input string
	want string
     }
     tests := []test{
	     {input: "6", want: "42.8"},
	     {input: "0", want: "32.0"},
     }

     for _, tc := range tests {
	     got, _ := CelsiusToFahrenheitString(tc.input)
	     if !(tc.want == got) {
		     t.Errorf("expected %s, got: %s", tc.want, got)
	     }
     }
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon 
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperatrmaaling i grader celsius
func TestCelsiusToFahrenheitLine(t *testing.T) {
     type test struct {
	input string
	want string
     }
     tests := []test{
	     {input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
	     //{input: "Kjevik;SN39040;18.03.2022 01:50", want: ""},

     }

     for _, tc := range tests {
	     got, _ := CelsiusToFahrenheitLine(tc.input)
	     if !(tc.want == got) {
		     t.Errorf("expected %s, got: %s", tc.want, got)
	     }
     }

	
}
