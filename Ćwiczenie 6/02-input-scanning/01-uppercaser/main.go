package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Użyj skanera, aby przekonwertować wiersze na wielkie litery, i
// wydrukuj je.
//
// 1. Odczytaj plik shakespeare.txt do swojego programu.
//
// 2. Zeskanuj dane wejściowe za pomocą bufio.NewScanner zobacz dokumentację.
//
// 3. Wydrukuj każdy wiersz.
//

func main() {
    file, err := os.Open("shakespeare.txt")
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()

    scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
        fmt.Println(strings.ToUpper(scanner.Text()))
    }
}
