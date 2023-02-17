package main

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

func (t *BinaryTree) Remove(value int) {

}
