package main

// ------------------------------------------------ ---------
// ĆWICZENIE: Wydrukuj temperaturę
//
// Wydrukuj aktualną temperaturę w Twojej okolicy za pomocą Printf
//
// Nie używaj czasownika %v
// Wyjście "nie powinno" wyglądać tak 29.50000 a 29,5
//
// OCZEKIWANY WYNIK
// Temperatura wynosi 29,5 stopnia.
// ------------------------------------------------ ---------

import "fmt"

func main() {
	fmt.Printf("Temperatura wynosi %.1f stopnia.", 29.5)
}
