package main

import "fmt"

// Node - узел  дерева
type Node struct {
	value int
	left  *Node // ссылка левое поддерево
	right *Node // ссылка правое поддерево
}

// BinaryTree - дерево
type BinaryTree struct {
	roor *Node // ссылка на корневой узел
}

// Insert - добавление элемента в дерево
func (t *BinaryTree) Insert(value int) {
	if t.roor == nil {
		//если дерево пустое создаем корень
		t.roor = &Node{
			value: value,
			left:  nil,
			right: nil,
		}
		return
	}
	node := t.roor
	for node != nil {
		if value < node.value {
			// если значение меньше текущего узла идём в левое поддерево
			if node.left == nil {
				// если левое поддерево пустое, создаём новый узел
				node.left = &Node{
					value: value,
					left:  nil,
					right: nil,
				}
				return
			}
			node = node.left
		} else {
			//если значение больше или равно идем в правое поддерево
			if node.right == nil {
				// если правое поддерево пустое, создаём новый узел
				node.right = &Node{
					value: value,
					left:  nil,
					right: nil,
				}
				return
			}
			node = node.right
		}
	}
}

// Remove - удаляет элемент из дерева
func (t *BinaryTree) Remove(value int) {
	t.roor = removeNode(t.roor, value)
}
func removeNode(node *Node, value int) *Node {
	// Проверка наличия узла в дереве
	if node == nil {
		return nil
	}
	//если значение совпало с заданным значением, начинается процесс удаления
	if value == node.value {
		// Без детей
		if node.left == nil && node.right == node {
			return nil
		}
		// 	Удаление с одним ребенком
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}

		// Удаление с 2 детьми, находим наименьший эл. в правом поддереве, который заменит удаляемый узел.
		// Значения этого эл. копируем в удаляемый узел и далее наименьший элемент удаляется из правого поддерева

		smallestRight := node.right //наименьшее в правом поддереве
		// в цикле перебор всех левых потомков правого дерева, когда левых потомков
		// больше не существует цикл прерывается
		for {
			if smallestRight != nil && smallestRight.left != nil {
				smallestRight = smallestRight.left
			} else {
				break
			}
		}
		// присваеваем значение удляемого узла и рекурсивно вызываем для удаления этого узла из правого поддерева
		node.value = smallestRight.value
		node.right = removeNode(node.right, smallestRight.value)
		// удаляемый узел заменяется наименьшим эл. из правого поддерева и возвращаем обновленное дерево
		return node
		// если значение меньше узла node.value, то рекурсивное удаление продолжится в левом поддереве
	} else if value < node.value {
		node.left = removeNode(node.left, value)
		return node
	} // если значение больше узла node.value, то рекурсивное удаление продолжится в правом поддереве
	node.right = removeNode(node.right, value)
	return node
}

// Find - поиск элемента в дереве
func (t *BinaryTree) Find(value int) bool {
	node := t.roor //записываем корень дерева

	for node != nil {
		if value == node.value {
			return true
		}
		if value < node.value { // идем в левое поддерево
			node = node.left
		} else { // идем в правое поддерево
			node = node.right
		}
	}
	return false
}
func (t *BinaryTree) Print() {
	printTree(t.roor, 0)
}
func printTree(node *Node, depth int) { // depth текущая глубина узла в дерева
	if node == nil {
		return
	}
	printTree(node.right, depth+1)
	fmt.Printf("%*s%d\n", depth*4, "", node.value)
	printTree(node.left, depth+1)
}

func main() {
	tree := BinaryTree{}
	tree.Insert(11)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(25)
	tree.Insert(17)
	tree.Insert(10)
	tree.Insert(250)
	tree.Insert(888)
	tree.Insert(112)
	tree.Insert(400)
	tree.Print()
	fmt.Println(tree.Find(11))
	fmt.Println(tree.Find(12))
	tree.Remove(888)
	tree.Remove(4)
	tree.Remove(3)
	tree.Print()
}
