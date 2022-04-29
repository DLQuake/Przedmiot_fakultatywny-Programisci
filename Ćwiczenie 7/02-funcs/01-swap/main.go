package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Napisz funkcję która zamieni wartości zmiennych poprzez przekazanie do funkcji adresów zmiennych
//
// 1. Zadeklaruj dwie zmienne typu float.
//
// 2. Zadeklaruj funkcję która zamieni wartości zmiennych po przez ich adres
//
// 3. Przekaż do funkcji adresy zmiennych.
//
// 4. Wydrukuj zmienne

func main() {
	var a float64 = 1.1
	var b float64 = 2.2

	swap(&a, &b)

	fmt.Println(a)
	fmt.Println(b)

}

func swap(a *float64, b *float64) {
	var temp float64
	temp = *a
	*a = *b
	*b = temp
}
