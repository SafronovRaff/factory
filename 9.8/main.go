package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	people := map[string]Man{
		"Вася щикатун":   {Name: "Вася", LastName: "Кунияде", Age: 51, Gender: "M", Crimes: 11},
		"Борис животное": {Name: "Борис", LastName: "Жудин", Age: 31, Gender: "M", Crimes: 17},
		"Ира кружка":     {Name: "Ирина", LastName: "Кружкина", Age: 17, Gender: "Ж", Crimes: 3},
		"Гена Букин":     {Name: "Олег", LastName: "Немещалов", Age: 24, Gender: "Ж", Crimes: 2},
		"Мохнорез":       {Name: "Иван", LastName: "Иванов", Age: 83, Gender: "Ж", Crimes: 1},
	}
	suspects := []string{"Гадя", "Надежда", "Борис животное", "Ира кружка"}
	var outlaw Man
	outlawBool := false
	for _, name := range suspects {
		human, ok := people[name]
		if !ok {
			continue
		}
		if human.Crimes > outlaw.Crimes {
			outlaw = human
			outlawBool = true
		}
	}
	if outlawBool {
		fmt.Println("подозреваемый, у которого наибольшее количество совершённых преступлений: ", outlaw.Name)
	}
	if !outlawBool {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}
}
