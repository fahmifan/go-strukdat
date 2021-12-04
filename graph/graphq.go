package graph

import (
	"fmt"
	"log"
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
	lock  sync.RWMutex
	edges map[Index]Weight

	Index Index
	Value interface{}
}

// AddEdge :nodoc:
func (n *Node) AddEdge(n2 Index, weight Weight) {
	n.lock.Lock()
	n.edges[n2] = weight
	n.lock.Unlock()
}

// Graph is a bidirectional weighted graph
type Graph struct {
	nodes []*Node
	lock  sync.RWMutex
}

func New() *Graph {
	return &Graph{}
}

func (g *Graph) AddNode(value interface{}) (index Index) {
	if value == nil {
		return
	}

	index = g.makeNodeIndex()
	node := &Node{Index: index, Value: value, edges: make(map[Index]Weight)}

	g.lock.Lock()
	g.nodes = append(g.nodes, node)
	g.lock.Unlock()

	return
}

func (g *Graph) makeNodeIndex() Index {
	return Index(len(g.nodes))
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
	if idx < 0 || int(idx) > g.nodeSize()-1 {
		return nil
	}

	g.lock.RLock()
	node := g.nodes[idx]
	g.lock.RUnlock()

	log.Println(node)

	node.lock.RLock()
	for edge := range node.edges {
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
		for edgeIdx, weight := range node.edges {
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

func (g *Graph) nodeSize() int {
	g.lock.RLock()
	size := len(g.nodes)
	g.lock.RUnlock()
	return size
}
