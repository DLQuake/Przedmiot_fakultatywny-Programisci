package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// 1. Zainicjuj stałe za pomocą iota.
// 2. Powinieneś znaleźć poprawną formułę dla jota.
//
// OGRANICZENIA
// 1. Usuń wartości inicjalne ze wszystkich stałych.
// 2. Następnie użyj iota raz do zainicjowania pierwszej stałej
//
// OCZEKIWANY WYNIK
// 9 10 11
// ------------------------------------------------ ---------
func main() {
	// const (
	// 	Nov = 11
	// 	Oct = 10
	// 	Sep = 9
	// )

	const (
		Sep = iota + 9
		Oct
		Nov
	)



	fmt.Println(Sep, Oct, Nov)
}
