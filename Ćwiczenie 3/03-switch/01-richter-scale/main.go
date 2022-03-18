package main

import (
	"fmt"
	"os"
	"strconv"
)

// ĆWICZENIE: Skala Richtera
//
// 1. Pobierz wielkość trzęsienia ziemi z wiersza poleceń.
// 2. Wyświetl odpowiedni opis.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive

//
// OCZEKIWANY WYNIK
//  go run main.go
//    Give me the magnitude of the num.
//
//  go run main.go ABC
//    I couldn't get that, sorry.
//
//  go run main.go 0.5
//    0.50 is micro
//
//  go run main.go 2.5
//    2.50 is very minor
//
//  go run main.go 3
//    3.00 is minor
//
//  go run main.go 4.5
//    4.50 is light
//
//  go run main.go 5
//    5.00 is moderate
//
//  go run main.go 6
//    6.00 is strong
//
//  go run main.go 7
//    7.00 is major
//
//  go run main.go 8
//    8.00 is great
//
//  go run main.go 11
//    11.00 is massive
// ---------------------------------------------------------
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the num.")
		return
	}

	num, err := strconv.ParseFloat(os.Args[1], 64)

	if 	err != nil {
		fmt.Println("I couldn't get that, sorry.")
	}
	switch {
	case num < 2.0:
		fmt.Println(num, "is micro")
	case num >= 2.0 && num < 3.0:
		fmt.Println(num, "is very minor")
	case num >= 3.0 && num < 4.0:
		fmt.Println(num, "is minor")
	case num >= 4.0 && num < 5.0:
		fmt.Println(num, "is light")
	case num >= 5.0 && num < 6.0:
		fmt.Println(num, "is moderate")
	case num >= 6.0 && num < 7.0:
		fmt.Println(num, "is strong")
	case num >= 7.0 && num < 8.0:
		fmt.Println(num, "is major")
	case num >= 8.0 && num < 10.0:
		fmt.Println(num, "is great")
	case num >= 10.0:
		fmt.Println(num, "is massive")
	}
}
