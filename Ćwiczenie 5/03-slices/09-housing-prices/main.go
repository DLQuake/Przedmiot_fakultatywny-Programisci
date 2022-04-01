package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ----------------------------------------------------------
// ĆWICZENIE:
//
// Otrzymaliśmy ceny mieszkań. Twoim zadaniem jest załadowanie danych do
// odpowiednio zadeklarowanych wycinków, a następnie wydrukowanie ich na ekranie.
//
// 1. Sprawdź oczekiwany wynik

// 2. Sprawdź poniższy kod
//
// 1. nagłówek: przechowuje nagłówki kolumn
// 2. dane: przechowuje rzeczywiste dane związane z każdą kolumną
// 3. separator: użyjesz go do oddzielenia danych
//
//
// 3. Przeanalizuj dane
//
// 1. Rozdziel go na wiersze za pomocą znaku nowej linii ("\n")
//
// 2. Dla każdego wiersza oddziel go za pomocą separatora (",")
//
//
// 4. Załaduj dane do oddzielnych wycinków
//
// 1. Załaduj lokalizacje do []string
// 2. Załaduj powierzchnię do []int
// 3. Załaduj liczbaSypialni do []int
// 4. Załaduj liczbaLazienek do []int
// 5. Załaduj cena do []int
//
//
// 5. Wydrukuj nagłówek
//
// 1. Oddziel to za pomocą separatora
//
// 2. Wydrukuj każdą kolumnę o szerokości 15 znaków ("%-15s")
//
//
// 6. Wydrukuj wiersze z wycinków, które utworzyłeś, wiersz po wierszu
//
//
// OCZEKIWANY WYNIK
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`
		separator = ","
	)

	var (
		location []string
		size     []int
		beds     []int
		baths    []int
		price    []int
	)

	for _, line := range strings.Split(data, "\n") {
		columns := strings.Split(line, separator)

		location = append(location, columns[0])
		size = append(size, atoi(columns[1]))
		beds = append(beds, atoi(columns[2]))
		baths = append(baths, atoi(columns[3]))
		price = append(price, atoi(columns[4]))
	}

	headerColumns := strings.Split(header, separator)
	fmt.Printf("%-15v", headerColumns[0])
	fmt.Printf("%-15v", headerColumns[1])
	fmt.Printf("%-15v", headerColumns[2])
	fmt.Printf("%-15v", headerColumns[3])
	fmt.Printf("%-15v\n", headerColumns[4])

	fmt.Printf("===========================================================================\n")


	for i := 0; i < len(location); i++ {
		fmt.Printf("%-15v", location[i])
		fmt.Printf("%-15v", size[i])
		fmt.Printf("%-15v", beds[i])
		fmt.Printf("%-15v", baths[i])
		fmt.Printf("%-15v\n", price[i])
	}


}

func atoi (s string) int {
	i, _ := strconv.Atoi(s)
	return i
}