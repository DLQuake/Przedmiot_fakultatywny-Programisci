package main

// ------------------------------------------------ ---------
// ĆWICZENIE: Multi przypisanie
//
// 1. Przypisz poprawne wartości do zmiennych
// aby dopasować do OCZEKIWANEGO WYJŚCIA
//
// 2. Wydrukuj zmienne
//
// WSKAZÓWKA
// Użyj wielu wywołań Println, aby wydrukować każde zdanie.
//
// OCZEKIWANY WYNIK
//  Air is good on Mars
//  It's true
//  It is 19.5 degrees
// ------------------------------------------------ ---------

import "fmt"
func main() {

	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Air is good on Mars", true, 19.5

	fmt.Println(planet)
	fmt.Println("It's ",isTrue)
	fmt.Println("It is ",temp," degrees")
}
