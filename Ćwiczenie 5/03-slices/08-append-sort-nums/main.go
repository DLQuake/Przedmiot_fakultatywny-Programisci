package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

// ------------------------------------------------ ---------
// ĆWICZENIE: dołączanie i sortowanie liczb
//
// Mamy []string, który powinien zawierać liczby.
//
// Twoim zadaniem jest przekonwertowanie []string na wycinek int.
//
// 1. Pobierz liczby z wiersza poleceń
//
// 2. Dołącz je do []int
//
// 3. Sortuj liczby
//
// 4. Wydrukuj je
//
// 5. Obsługuj przypadki błędów
//
//
// OCZEKIWANY WYNIK
//
//  go run main.go
//    provide a few numbers
//
//  go run main.go 4 6 1 3 0 9 2
//    [0 1 2 3 4 6 9]
//
//  go run main.go a b c
//    []
//
//  go run main.go 4 a 1 c d 9
//    [1 4 9]
//
// ---------------------------------------------------------

func main() {
	numbers := []string{}

	if len(os.Args[1:]) == 0 {
		fmt.Println("provide a few numbers")
		return
	}

	for i := 1; i < len(os.Args); i++ {
		numbers = append(numbers, os.Args[i])
	}

	ints := []int{}

	for _, s := range numbers {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		ints = append(ints, i)
	}

	sort.Ints(ints)

	fmt.Printf("%v\n", ints)
}
