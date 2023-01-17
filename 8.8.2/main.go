package main

import (
	Au "factory/8.8.2/Auto"
	"fmt"
)

func main() {
	BMW := Au.ConstructAuto("bmw", "x", 5000, 4000, 2000, 362, 130)
	Dodge := Au.ConstructAuto("Dodge", "Avenger", 6000, 4000, 1800, 291, 160)
	Mercedes := Au.ConstructAuto("Mercedes", "c200", 5000, 400, 23000, 287, 150)
	printAuto(BMW)
	printAuto(Dodge)
	printAuto(Mercedes)
}

func printAuto(aut Au.Auto) {

	fmt.Println(aut.Brand(), aut.Model(), aut.EnginePower(), aut.MaxSpeed())

	switch aut.Brand() {
	case "Dodge":
		aut, ok := aut.(Au.Dimensions)
		if ok {
			fmt.Println(aut.Width().Get(Au.Inch), aut.Length().Get(Au.Inch), aut.Width().Get(Au.Inch))
		}

	default:
		aut, ok := aut.(Au.Dimensions)
		if ok {
			fmt.Println(aut.Width().Get(Au.CM), aut.Height().Get(Au.CM), aut.Length().Get(Au.CM))
		}

	}
}
