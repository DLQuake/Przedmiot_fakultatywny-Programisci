package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
// Napraw kod.
//
// WSKAZÓWKI
// maksymalnie int8 może wynosić 127
// maksimum int16 może wynosić 32767
//
// OCZEKIWANY WYNIK
// 1127
// ------------------------------------------------ ---------

func main() {
	min := int8(127)
	max := int16(1000)
	fmt.Println(int16(min) + max)
}
