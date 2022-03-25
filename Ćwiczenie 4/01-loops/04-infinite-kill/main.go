package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// 1. Utwórz nieskończoną pętlę
// 2. Na każdym kroku pętli wypisz losowy znak.
// 3. Uśpij program przez 250 milisekund.
// 4. Uruchom program i poczekaj kilka sekund
// następnie zabij go za pomocą klawiszy CTRL+C
//
// OGRANICZENIA
// 1. Wydrukuj losowo jeden z tych znaków: \ / | -
// 2. Przed wydrukowaniem znaku wypisz również
// sekwencja ucieczki: \r
//
// "\r/", lub tak: "\r|" i tak dalej...
//
// 3. UWAGA: Jeśli używasz Go Playground, użyj „\f” zamiast „\r”
//
// WSKAZÓWKI
// 1. Użyj `time.Sleep` do uśpienia aplikacji.
// 2. Powinieneś przekazać do niego wartość `time.Duration`.
// 3. Zapoznaj się z dokumentacją Go online dla
// Stała `milisekundy`.
// 4. Podczas drukowania nie używaj znaku nowej linii!
// Zamiast tego użyj Print lub Printf.
//

// OCZEKIWANY WYNIK
// - Program powinien wyświetlać następujące komunikaty w dowolnej kolejności.
// - I pierwszy znak (\, -, / lub |) powinien być losowy
// wyświetlane.
//  \ Please Wait. Processing....
//  - Please Wait. Processing....
//  / Please Wait. Processing....
//  | Please Wait. Processing....
// ------------------------------------------------ ---------
func main() {

	for {
		var c string

		switch rand.Intn(4) {
		case 0:
			c = "\\"
		case 1:
			c = "/"
		case 2:
			c = "|"
		case 3:
			c = "-"
		}
		fmt.Printf("\r%s Please Wait. Processing....", c)
		time.Sleep(time.Microsecond * 250)
	}
}
