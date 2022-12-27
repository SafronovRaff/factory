package main

import (
	el "factory/8.8/electronic"
	"fmt"
)

func main() {

	Iphone := el.DesignerApplePhone("Iphone13pro")
	Android := el.DesignerAndroidPhone("Samsung", "a5")
	Panasonic := el.DesignerRadioPhone("Panasonic", "KX-TS2350RUW")
	printCharacteristics(Iphone)
	printCharacteristics(Android)
	printCharacteristics(Panasonic)
}

func printCharacteristics(phone el.Phone) {
	fmt.Printf("Brand- %v \n Type- %v \n Model- %v \n  ", phone.Brand(), phone.Type(), phone.Model())

	switch phone.Type() {
	case "smartphone":
		phone, ok := phone.(el.Smartphone)
		if ok {
			fmt.Printf("OS- %v\n\n\n ", phone.OS())
		}

	case "radioPhone":
		phone, ok := phone.(el.StationPhone)
		if ok {
			fmt.Printf("ButtonsCount- %v\n\n\n ", phone.ButtonsCount())
		}
	}

}
