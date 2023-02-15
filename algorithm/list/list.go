package main

import "fmt"

// ErrWrongList - ошибка при создании не существующего в списке инденкса
var ErrWrongList = fmt.Errorf("неверный индекс")

// IntNode - описание типа узел списка
type IntNode struct {
	Value int
	Next  *IntNode
}

// intList - описание типа список целых чисел
type IntList struct {
	size int
	Head *IntNode
}

// NewIntNode - создание нового узла списка
func NewIntNode(value int) *IntNode {
	return &IntNode{
		Value: value,
		Next:  nil,
	}
}

// NewIntList -создание нового списка целых чисел
func NewIntList() *IntList {
	return &IntList{
		size: 0,
		Head: nil,
	}
}

// Size - получение размера списка
func (I IntList) Size() int {
	return I.size
}

// Get - получение произвольного элеменнта списка
func (I IntList) Get(index int) (*IntNode, error) {
	if index < 0 || index >= I.Size() {
		return nil, ErrWrongList
	}
	node := I.Head
	for i := 1; i <= index; i++ {
		node = node.Next
	}
	return node, nil
}

// Set - обновление произвольного элемента списка
func (I *IntList) Set(el, index int) error {
	if index < 0 || index >= I.Size() {
		return ErrWrongList
	}
	node, err := I.Get(index)
	if err != nil {
		return err
	}
	node.Value = el
	return nil
}

// Add- добавление нового элемента в начало списка
func (I *IntList) Add(el int) {
	newNode := NewIntNode(el)
	if I.Head == nil {
		I.Head = newNode
	} else {
		newNode.Next = I.Head
		I.Head = newNode
	}
	I.size++
}

// Insert - вставка элемента в произвольную позицию
func (I *IntList) Insert(el, index int) error {
	if index < 0 || index >= I.size {
		return ErrWrongList
	}
	newNode := NewIntNode(el)
	if index == 0 {
		I.Add(el)
		return nil
	}
	node, err := I.Get(index - 1)
	if err != nil {
		return err
	}
	newNode.Next = node.Next
	node.Next = newNode
	I.size++
	return nil
}

// Remove - удаление элемента в произвольной позиции
func (I *IntList) Remove(index int) error {
	if index < 0 || index >= I.size {
		return ErrWrongList
	}
	node, err := I.Get(index - 1)
	if err != nil {
		return err
	}
	node.Next = node.Next.Next
	I.size--
	return nil
}

// Print - печать списка
func (I IntList) Print() {
	node := I.Head
	if node != nil {
		for node != nil {
			fmt.Printf("%d \t ", node.Value)
			node = node.Next
		}
		fmt.Printf("\n ")
	} else {
		fmt.Println("список пуст")
	}

}

func main() {
	list := NewIntList()
	list.Print()
	list.Add(2)
	list.Add(1)
	list.Add(0)
	list.Print()
	list.Insert(0, 1)
	list.Print()
	list.Remove(1)
	list.Print()
	list.Size()
}
