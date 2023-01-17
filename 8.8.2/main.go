package main

import (
	Au "factory/8.8.2/Auto"
	"fmt"
)

func main() {
	BMW := Au.NewConstructAuto("bmw", "x", 500, 300, 200, 360, 120, "cm")

	Mercedes := Au.NewConstructAuto("Mercedes", "c200", 500, 300, 230, 287, 286, "cm")

	Dodge := Au.NewConstructAuto("Dodge", "Avenger", 600, 300, 180, 291, 320, "cm")

	/*
		//works
		fmt.Println("Brand", BMW.Brand(), "Model", BMW.Model(), "EnginePower", BMW.EnginePower(), "MaxSpeed", BMW.MaxSpeed(),
			"Length", BMW.Length().Get(Au.CM), "Height", BMW.Height().Get(Au.CM), "Width", BMW.Width().Get(Au.CM))

		fmt.Println("Brand", Mercedes.Brand(), "Model", Mercedes.Model(), "EnginePower", Mercedes.EnginePower(), "MaxSpeed",
			Mercedes.MaxSpeed(), "Length", Mercedes.Length().Get(Au.CM), "Height", Mercedes.Height().Get(Au.CM), "Width", Mercedes.Width().Get(Au.CM))

		fmt.Println("Brand", Dodge.Brand(), "Model", Dodge.Model(), "EnginePower", Dodge.EnginePower(), "MaxSpeed",
			Dodge.MaxSpeed(), "Length", Dodge.Length().Get(Au.Inch), "Height", Dodge.Height().Get(Au.Inch), "Width", Dodge.Width().Get(Au.Inch))*/

	printAuto(BMW)
	printAuto(Mercedes)
	printAuto(Dodge)
}

func printAuto(aut *Au.AllAuto) {
	fmt.Println("Brand", aut.Brand(), "Model", aut.Model(), "EnginePower", aut.EnginePower())
	switch aut.Brand() {
	case "Dodge":
		fmt.Println("Width", aut.Width().Get(Au.Inch), "Height", aut.Height().Get(Au.Inch), "Length", aut.Length().Get(Au.Inch))

	default:
		fmt.Println("Width", aut.Width().Get(Au.CM), "Height", aut.Height().Get(Au.CM), "Length", aut.Length().Get(Au.CM))

	}
}
