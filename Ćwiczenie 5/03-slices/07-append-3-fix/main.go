package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Napraw problemy w poniższym kodzie.
//
// BONUS
//
// Uprość kod.
//
// OCZEKIWANY WYNIK
//
// toppings: [black olives green peppers onions extra cheese]
//
// ------------------------------------------------ ---------

func main() {
	// toppings := []int{"black olives", "green peppers"}

	// var pizza [3]string
	// append(pizza, ...toppings)
	// pizza = append(toppings, "onions")
	// toppings = append(pizza, extra cheese)

	// fmt.Printf("pizza       : %s\n", pizza)

	toppings := []string{"black olives", "green peppers"}
	toppings = append(toppings, "onions")
	toppings = append(toppings, "extra cheese")
	fmt.Printf("toppings: %s\n", toppings)
}
