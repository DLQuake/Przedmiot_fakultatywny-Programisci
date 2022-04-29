package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Zadelkaruj dwie struktury Artist(Name, Surname, Age) i oraz Car(Brand, HorsePower, Price)
//
// 1. Wydrukuj obie strkutury
//
// 2. Strukturę Artist wypełnij za pomocą literału struktury
//
// 3. Strukturę Car wypełnij za pomocą notacji kropki (zamienna typu tej struktury).
//
// 4. Wydrukuj obie sktruktury.
//

type Artist struct {
	Name    string
	Surname string
	Age     int
}

type Car struct {
	Brand      string
	HorsePower int
	Price      int
}

func main() {
	artist := Artist{}
	car := Car{}

	fmt.Println(artist)
	fmt.Println(car)

	artist = Artist{
		Name:    "Tom",
		Surname: "Biden",
		Age:     60,
	}

	car.Brand = "Ford"
	car.HorsePower = 150
	car.Price = 20000

	fmt.Println(artist)
	fmt.Println(car)

}
