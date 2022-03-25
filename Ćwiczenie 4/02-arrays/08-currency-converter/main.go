package main

import (
	"fmt"
	"os"
	"strconv"
)

// ------------------------------------------------ ---------
// ĆWICZENIE: Przelicznik walut
//
// W tym ćwiczeniu wyświetlisz przeliczniki walut
// w stosunku do USD.
//
// 1. Zadeklaruj kilka stałych za pomocą iota. Będą kluczami
// tablicy.
//
// 2. Utwórz tablicę zawierającą współczynniki konwersji.
//
// Powinieneś użyć elementów z kluczami i zawartości, którą zadeklarowałeś wcześniej.
//
// 3. Pobierz kwotę USD do przeliczenia z wiersza poleceń.
//
// 4. Obsłuż przypadki błędów w przypadku brakujących lub nieprawidłowych danych wejściowych.
//
// 5. Wydrukuj kursy wymiany.
//
// OCZEKIWANY WYNIK
//   go run main.go
//     Please provide the amount to be converted.
//
//   go run main.go invalid
//     Invalid amount. It should be a number.
//
//   go run main.go 10.5
//     10.50 USD is 9.24 EUR
//     10.50 USD is 8.19 GBP
//     10.50 USD is 1186.71 JPY
//
//   go run main.go 1
//     1.00 USD is 0.88 EUR
//     1.00 USD is 0.78 GBP
//     1.00 USD is 113.02 JPY
// ---------------------------------------------------------

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please provide the amount to be converted")
		return
	}

	usd, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	const (
		eur = iota
		gbp
		jpy
	)

	conv := [3]float64{0.88, 0.78, 113.02}


	fmt.Println(usd,"USD is",conv[eur] * usd,"EUR")
	fmt.Println(usd,"USD is",conv[gbp] * usd,"GBP")
	fmt.Println(usd,"USD is",conv[jpy] * usd,"JPY")
}
