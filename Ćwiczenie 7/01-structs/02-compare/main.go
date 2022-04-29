package main

import (
	"fmt"
	"reflect"
)

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Modyfikując poprzednie ćwiczenie
//
// 1. Utwórz zmienną tego samego typu dowolnej struktury
//
// 2. Wypełnij ją dowolnym sposobem
//
// 3. Pierwszy raz takimi samymi wartości drugi raz wartości zmiennych powinny być inne
//
// 4. Porównaj obie strkutury za pomocą operatora == oraz użyj paczki reflect (porównaj te struktury 2 razy, raz z takimi
//    samymi wartościami raz z różnymi
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
	var artist Artist
	var car Car

	artist = Artist{
		Name:    "Tom",
		Surname: "Biden",
		Age:     60,
	}

	car.Brand = "Ford"
	car.HorsePower = 150
	car.Price = 20000

	artist2 := Artist{
		Name:    "Tom",
		Surname: "Biden",
		Age:     60,
	}

	car2 := Car{
		Brand:      "Ford",
		HorsePower: 150,
		Price:      20000,
	}

	fmt.Println(reflect.DeepEqual(artist, artist2))
	fmt.Println(reflect.DeepEqual(car, car2))

	fmt.Println(artist == artist2)
	fmt.Println(car == car2)
}
