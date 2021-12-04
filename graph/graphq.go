package graph

import (
	"fmt"
	"os"
	"sync"
)

// Index ..
type Index int

// Weight ..
type Weight int

type Edge struct {
	StartIdx Index
	EndIdx   Index
	Weight   Weight
}

type Node struct {
	lock sync.RWMutex

	Edges map[Index]Weight
	Index Index
	Value interface{}
}

// AddEdge add an Edge from the Node to n2
func (n *Node) AddEdge(n2 Index, weight Weight) {
	n.lock.Lock()
	n.Edges[n2] = weight
	n.lock.Unlock()
}

// Graph is a bidirectional weighted adjacency list graph
type Graph struct {
	nodes []*Node
	lock  sync.RWMutex
}

// New return new Graph
func New() *Graph {
	return &Graph{}
}

// Node return the node of idx
func (g *Graph) Node(idx Index) *Node {
	if !g.validateIdx(idx) {
		return nil
	}
	g.lock.RLock()
	node := g.nodes[idx]
	g.lock.RUnlock()
	return node
}

// AddNode add new node to the graph
func (g *Graph) AddNode(value interface{}) (index Index) {
	if value == nil {
		return
	}

	index = g.nextIndex()
	node := &Node{Index: index, Value: value, Edges: make(map[Index]Weight)}

	g.lock.Lock()
	g.nodes = append(g.nodes, node)
	g.lock.Unlock()

	return
}

// AddEdge add an weigheted edge from n1 --> n2
func (g *Graph) AddEdge(n1, n2 Index, weight Weight) {
	if n1 < 0 || n2 < 0 {
		return
	}

	g.lock.RLock()
	node1 := g.nodes[int(n1)]
	node2 := g.nodes[int(n2)]
	g.lock.RUnlock()

	if node1 == nil || node2 == nil {
		fmt.Fprintf(os.Stderr, "n1 or n2 is not found: n1:%d n2:%d", n1, n2)
		return
	}

	node1.AddEdge(n2, weight)
}

// Neighbors return all neighbors index of the idx
func (g *Graph) Neighbors(idx Index) (neigbors []Index) {
	if idx < 0 || int(idx) > g.nodeLen()-1 {
		return nil
	}

	g.lock.RLock()
	node := g.nodes[idx]
	g.lock.RUnlock()

	node.lock.RLock()
	for edge := range node.Edges {
		neigbors = append(neigbors, edge)
	}
	node.lock.RLock()

	return neigbors
}

// Edges return all edges in the graph
func (g *Graph) Edges() (edges []Edge) {
	g.lock.RLock()
	nodes := g.nodes
	g.lock.RUnlock()

	for _, node := range nodes {
		if node == nil {
			continue
		}

		node.lock.RLock()
		for edgeIdx, weight := range node.Edges {
			edges = append(edges, Edge{
				StartIdx: node.Index,
				EndIdx:   edgeIdx,
				Weight:   weight,
			})
		}
		node.lock.RUnlock()
	}
	return edges
}

func (g *Graph) nodeLen() int {
	g.lock.RLock()
	size := len(g.nodes)
	g.lock.RUnlock()
	return size
}

func (g *Graph) nextIndex() Index {
	return Index(g.nodeLen())
}

func (g *Graph) validateIdx(idx Index) bool {
	if idx < 0 || int(idx) > g.nodeLen()-1 {
		return false
	}
	return true
}

func (g *Graph) BFS(rootIdx Index, cb func(idx Index)) {
	queue := NewQueue()
	visitedIdx := make([]bool, g.nodeLen()+1)

	queue.Enqueue(rootIdx)
	visitedIdx[rootIdx] = true

	for {
		if queue.IsEmpty() {
			break
		}

		el := queue.Pop()
		elIdx := el.Value.(Index)

		if cb != nil {
			cb(elIdx)
		}

		neighbors := g.Neighbors(elIdx)
		for _, i := range neighbors {
			if !visitedIdx[i] {
				queue.Enqueue(Index(i))
				visitedIdx[i] = true
				continue
			}
		}
	}
}
