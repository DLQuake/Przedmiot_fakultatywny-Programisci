package main

import (
	"fmt"
	"os"
	"strconv"
)

// ------------------------------------------------ ---------
// ĆWICZENIE: Sortowanie
//
// Twoim celem jest posortowanie podanych liczb z wiersza poleceń.
//
// 1. Pobierz liczby z wiersza poleceń.
//
// 2. Utwórz tablicę i przypisz do niej podane liczby.
//
// 3. Posortuj podane liczby i wydrukuj je.
//
// OGRANICZENIE
// + Można podać maksymalnie 5 liczb
// + Jeśli jeden z argumentów nie jest prawidłową liczbą, pomiń go
//
// WSKAZÓWKI
// + Do sortowania liczb można użyć algorytmu sortowania bąbelkowego.
// Proszę obejrzeć to: https://youtu.be/nmhjrI-aW5o?t=7
//
// OCZEKIWANY WYNIK
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
// ---------------------------------------------------------

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Please give me up to 5 numbers.")
		return
	}

	if len(os.Args[1:]) > 5 {
		fmt.Println("Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	liczba1, _ := strconv.Atoi(os.Args[1])
	liczba2, _ := strconv.Atoi(os.Args[2])
	liczba3, _ := strconv.Atoi(os.Args[3])
	liczba4, _ := strconv.Atoi(os.Args[4])
	liczba5, _ := strconv.Atoi(os.Args[5])

	tablica := [5]int{liczba1,liczba2,liczba3,liczba4,liczba5}

	fmt.Println(tablica)

	for i := 0; i < len(tablica)-1; i++ {
		for j := 0; j < len(tablica)-i-1; j++ {
			if tablica[j] > tablica[j+1] {
				temp := tablica[j]
				tablica[j] = tablica[j+1]
				tablica[j+1] = temp
			}
		}
	}

	fmt.Println(tablica)
}
