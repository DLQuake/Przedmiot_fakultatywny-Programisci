package main

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Oblicz ile minut jest w dwóch tygodniach.
//
//  KROKI:
// 1. Zadeklaruj stałą `minsPerDay` i zainicjuj ją
// do liczby minut w ciągu dnia
//
// 2. Zadeklaruj stałą `weekDays` i zainicjuj ją
// do 7.
//
// 3. Użyj printf, aby wydrukować całkowitą liczbę minut
//     w ciągu dwóch tygodni.
//
// OCZEKIWANY WYNIK
// There are 20160 minutes in 2 weeks.
// ------------------------------------------------ ---------
import "fmt"

func main() {
	const minsPerDay int = 1440
	const weekDays int = 7

	fmt.Printf("There are %d minutes in %d weeks.",minsPerDay*2*weekDays, weekDays*2)
}
