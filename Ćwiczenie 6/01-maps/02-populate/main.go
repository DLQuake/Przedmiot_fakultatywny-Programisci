package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Dodaj elementy do map, które zadeklarowałeś w
// pierwszym ćwiczeniu i wypróbuj je, szukając po kluczach.
//
// Po wykonaniu ćwiczenia usuń dane i sprawdź
// czy twój program nadal działa.
//
// 1. Numery telefonów według nazwiska
// --------------------------
// Bowen 202-555-0179
// dulin 03.37.77.63.06
// Greco 03489940240
//
// Wydrukuj numer telefonu dulina.
//
//
// 2. Dostępność produktu według ID produktu
// ----------------
// 617841573 prawda
// 879401371 fałsz
// 576872813 prawda
//
// Czy jest dostępny numer produktu 879401371?
//
//
// 3. Wiele numerów telefonów według nazwiska
// ------------------------------------------------ ------
// Bowen [202-555-0179]
// dulin [03.37.77.63.06 03.37.70.50.05 02.20.40.10.04]
// Greko [03489940240 03489900120]
//
// Jaki jest drugi numer telefonu Greko?
//
//
// 4. Koszyk zakupów według identyfikatora klienta
// -------------------------------
// 100 [617841573:4 576872813:2]
// 101 [576872813:5 657473833:20]
// 102 []
//
// Ile z 576872813 kupi klient 101?
// (identyfikator produktu) (identyfikator klienta)
//
//
// OCZEKIWANY WYNIK
//
// 1. Uruchom program, aby zobaczyć wynik
// 2. Oto wynik z pustymi mapami:
//
//      dulin's phone number: N/A
//      Product ID #879401371 is not available
//      greco's 2nd phone number: N/A
//      Customer #101 is going to buy 5 from Product ID #576872813.
//
// ------------------------------------------------ ---------

func main() {

	phoneNumberByLastName := map[string]string{
		"Bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"Greco": "03489940240",
	}

	productAvailability := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}

	phoneNumbersByLastName := map[string][]string{
		"Bowen":   {"202-555-0179"},
		"Dulin":   {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"Greco":   {"03489940240", "03489900120"},
	}

	shoppingCart := map[int]map[int]int{
		100: {617841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
		102: {},
	}

	fmt.Println("dulin's phone number:",phoneNumberByLastName["dulin"])
	fmt.Println("Product ID #879401371 is",productAvailability[879401371])
	fmt.Println("greco's 2nd phone number:",phoneNumbersByLastName["Greco"][1])
	fmt.Println("Customer #101 is going to buy",shoppingCart[101][576872813],"from Product ID #576872813.")
}
