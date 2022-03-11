package main

// ------------------------------------------------ ---------
// ĆWICZENIE: Wydrukuj typ
//
// Wydrukuj typ i wartość 3 za pomocą fmt.Printf
// Wydrukuj typ i wartość 3.14 za pomocą fmt.Printf
// Wydrukuj typ i wartość "hello" za pomocą fmt.Printf
// Wydrukuj typ i wartość true za pomocą fmt.Printf
//
// OCZEKIWANY WYNIK
// Typ 3 to int
// Typ 3.14 to float64
// Typ hello to string
// Typ true to bool
// ------------------------------------------------ ---------

import "fmt"

func main() {

	i, f, s, b := 3, 3.14, "hello", true

	fmt.Printf("Typ %d to %[1]T\n",i)
	fmt.Printf("Typ %g to %[1]T\n",f)
	fmt.Printf("Typ %s to %[1]T\n",s)
	fmt.Printf("Typ %t to %[1]T\n",b)
}
