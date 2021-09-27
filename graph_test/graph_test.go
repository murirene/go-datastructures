package graph_test

import (
    "testing"
    "fmt"
    "go-data-structure/graph"
)

func TestMakeAdjacencyMatrixGraph(t *testing.T) {
    G := graph.MakeAdjacencyMatrixGraph()
    if &G == nil {
        t.Fatal("Failed at creating matrix")
    }
    G.AddVertix("A")
    str := fmt.Sprintf("%v", G)
    if str != "(A)" {
        t.Fatalf("%v != (A)", str)
    }
    G.AddVertix("B")
    G.AddVertix("C")
    G.AddVertix("D")
    G.AddVertix("E")

    str = fmt.Sprintf("%v", G)
    if str != "(A)(B)(C)(D)(E)" {
        t.Fatalf("%v != (A)(B)(C)(D)(E)", str)
    }
}
