package graph

import (
	"log"
	"testing"
)

/**
1 ----- 2
|	     \
3 -- 4 -- 5
*/
func TestGraph(t *testing.T) {
	gr := New()
	n1 := gr.AddNode(1)
	n2 := gr.AddNode(2)
	n3 := gr.AddNode(3)
	n4 := gr.AddNode(4)
	n5 := gr.AddNode(5)

	gr.AddEdge(n1, n2, 1)
	gr.AddEdge(n1, n3, 1)
	gr.AddEdge(n3, n1, 3)
	gr.AddEdge(n3, n4, 1)
	gr.AddEdge(n4, n3, 1)
	gr.AddEdge(n4, n5, 1)
	gr.AddEdge(n2, n5, 1)

	nbs := gr.Neighbors(n3)
	mapNbs := make(map[Index]struct{})
	log.Print(nbs)
	for _, nb := range nbs {
		mapNbs[nb] = struct{}{}
	}
	if _, ok := mapNbs[n1]; !ok {
		t.Fatal("n1 not found")
	}
	if _, ok := mapNbs[n4]; !ok {
		t.Fatal("n4 not found")
	}
}

func TestGraph_Edges(t *testing.T) {
	gr := New()
	n1 := gr.AddNode(1)
	n2 := gr.AddNode(2)
	n3 := gr.AddNode(3)
	n4 := gr.AddNode(4)
	n5 := gr.AddNode(5)

	gr.AddEdge(n1, n2, 1)
	gr.AddEdge(n1, n3, 1)
	gr.AddEdge(n3, n1, 3)
	gr.AddEdge(n3, n4, 1)
	gr.AddEdge(n4, n3, 1)
	gr.AddEdge(n4, n5, 1)
	gr.AddEdge(n2, n5, 1)

	edges := gr.Edges()
	found := false
	for _, edge := range edges {
		if edge.StartIdx == n3 && edge.EndIdx == n1 && edge.Weight == 3 {
			found = true
		}
	}

	if !found {
		t.Log(edges)
		t.Fatalf("not found: startIdx:%d, endIdx:%d, weight:%d", n3, n1, 3)
	}

}
