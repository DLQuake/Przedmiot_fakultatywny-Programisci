package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Napraw poniższy kod
//
// 1. Tak aby nie wywoływała się panica oraz była możliwość modyfikacji obiektu

type Actress struct {
	Name string
}

func main() {
	var actress *Actress
	fmt.Printf("1) %+v\n", actress)
	actress = &Actress{Name: "Marilyn Monroe"}
	changeName(actress)

	fmt.Printf("3) %+v\n", actress)
}

func changeName(actress *Actress) {
	if actress == nil {
		return
	}
	actress.Name = "Dominic Toretto"
}

