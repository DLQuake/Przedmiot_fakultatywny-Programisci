package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Czy możesz uprościć instrukcję if w poniższym kodzie?
//
//  Kiedy:
// isSphere == true i radius jest równy lub większy niż 200
//
// Wypisze "It's a big sphere."
//
// W przeciwnym razie wypisze "I don't know."
//
// OCZEKIWANY WYNIK
// "It's a big sphere."
// ------------------------------------------------ ---------

func main() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	// var big bool

	// if radius >= 50 {
	// 	if radius >= 100 {
	// 		if radius >= 200 {
	// 			big = true
	// 		}
	// 	}
	// }

	// if big != true {
	// 	fmt.Println("I don't know.")
	// } else if !(isSphere == false) {
	// 	fmt.Println("It's a big sphere.")
	// } else {
	// 	fmt.Println("I don't know.")
	// }

	if isSphere == true && radius >= 200 {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}
