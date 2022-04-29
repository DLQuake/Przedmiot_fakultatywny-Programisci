package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Utwórz strukturę o nazwie Secret z jednym polem Ids typu string
//
// 1. Napisz funkcję która wygeneruje losowy ciąg znaków, użyj funkcji rand.Intn() oraz stałej letterBytes
//    funkcja powinna przyjmować jeden parametr typu int jak długi ma być to ciąg znaków
//
// 2. Napisz kolejną funkcję która będzie wywoływała powyższą oraz będzie przyjmowała 2 parametry, liczbę id
//    do wygenerowania oraz ich długość.
//    Zwrócony id z powyższej funkcji powinien być zapisany do jednej zmiennej oddzielony nową linią.
//    Użyj do tego strings.Builder
//
// 3. Do wcześniej utworzonej struktury przypisz funkcję do konkretnego pola.
// Wypisz wynik

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Secret struct {
	Ids string
}

func generateId(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func generateIds(length, count int) string {
	var builder strings.Builder
	for i := 0; i < count; i++ {
		builder.WriteString(generateId(length))
		builder.WriteString("\n")
	}
	return builder.String()
}

func main() {
	secret := Secret{
		Ids: generateIds(5, 10),
	}

	fmt.Println(secret)
}
