package main

import (
	"fmt"
	"os"
	"strconv"
)

// ------------------------------------------------ ---------
// ĆWICZENIE: Sumuj do N
//
// 1. Pobierz dwie liczby z wiersza poleceń: min i max
// 2. Zamień je na liczby całkowite (za pomocą Atoi)
// 3. Używając pętli, zsumuj liczby między min i max
//
// OGRANICZENIA
// 1. Sprawdź również, czy min < max.
//
// WSKAZÓWKA
// Do konwersji liczb możesz użyć `strconv.Atoi`.
//
// OCZEKIWANY WYNIK
// Załóżmy, że użytkownik uruchamia go tak:
//
// go, run main.go 1 10
//
// Oczekiwany wynik:
//
// 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ------------------------------------------------ ---------

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Podaj min oraz max")
		return
	}

	min, err := strconv.Atoi(os.Args[1])
	max, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("To nie są liczby")
	}

	suma := 0
	if min < max {
		for i := min; i <= max; i++ {
			suma += i
		}

		for i := min; i <= max; i++ {
			fmt.Print(i, " + ")
			if i == max - 1 {
				fmt.Print(i+1)
				break
			}
		}
		fmt.Println(" =", suma)
	} else {
		fmt.Println("min musi być mniejsze od max")
	}

}
