package main

import (
	"fmt"
	"os"
	"strconv"
)

// ------------------------------------------------ ---------
// ĆWICZENIE: Nieparzyste lub parzyste
//
// 1. Pobierz liczbę z wiersza poleceń.
// 2. Sprawdź, czy liczba jest nieparzysta, parzysta i podzielna przez 8.
//
// OGRANICZENIE
// Obsłuż błąd. Jeśli argument nie jest prawidłowym numerem, daj znać użytkownikowi.
// OCZEKIWANY WYNIK
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------
func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Pick a number")
	} else {
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(os.Args[1], "is not a number")
		} else if num%2 == 1 {
			fmt.Println(num, "is an odd number")
		} else if num%2 == 0 && num%8 == 0 {
			fmt.Println(num, "is an even number and it's divisible by 8")
		} else {
			fmt.Println(num, "is an even number")
		}
	}
}
