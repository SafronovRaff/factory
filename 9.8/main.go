package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func (m Man) Count() int {
	return m.Crimes
}

func (m Man) Data() string {
	return fmt.Sprintf("Имя: %s,Кол-во орестов: %v", m.Name, m.Crimes)
}

var People = map[string]Man{
	"Вася щикатун":   {Name: "Вася", LastName: "Кунияде", Age: 51, Gender: "M", Crimes: 11},
	"Борис животное": {Name: "Борис", LastName: "Жудин", Age: 31, Gender: "M", Crimes: 17},
	"Ира кружка":     {Name: "Ирина", LastName: "Кружкина", Age: 17, Gender: "Ж", Crimes: 3},
	"Гена Букин":     {Name: "Олег", LastName: "Немещалов", Age: 24, Gender: "Ж", Crimes: 2},
	"Мохнорез":       {Name: "Иван", LastName: "Иванов", Age: 83, Gender: "Ж", Crimes: 1},
}

func main() {

	suspects := []string{"Гадя", "Надежда", "Борис животное", "Ира кружка"}
	/*var outlaw Man
	outlawBool := false
	for _, name := range suspects {
		human, ok := People[name]
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
	}*/

	fmt.Println(Criminal(People, suspects))
}

func Criminal(m map[string]Man, s []string) string {
	var criminals []string
	var keyCrimes string
	for _, name := range s {
		_, ok := m[name]
		if ok {
			criminals = append(criminals, name)
		}
	}
	if len(criminals) == 0 {
		return "В базе данных нет информации по запрошенным подозреваемым"
	}
	count := 0
	for _, crimes := range s {
		if People[crimes].Count() > count {
			count = People[crimes].Count()
			keyCrimes = crimes
		}
	}
	return People[keyCrimes].Data()
}
