package main

import (
	"fmt"
	"os"
	"strings"
)

// ĆWICZENIE: Manipulator ciągów
// 1. Pobierz operację jako pierwszy argument.
// 2. Pobierz ciąg znaków, którym chcesz manipulować, jako drugi argument.
//
// WSKAZÓWKA
// Funkcje manipulacyjne można znaleźć w pakiecie strings  w dokumentacji Go (ToLower, ToUpper, Title).
//
// OCZEKIWANY WYNIK
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------
func main() {

	if len(os.Args[1:]) != 2 {
		fmt.Println("[command] [string]")
		fmt.Println("Available commands: lower, upper and title")
	}

	switch {
	case os.Args[1] == "lower":
		fmt.Println(strings.ToLower(os.Args[2]))
	case os.Args[1] == "upper":
		fmt.Println(strings.ToUpper(os.Args[2]))
	case os.Args[1] == "title":
		fmt.Println(strings.Title(os.Args[2]))
	default:
		fmt.Println("Unknown command: ", os.Args[1])
	}
}
