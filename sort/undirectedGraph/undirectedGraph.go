// Не орентированный граф и алгоритм поиска в ширину
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
	fmt.Println(g.BFS(2))
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	fmt.Println(g.BFS(4))
	fmt.Println(g)
}

// поиск в ширину BFS
func (g *Graph) BFS(start int) []int {
	visited := make(map[int]bool) // словарь для отслеживания посещенных вершин
	queue := make([]int, 0)       // очередь вершин, которые нужно обработать
	res := make([]int, 0)         // список посещенных вершин
	visited[start] = true         // помечаем начальную как посещенную
	queue = append(queue, start)  // добавляем начальную вершину в очередь

	//очередь не пуста, обрабатываем вершины
	for len(queue) > 0 {
		node := queue[0]        // порвая вершина из очереди
		queue = queue[1:]       // удаляем из очереди
		res = append(res, node) // добавляем в список посещенных

		//обработка соседей текущей вершины
		for _, neighbor := range g.nodes[node] {
			if !visited[neighbor] { // если сосед еще не посещен
				visited[neighbor] = true        //помечаем как посещенный
				queue = append(queue, neighbor) //добавляем соседа в очередь
			}
		}
	}
	return res // возвращает список посещенных вершин
}
