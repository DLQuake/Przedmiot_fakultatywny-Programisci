package main

import (
	"fmt"
	"time"
)

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// 1. Utwórz następujące wycinki zerowe:
// + Dodatki do pizzy
// + Godziny odlotów
// + Lata ukończenia studiów
// + Włączenie/wyłączenie świateł w pomieszczeniu
//
// 2. Dołącz do nich kilka elementów (wykorzystaj swoją kreatywność!)
//
// 3. Wydrukuj wszystkie wycinki
//
//
// OCZEKIWANY WYNIK (przykładowy)
//
// pizza : [cebula pepperoni extra ser]
//
// dyplomy : [1998 2005 2018]
//
// odloty : [2019-01-28 15:09:31.294594 +0300 +03 m=+0.000325020
// 2019-01-29 15:09:31.294594 +0300 +03 m=+86400.000325020
// 2019-01-30 15:09:31.294594 +0300 +03 m=+172800.000325020]
//
// lights : [prawda fałsz prawda]
//
//
// WSKAZÓWKI
// + W przypadku godzin odlotu użyj typu time.Time. Sprawdź jego dokumentację.
//
//  now := time.Now() -> Podaje aktualny czas
//  now.Add(time.Hour*24)-> Zwraca czas time.Time 24 godziny po `now`
//
// + W przypadku lat ukończenia szkoły możesz użyć typu int
// ------------------------------------------------ ---------

func main() {
	var dodatkiDoPizzy = []string{}
	var godzinyOdlotow = []time.Time{}
	var lataUkonczeniaStudiow = []int{}
	var wlaczenieSwiatla = []bool{}


	dodatkiDoPizzy = append(dodatkiDoPizzy, "cebula", "pepperoni", "extra ser")
	godzinyOdlotow = append(godzinyOdlotow, time.Now().Add(time.Hour*24), time.Now().Add(time.Hour*24*2), time.Now().Add(time.Hour*24*3))
	lataUkonczeniaStudiow = append(lataUkonczeniaStudiow, 1998, 2005, 2018)
	wlaczenieSwiatla = append(wlaczenieSwiatla, true, false, true)

	fmt.Println("pizza :", dodatkiDoPizzy)
	fmt.Println("dyplomy :", lataUkonczeniaStudiow)
	fmt.Println("odloty :", godzinyOdlotow)
	fmt.Println("lights :", wlaczenieSwiatla)

}
