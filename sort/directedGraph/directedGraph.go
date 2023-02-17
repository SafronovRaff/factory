package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

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

// addEdge - добавление ребер
func (n *Node) addEdge(dest *Node, weight int) {
	edge := &Edge{
		dest:   dest,
		weight: weight,
	}
	n.edges = append(n.edges, edge)
}

// addNodeWithEdge - добавлять новый узел в граф, а затем добавляет ребро между этим узлом и заданным узлом-назначением
func (g *Graph) addNodeWithEdge(id, destID, weight int) {
	var sourceNode *Node // исходный узел
	var destNode *Node   //узел назначения
	// проверяем наличие исходного узла в графе
	for _, n := range g.nodes {
		if n.id == id {
			sourceNode = n
			break
		}
	}
	//если исходный узел не найден то добавляем его в граф
	if sourceNode == nil {
		sourceNode = &Node{
			id:    id,
			edges: make([]*Edge, 0)}
		g.nodes = append(g.nodes, sourceNode)
	}
	// поиск узла назначения
	for _, n := range g.nodes {
		if n.id == destID {
			destNode = n
			break
		}
	}
	//если узел назначения не найден, добавляем в граф
	if destNode == nil {
		destNode = &Node{
			id:    destID,
			edges: make([]*Edge, 0),
		}
		g.nodes = append(g.nodes, destNode)
	}

	//добавляем ребро исходным узлом и узлом назначения
	sourceNode.addEdge(destNode, weight)

}

// Print - печать графа
func (g *Graph) Print() {
	for _, n := range g.nodes {
		fmt.Printf("Node %d: ", n.id)
		for _, e := range n.edges {
			fmt.Printf("(%d, %d) ", e.dest.id, e.weight)
		}
		fmt.Printf("\n")
	}
}

// printShortestPath -печать кратчайшего пути
func (g *Graph) printShortestPath(sourceID, destID int) {
	//получаем короткий путь
	path := g.shortestPath(sourceID, destID)

	if path == nil {
		fmt.Printf("путь из %d в узел %d не найден \n", sourceID, destID)
	}
	//строим путь
	var pathStr string
	for i, nodeID := range path {
		if i > 0 {
			pathStr += " -> "
		}
		pathStr += strconv.Itoa(nodeID)
	}

	fmt.Printf("Кратчайший путь  из узла %d в узел %d : %s \n", sourceID, destID, pathStr)

}

//shortestPath - определения кратчайшего пути между любой парой вершин
//алгоритм поиска в ширину(BFS) с учетом весов ребер,
//находит кратчайший путь от начальной вершины до всех других вершин в графе.

func (g *Graph) shortestPath(sourceID, destID int) []int {
	// карта расстояний для каждого узла
	distances := make(map[int]int)

	//карта предшественников для каждого узла
	predecessors := make(map[int]*Node)

	//очередь для обхода в ширину
	queue := list.New()

	//растояние от исходного узла до него самого - 0
	distances[sourceID] = 0

	//добавляем исходный узел в очередь
	//PushBack вставляет новый элемент e со значением v в конце списка l и возвращает e.
	queue.PushBack(g.getNode(sourceID))

	//пока очередь не пуста
	for queue.Len() > 0 {
		//получаем первый узел из очереди.
		//Remove удаляет e из l, если e является элементом списка l. Он возвращает значение элемента e.Value. Элемент не должен быть нулевым.
		node := queue.Remove(queue.Front()).(*Node)

		//перебираем все ребра узла
		for _, edge := range node.edges {

			// считаем растояние узла до узла - назначения
			distance := distances[node.id] + edge.weight

			//если растояние до узла - назначения не установленно, или новое расстояние меньше текущего
			if _, ok := distances[edge.dest.id]; !ok || distance < distances[edge.dest.id] {
				//устанавливаем новое расстояние для узла - назначения
				distances[edge.dest.id] = distance

				//устанавливаем предшественника для узла - назначения
				predecessors[edge.dest.id] = node
				//добавляем узел - назначения в очередь

				//добавляем узел - назначения в очередь
				queue.PushBack(edge.dest)
			}
		}
	}
	// если расстояние до узла - назначения не установленно, путь не найден
	if _, ok := distances[destID]; !ok {
		return nil
	}
	//собираем путь от узла - назначения до исходного узла
	path := []int{destID}

	for node := predecessors[destID]; node != nil; node = predecessors[node.id] {
		path = append([]int{node.id}, path...)
	}
	return path
}

// getNode - находит узел в графе по Id
func (g *Graph) getNode(id int) *Node {
	for _, node := range g.nodes {
		if node.id == id {
			return node
		}
	}
	return nil
}

// randomGraph - заполняет граф случайными значениями принимает колл узлов, колл рёбер и макс вес ребра.
func (g *Graph) randomGraph(numNode, numEdges, maxWeight int) {
	//	генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// добавляем узлы графа
	for i := 0; i < numNode; i++ {
		g.addNode(i)
	}

	//добавляем случайные ребра
	for i := 0; i < numEdges; i++ {
		//случайный узел - источник
		sourseID := rand.Intn(numNode)

		//случайный узел - назначени
		destID := rand.Intn(numNode)

		//случайный вес ребра
		weight := rand.Intn(maxWeight) + 1

		//добавляем ребро в граф
		g.addEdgeWeight(sourseID, destID, weight)
	}

}

// addEdgeWeight - получет индификаторы 2х узлов и вес ребра между ними
func (g *Graph) addEdgeWeight(sourseID, destID, weight int) {

	sourseNode := g.getNode(sourseID) // getNode - находит узел в графе по Id
	destNodec := g.getNode(destID)    // создаём новое ребро
	//добавляем ребро из исходного в назначенное

	edge := &Edge{
		dest:   destNodec,
		weight: weight,
	}
	sourseNode.edges = append(sourseNode.edges, edge)
}

func main() {
	g := Graph{}
	/*	node1 := g.addNode(1)
		node2 := g.addNode(2)
		node3 := g.addNode(3)
		node4 := g.addNode(4)
		node1.addEdge(node2, 1)
		node1.addEdge(node3, 2)
		node2.addEdge(node4, 3)
		node3.addEdge(node4, 4)
	*/
	g.addNodeWithEdge(1, 2, 1)
	g.addNodeWithEdge(1, 3, 2)
	g.addNodeWithEdge(2, 4, 3)
	g.addNodeWithEdge(3, 4, 4)

	g.Print()
	g.printShortestPath(1, 4)

	g.randomGraph(15, 40, 3)
	g.Print()

	g.printShortestPath(11, 6)
	g.printShortestPath(2, 7)
	g.printShortestPath(3, 9)
	g.printShortestPath(14, 7)
	g.printShortestPath(6, 10)
	g.printShortestPath(13, 1)

}
