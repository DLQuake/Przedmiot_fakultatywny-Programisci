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
// Utwórz klon unix'owego polecenia grep. Grep to narzędzie wiersza poleceń do
// wyszukiwanie danych w postaci zwykłego tekstu w przeszukiwanym wierszu, które pasują do określonego
//  wzoru
//
// 1. Wprowadź plik shakespeare.txt do swojego programu.
//
// 2. Zaakceptuj argument wiersza poleceń dla wzorca
//
// 3. Wydrukuj tylko te linie, które zawierają ten wzór
//
// 4. Jeśli nie podano wzoru, wydrukuj wszystkie linie
//
//
// OCZEKIWANY WYNIK
//
// go run main.go come < shakespeare.txt
//
//    come night come romeo come thou day in night
//    come gentle night come loving black-browed night
//
// ------------------------------------------------ ---------

func main() {

	file, err := os.Open("shakespeare.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Podaj wzorzec")
		return
	}

	scanner := bufio.NewScanner(file)

	if len(args) == 0 {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		pattern := args[0]
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), pattern) {
				fmt.Println(scanner.Text())
			}
		}
	}

}
