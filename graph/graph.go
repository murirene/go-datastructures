package graph

import (
    "fmt"
    "strings"
)

type VISITED int

const (
        WHITE VISITED = iota
        GREY
        BLACK
    )

type Vertix struct {
    value string
    key int
    color VISITED
    parent *Vertix
    distance int
}

func (v Vertix) String() string {
    return fmt.Sprintf("(%s)", v.value)
}

type AdjacencyMatrixGraph struct {
    vertices map[string]Vertix
    matrix [][]int
}

func MakeAdjacencyMatrixGraph() AdjacencyMatrixGraph {
    vertices := make(map[string]Vertix, 0)
    matrix := make([][]int, 0)
    return AdjacencyMatrixGraph {
        vertices: vertices,
        matrix: matrix,
    }
}

func (g *AdjacencyMatrixGraph) AddVertix(key string) bool {
    _, ok := g.vertices[key] 
    if ok == true {
        return false
    }
    v := Vertix {
        value: key,
        key: len(g.vertices),
        color: WHITE,
        parent: nil,
        distance: 0,
    }

    matrix := make([][]int, len(g.vertices)+1)
    for i, _ := range matrix {
        matrix[i] = make([]int, len(g.vertices)+1)
    }

    for i:=0;i<len(g.matrix);i++ {
        for j:=0;j<len(g.matrix);j++ {
            matrix[i][j] = g.matrix[i][j]
        }
    }

    g.matrix = matrix
    g.vertices[key] = v

    return true
}

func (g AdjacencyMatrixGraph) String() string {
    var sb strings.Builder
    for key,_ := range g.vertices {
        sb.WriteString(fmt.Sprintf("(%s)", key))
    }
    return sb.String() 
}

