package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE: Multi przypisanie
//
// 1. Przypisz wartośc "go" do zmiennej `lang`
// i przypisz 2 do zmiennej version
// używając instrukcji wielokrotnego przypisania
//
// 2. Wydrukuj zmienne
//
// OCZEKIWANY WYNIK
//  go version 2
// ------------------------------------------------ ---------

func main() {

	var (
		lang    string
		version int
	)

	lang, version = "go", 2

	fmt.Println(lang, "version", version)
}
