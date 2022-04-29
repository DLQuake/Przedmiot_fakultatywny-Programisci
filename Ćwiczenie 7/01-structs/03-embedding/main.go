package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Zadelkaruj dwie struktury Team(League, Country) oraz Player(Name, Surname, Age)
//
// 1. Osadź strukturę Team w strukturze Player
//
// 2. Wypełnij ją dowolnym sposobem
//
// 4. Wydrukuj strukturę
//

type Team struct {
	League  string
	Country string
}

type Player struct {
	Name    string
	Surname string
	Age     int
	Team
}

func main() {
	player := Player{
		Name:    "John",
		Surname: "Doe",
		Age:     50,
		Team: Team{
			League:  "Premier League",
			Country: "England",
		},
	}

	fmt.Println(player)
}
