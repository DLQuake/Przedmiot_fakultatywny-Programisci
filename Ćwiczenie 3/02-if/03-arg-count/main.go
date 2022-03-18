package main

import (
	"fmt"
	"os"
)

// ------------------------------------------------ ---------
// ĆWICZENIE: Liczba argumentów
//
// 1. Pobierz argumenty z wiersza poleceń.
// 2. Wydrukuj oczekiwane wyniki poniżej w zależności od liczby argumentów.
//
// OCZEKIWANY WYNIK
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	argument := os.Args[1:]

	length := len(os.Args[1:])

	if length == 0 {
		fmt.Println("Give me args")
	} else if length == 1 {
		fmt.Println("There is one: ",argument)
	} else if length == 2 {
		fmt.Println("There is two: ",argument)
	} else {
		fmt.Println("There are ",length," arguments")
	}
}
