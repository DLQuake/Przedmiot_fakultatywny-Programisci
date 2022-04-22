package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Utwórz i wydrukuj następujące mapy.
//
// 1. Numery telefonów według nazwiska
// 2. Dostępność produktu według ID produktu
// 3. Wiele numerów telefonów według nazwiska
// 4. Koszyk zakupów według identyfikatora klienta
//
// Każdy produkt w koszyku ma identyfikator produktu i ilość. Dzięki mapie możesz stwierdzić:
// „Pan X kupił banany Y”
// ---------------------------------------------------------

func main() {

	// #1
	// Key        : Last name
	// Element    : Phone number

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable

	// #3
	// Key        : Last name
	// Element    : Phone numbers


	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity

	phoneNumberByLastName := map[string]string{
		"Smith":   "+48123456789",
		"Johnson": "+48123456789",
		"Williams": "+48123456789",
		"Jones":   "+48123456789",
		"Brown":   "+48123456789",
	}

	productAvailability := map[int]bool{
		1: true,
		2: false,
		3: true,
		4: false,
	}

	phoneNumbersByLastName := map[string][]string{
		"Smith":   {"+48123456789", "+48123456789"},
		"Johnson": {"+48123456789", "+48123456789"},
		"Williams": {"+48123456789", "+48123456789"},
		"Jones":   {"+48123456789", "+48123456789"},
	}

	shoppingCart := map[int]map[int]int{
		1: {576872813: 1, 657473833: 2},
		2: {617841573: 5, 576872813: 3},
		3: {123456789: 1, 987654321: 5},
	}

	fmt.Println("#1 Numery telefonów według nazwiska")
	for lastName, phoneNumber := range phoneNumberByLastName {
		fmt.Printf("%s: %s\n", lastName, phoneNumber)
	}

	fmt.Println("#2 Dostępność produktu według ID produktu")
	for productID, available := range productAvailability {
		fmt.Printf("%d: %t\n", productID, available)
	}

	fmt.Println("#3 Wiele numerów telefonów według nazwiska")
	for lastName, phoneNumbers := range phoneNumbersByLastName {
		fmt.Printf("%s: %s\n", lastName, phoneNumbers)
	}

	fmt.Println("#4 Koszyk zakupów według identyfikatora klienta")
	for customerID, products := range shoppingCart {
		fmt.Printf("Customer %d:\n", customerID)
		for productID, quantity := range products {
			fmt.Printf("\t%d: %d\n", productID, quantity)
		}
	}

}
