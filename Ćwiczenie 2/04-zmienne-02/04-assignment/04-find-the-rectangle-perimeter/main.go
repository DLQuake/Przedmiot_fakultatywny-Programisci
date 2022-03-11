package main

// ĆWICZENIE: Obwód prostokąta
//
// 1. Oblicz obwód prostokąta
// Jego szerokość to 5
// Jego wysokość to 6
//
// 2. Przypisz wynik do zmiennej `perimeter`
//
// 3. Wydrukuj zmienną `perimeter`
//
// WSKAZÓWKA
// Formuła = 2 razy szerokość i wysokość
//
// OCZEKIWANY WYNIK
// 22

import "fmt"

func main() {

	var (
		perimeter        int
		width, height = 5, 6
	)

	perimeter = (width + height)*2

	fmt.Println(perimeter)
}
