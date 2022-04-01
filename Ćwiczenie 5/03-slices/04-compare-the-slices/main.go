package main

import (
	"fmt"
	"sort"
	"strings"
)

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// 1. Podziel ciąg tekstowy na kawałki
//
// 2. Posortuj wycinki
//
// 3. Porównaj, czy wycinki są równe, czy nie
//
//
// OCZEKIWANY WYNIK
//
//   They are equal.
//
//
// WSKAZÓWKI
//
// 1. funkcja strings.Split  dzieli ciąg i zwraca wycinek typu string
// 2. Porównanie wycinków: Najpierw sprawdź, czy mają taką samą długość, dopiero potem je porównaj.

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	stringsA := strings.Split(namesA, ", ")
	// fmt.Println(stringsA)
	// fmt.Println(namesB)
	fmt.Println(len(stringsA))
	fmt.Println(len(namesB))

	sort.Strings(stringsA)
	sort.Strings(namesB)

	if len(stringsA) == len(namesB) {
		for i := 0; i < len(stringsA); i++ {
			if stringsA[i] != namesB[i] {
				return
			}
		}
		fmt.Println("They are equal")
	}
}