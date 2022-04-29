package main

import (
	"fmt"
	"strconv"
)

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Napisz funkcję convert która jako parametr przyjmuje funkcję anonimową, ta funkcja powinna przyjmować int i zwracać string
// do implementacji funkcji użyj strconv.Itoa(i)
//
// Wypisz wynik

func convert(f func(int) string) string {
	return f(42)
}

func main() {
	fmt.Println(convert(func(i int) string {
		return strconv.Itoa(i)
	}))
}