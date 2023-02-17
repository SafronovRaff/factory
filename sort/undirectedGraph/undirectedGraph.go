// реализация неорентированного графа и алгоритма поиска в ширину
package main

import "fmt"

// Graph -структура графа, где ключи - узлы, а значения списки смежных узлов
type Graph struct {
	nodes map[int][]int
}

// AddEdge - добавляет ребро между 2 узлами графа
func (g *Graph) AddEdge(u, v int) {
	//пороверяем на существование узлов nodes в структуре графа
	if g.nodes == nil {
		g.nodes = make(map[int][]int)
	}
	// добавлем узкл в список смежных узлов графа
	g.nodes[u] = append(g.nodes[u], v)
	g.nodes[v] = append(g.nodes[v], u)
}
func main() {
	g := Graph{}
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	fmt.Println(g) //{map[0:[1] 1:[0 2] 2:[1 3] 3:[2]]}

}
 // поиск в ширину BFS
func ()  {
	
}