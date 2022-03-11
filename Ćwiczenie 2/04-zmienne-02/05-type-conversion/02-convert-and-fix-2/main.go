package main

// ------------------------------------------------ ---------
// ĆWICZENIE: Konwertuj i napraw nr 2
//
// Napraw kod za pomocą wyrażenia konwersji.
//
// OCZEKIWANY WYNIK
// 10,5
// ------------------------------------------------ ---------

import "fmt"

func main() {
	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)
}
