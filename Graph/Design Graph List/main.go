package main

import "fmt"

// Graph hold graph data structure
type Graph struct {
	Vertices []*Vertex
}

// Vertex hold value of vertices
type Vertex struct {
	Key     int
	Adjency []*Vertex
}

var visited = make(map[int]bool)

func (g *Graph) addVertex(key int) {
	vertex := &Vertex{
		Key: key,
	}

	if ok := contain(g.Vertices, key); ok {
		fmt.Println("Vertex already exist in graph")
	} else {
		g.Vertices = append(g.Vertices, vertex)
	}
}

func (g *Graph) addEdge(u, v int) {
	fromVertex := g.getVertex(u)
	toVertex := g.getVertex(v)

	// check if vertex not exist print error
	if fromVertex == nil || toVertex == nil {
		fmt.Println("Vertex not exist in a graph")
	} else if contain(fromVertex.Adjency, toVertex.Key) {
		fmt.Println("Vertex", fromVertex.Key, "already connected with", toVertex.Key)
	} else {
		// directed graph
		// fromVertex.Adjency = append(fromVertex.Adjency, toVertex)

		//undirected graph
		fromVertex.Adjency = append(fromVertex.Adjency, toVertex)
		toVertex.Adjency = append(toVertex.Adjency, fromVertex)
	}
}

func (g *Graph) getVertex(key int) *Vertex {
	for k, v := range g.Vertices {
		if v.Key == key {
			return g.Vertices[k]
		}
	}

	return nil
}

func contain(vertex []*Vertex, key int) bool {
	for _, v := range vertex {
		if v.Key == key {
			return true
		}
	}

	return false
}

// DFS ...
func DFS(startVertex *Vertex) {
	if startVertex == nil {
		return
	}

	visited[startVertex.Key] = true

	fmt.Print(startVertex.Key, " ")

	for _, v := range startVertex.Adjency {
		if visited[v.Key] {
			continue
		}

		DFS(v)
	}
}

// BFS ...
func BFS(startVertex *Vertex) {
	visited := make(map[int]bool)

	queue := make([]*Vertex, 0)
	queue = append(queue, startVertex)
	queueIndex := 0

	if startVertex == nil {
		return
	}

	for len(queue) > 0 {
		vertex := queue[0]
		if visited[vertex.Key] {
			queue = queue[queueIndex+1:]
			continue
		}

		visited[vertex.Key] = true
		fmt.Print(vertex.Key, " ")

		for _, v := range vertex.Adjency {
			queue = append(queue, v)
		}

		queue = queue[queueIndex+1:]
		// fmt.Println("len", len(queue))
	}
}

func print(graph *Graph) {
	for _, v := range graph.Vertices {
		fmt.Print("Vertex ", v.Key, " : ")

		for _, v := range v.Adjency {
			fmt.Print(v.Key, " ")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("--- Graph Design Adjency List ---")

	graph := &Graph{}

	for i := 0; i < 5; i++ {
		graph.addVertex(i)
	}

	// graph.addEdge(1, 9)
	// graph.addEdge(1, 5)
	// graph.addEdge(1, 2)
	// graph.addEdge(5, 8)
	// graph.addEdge(8, 9)
	// graph.addEdge(9, 10)

	graph.addEdge(0, 1)
	graph.addEdge(0, 2)
	graph.addEdge(0, 3)
	graph.addEdge(1, 4)

	print(graph)

	fmt.Print("DFS ")
	DFS(graph.Vertices[0])

	fmt.Println()

	fmt.Print("BFS ")
	BFS(graph.Vertices[0])
	fmt.Println()

	// fmt.Println(graph.Vertices[0])
}
