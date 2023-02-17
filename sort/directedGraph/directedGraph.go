package main

import "fmt"

type Graph struct {
	nodes []*Node
}

// Node - список узлов
type Node struct {
	id    int // идентификатор узла
	edges []*Edge
}

// список исходящих ребер
type Edge struct {
	dest   *Node // конечный узел
	weight int   // вес
}

/* // хуита
// addNode - добавление элемента в граф
func (g *Graph) addNode(id int) *Node {
	// проверяем граф на узел с таким идентификатором
	for _, n := range g.nodes {
		if n.id == id {
			return n
		}
	}
	// если узла нет, то создаем новый узел и добавляем в граф
	node := &Node{
		id:    id,
		edges: make([]*Edge, 0),
	}
	g.nodes = append(g.nodes, node)
	return node
}
//addEdge - добавление ребер
func (n *Node)addEdge( dest *Node, weight int)  {
	edge := &Edge{
		dest:   dest,
		weight: weight,
	}
	n.edges = append(n.edges,edge)
}
*/

func (g *Graph) Print() {
	for _, n := range g.nodes {
		fmt.Printf("Node %d: ", n.id)
		for _, e := range n.edges {
			fmt.Printf("(%d, %d) ", e.dest.id, e.weight)
		}
		fmt.Printf("\n")
	}
}

func main() {
	g := Graph{}
	/*	node1 := g.AddNode(1)
		node2 := g.AddNode(2)
		node3 := g.AddNode(3)
		node4 := g.AddNode(4)
		node1.edges = append(node1.edges,&)*/

}
