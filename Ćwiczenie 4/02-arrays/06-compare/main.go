package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE: Porównaj tablice
//
// 1. Odkomentuj kod
//
// 2. Napraw problemy, aby tablice stały się porównywalne
//
// OCZEKIWANY WYNIK
//  prawda
//  prawda
//  fałszywe
// ------------------------------------------------ ---------

func main() {
	// week := [...]string{"Monday", "Tuesday"}
	// wend := [4]string{"Saturday", "Sunday"}

	// fmt.Println(week != wend)

	// evens := [...]int{2, 4, 6, 8, 10}
	// evens2 := [...]int32{2, 4, 6, 8, 10}

	// fmt.Println(evens == evens2)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.
	// image := [5]byte{'h', 'i'}
	// var data [5]byte

	// fmt.Println(data == image)

	week := [...]string{"Monday", "Tuesday"}
	wend := [2]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	evens := [...]int32{2, 4, 6, 8, 10}
	evens2 := [...]int32{2, 4, 6, 8, 10}

	fmt.Println(evens == evens2)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.
	image := [5]uint8{'h', 'i'}
	var data [5]byte

	fmt.Println(data == image)
}
