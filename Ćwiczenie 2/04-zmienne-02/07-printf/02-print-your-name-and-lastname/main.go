package main

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Wydrukuj swoje imię i nazwisko za pomocą Printf
//
// OCZEKIWANY WYNIK
// Nazywam się X, a moje nazwisko to XY.
//
// Przechowuj specyfikator formatowania (pierwszy argument Printf)
// w zmiennej.
// Następnie przekaż go do printf
// ------------------------------------------------ ---------
import "fmt"

func main() {
	X, Y := "Dominik", "Lewczyński"
	fmt.Printf("Nazywam się %s, a moje nazwisko to %s.", X, X+Y)
}
