package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE: Iota
//
// Użyj iota, aby zainicjować stałe pór roku
//
// WSKAZÓWKA
// Możesz zmienić kolejność stałych.
//
// OCZEKIWANY WYNIK
// 12 3 6 9
// ------------------------------------------------ ---------

func main() {
	// UWAGA: Powinieneś usunąć wszystkie inicjatory poniżej
	// Następnie użyj iota, aby to naprawić.
	// const (
	// 	Winter = 12
	// 	Spring = 3
	// 	Summer = 6
	// 	Fall   = 9
	// )

	const (
		Spring  = (iota + 1) * 3
		Summer
		Fall
		Winter
	)
	fmt.Println(Winter, Spring, Summer, Fall)
}
