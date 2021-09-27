package graph_test

import (
	"fmt"
	"go-data-structure/graph"
	"testing"
)

func TestMakeAdjacencyMatrixGraph(t *testing.T) {
	G := graph.MakeAdjacencyMatrixGraph()
	if &G == nil {
		t.Fatal("Failed at creating matrix")
	}
	G.AddVertix("A")
	str := fmt.Sprintf("%v", G)
	if str != "(A)->[]" {
		t.Fatalf("%v != (A)->[]", str)
	}
	G.AddVertix("B")
	G.AddVertix("C")
	G.AddVertix("D")
	G.AddVertix("E")
    G.AddEdge("A", "B")
    G.AddEdge("A", "C")
    G.AddEdge("B", "C")
    G.AddEdge("B", "E")
    G.AddEdge("B", "D")
    G.AddEdge("C", "D")
    G.AddEdge("D", "E")
	str = fmt.Sprintf("%v", G)
	if str != "(A)->[B,C] (B)->[C,D,E] (C)->[D] (E)->[]" {
		t.Fatalf("%v != (A)->[B,C] (B)->[C,D,E] (C)->[D] (E)->[]", str)
	}
}
