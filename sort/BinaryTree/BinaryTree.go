// binary tree
package main

import "fmt"

// Node - узел дерева в рекурсии
type Node struct {
	key   int
	left  *Node
	right *Node
}

// Bt - структура дерева
type Bt struct {
	root *Node
	size int
}

// NewBt - конструктор дерева
func NewBt() *Bt {
	return &Bt{
		root: nil,
		size: 0,
	}
}

// Size - возвращает размер дерева
func (t *Bt) Size() int {
	return t.size
}

// Insert - добавление элементов в дерево
func (t *Bt) Insert(value int) {
	if t.root == nil {
		t.root = &Node{
			key:   value,
			left:  nil,
			right: nil,
		}
	}
	t.size++
	t.root.insert(&Node{
		key:   value,
		left:  nil,
		right: nil,
	})
}

// insert - рекурсивный метод вставки узла
func (root *Node) insert(newNode *Node) {
	//если существует пропускаем
	if root.key == newNode.key {
		return
	}
	// правое дерево
	if root.key < newNode.key {
		if root.right == nil {
			root.right = newNode
		} else {
			root.right.insert(newNode)
		}
	} else {
		if root.left == nil {
			root.left = newNode
		} else {
			root.left.insert(newNode)
		}
	}
}

// Searchh- поиск эл в дереве
func (t *Bt) Search(value int) bool {
	t.size--
	return searchEl(t.root, value)
}

// searchEl - рекурсивный поиск
func searchEl(root *Node, value int) bool {
	if root != nil {
		if value == root.key {
			return true
		} else if value > root.key {
			return searchEl(root.right, value)
		} else {
			return searchEl(root.left, value)
		}
	}
	return false
}

// Delete - удалить элемент
func (t *Bt) Delete(value int) bool {
	if !t.Search(value) || t.root == nil {
		return false
	}
	if t.root.key == value {
		tempRoor := &Node{
			key:   0,
			left:  nil,
			right: nil,
		}
		tempRoor.left = t.root
		r := delete(t.root, tempRoor, value)
		t.root = tempRoor.left
		return r
	}
	return delete(t.root.left, t.root, value) || delete(t.root.right, t.root, value)
}

func delete(root *Node, parent *Node, value int) bool {
	switch {
	case root.key == value:
		if root.left != nil && root.right != nil {
			root.key = minValue(root.right)
			return delete(root.right, root, root.key)
		}
		link(parent, root)
		return true

	case root.key > value:
		if root.left == nil {
			return false
		}
		return delete(root.left, root, value)

	case root.key < value:
		if root.right == nil {
			return false
		}
		return delete(root.right, root, value)
	}
	return false
}

// link - соединяет узлы
func link(parent *Node, root *Node) {
	if parent.left == root {
		if root.left != nil {
			parent.left = root.left
		} else {
			parent.left = root.right
		}
	} else if parent.right == root {
		if root.left != nil {
			parent.left = root.left
		} else {
			parent.right = root.right
		}
	}
}

func minValue(root *Node) int {
	if root != nil {
		if root.left == nil {
			return root.key
		}
		return minValue(root.left)
	}
	return root.key
}

func maxValue(root *Node) int {
	if root != nil {
		if root.right == nil {
			return root.key
		}
		return maxValue(root.right)
	}
	return root.key
}

// печать мин значения
func (t *Bt) printMin() {
	fmt.Println(minValue(t.root))
}

// печатает макс значения
func (t *Bt) printMax() {
	fmt.Println(maxValue(t.root))
}

func (t *Bt) PrinfTree() {
	printNode(t.root)
}

// printNode - печать по слева на право, обход левое дерево, корень и правое дерево
func printNode(root *Node) {
	if root != nil {
		printNode(root.left)
		fmt.Println(root.key)
		printNode(root.right)

	}
}

func main() {
	tree := NewBt()
	tree.Insert(15)
	tree.Insert(1)
	tree.Insert(2)
	fmt.Println("Size: ", tree.Size())
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(3)
	fmt.Println("Size: ", tree.Size())
	fmt.Println("Min element: ")
	tree.printMin()
	fmt.Println("Search 4: ", tree.Search(4))
	fmt.Println("Search 3: ", tree.Search(3))
	tree.Insert(14)
	tree.Insert(17)
	tree.Insert(31)
	fmt.Println("Size: ", tree.Size())
	tree.PrinfTree()
	fmt.Println("Max element: ")
	tree.printMax()
	fmt.Println("Delete 31: ", tree.Delete(31))
	fmt.Println("Delete 5: ", tree.Delete(5))
	fmt.Println("Size: ", tree.Size())
	tree.PrinfTree()
}
