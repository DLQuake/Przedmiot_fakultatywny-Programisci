package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Utwórz program, który drukuje wszystkie i niepowtarzalne słowa
// ze strumienia wejściowego.
//
// 1. Wprowadź plik shakespeare.txt do swojego programu.
//
// 2. Zeskanuj dane wejściowe za pomocą funkcji bufio.NewScanner.
//
// 3. Skonfiguruj skaner do skanowania w poszukiwaniu słów.
//
// 4. Policz unikalne słowa za pomocą mapy.
//
// 5. Wydrukuj wszystkie i niepowtarzalne słowa.
//
//
// OCZEKIWANY WYNIK
//
// There are 99 words, 70 of them are unique.
//
// ------------------------------------------------ ---------

func main() {

	file, err := os.Open("shakespeare.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	words := make(map[string]bool)

	allwords := 1
	unique := 0

	for scanner.Scan() {
		allwords++
		if _, ok := words[scanner.Text()]; !ok {
			unique++
			words[scanner.Text()] = true
		}
	}

	fmt.Printf("There are %d words, %d of them are unique.\n", allwords, unique)
}