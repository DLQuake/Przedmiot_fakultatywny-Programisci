package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE: Iota
// 1. Zainicjuj wiele stałych za pomocą iota.
// 2. Postępuj zgodnie z instrukcjami zawartymi w kodzie.
//
// OCZEKIWANY WYNIK
// 1 2 3
// ------------------------------------------------ ---------

func main() {
	// WSKAZÓWKA: To jest poprawna deklaracja stałej
	// podkreślnik może być użyty zamiast nazwy
	const _ = iota
	//    ^- to tylko nazwa

	// Teraz użyj iota i zainicjuj następujące stałe
	// "automatycznie" odpowiednio na 1, 2 i 3.
	const (
		_ = iota
		Jan
		Feb
		Mar
	)

	// To powinno wypisać: 1 2 3
	// Nie: 0 1 2
	fmt.Println(Jan, Feb, Mar)
}
