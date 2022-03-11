package main

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// 1. Zmień wartość zmiennej `color` na "dark green"
//
// 2. Nie przypisuj "dark green" bezpośrednio do `color`
//
// Zamiast tego użyj ponownie zmiennej `color`
// podczas przypisywania do `color`
//
// 3. Wydrukuj wartość tej zmiennej
//
// OCZEKIWANY WYNIK
//  dark green
// ------------------------------------------------ ---------

import "fmt"

func main() {
	color := "green"

	color = "dark "+ color

	fmt.Println(color)
}
