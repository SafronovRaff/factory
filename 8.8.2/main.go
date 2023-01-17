package main

import (
	Au "factory/8.8.2/Auto"
	"fmt"
)

func main() {
	BMW := Au.ConstructAuto("bmw", "x", 5, 4, 2, 362, 130)
	Dodge := Au.ConstructAuto("Dodge", "Avenger", 6, 4, 1.8, 291, 160)

	printAuto(BMW)
	printAuto(Dodge)
}

func printAuto(auto Au.Auto, dim Au.Dimensions) {

	fmt.Println(auto.Brand(), auto.Model(), auto.EnginePower(), auto.MaxSpeed(), dim.Width(), dim.Length(), dim.Height())

	switch Au.Auto.Brand(auto) {
	case "Dodge":
		fmt.Println(dim.Width().Get(Au.Inch), dim.Length().Get(Au.Inch), dim.Width().Get(Au.Inch))
	default:
		fmt.Println(dim.Width().Get(Au.CM), dim.Height().Get(Au.CM), dim.Length().Get(Au.CM))
	}
}
