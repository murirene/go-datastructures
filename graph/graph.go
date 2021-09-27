package graph

import (
	"errors"
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
	value    string
	key      int
	color    VISITED
	parent   *Vertix
	distance int
}

func (v Vertix) String() string {
	return fmt.Sprintf("(%s)", v.value)
}

type AdjacencyMatrixGraph struct {
	vertices map[string]Vertix
	matrix   [][]int
}

func MakeAdjacencyMatrixGraph() AdjacencyMatrixGraph {
	vertices := make(map[string]Vertix, 0)
	matrix := make([][]int, 0)
	return AdjacencyMatrixGraph{
		vertices: vertices,
		matrix:   matrix,
	}
}

func (g *AdjacencyMatrixGraph) AddVertix(key string) bool {
	_, ok := g.vertices[key]
	if ok == true {
		return false
	}
	v := Vertix{
		value:    key,
		key:      len(g.vertices),
		color:    WHITE,
		parent:   nil,
		distance: 0,
	}

	matrix := make([][]int, len(g.vertices)+1)
	for i, _ := range matrix {
		matrix[i] = make([]int, len(g.vertices)+1)
	}

	for i := 0; i < len(g.matrix); i++ {
		for j := 0; j < len(g.matrix); j++ {
			matrix[i][j] = g.matrix[i][j]
		}
	}

	g.matrix = matrix
	g.vertices[key] = v

	return true
}

func (g *AdjacencyMatrixGraph) AddEdge(key string, keyEdge string) (bool, error) {
	v1, ok1 := g.vertices[key]
	v2, ok2 := g.vertices[keyEdge]

	if ok1 == false || ok2 == false {
		return false, errors.New("Vertex does not exist")
	}

	if len(g.matrix) <= v1.key || len(g.matrix) <= v2.key {
		return false, errors.New("Adjency Matrix Key is not within range")
	}

	g.matrix[v1.key][v2.key] = 1

	return true, nil
}

func (g AdjacencyMatrixGraph) String() string {
    var ( getKey = func(vMap map[string]Vertix, idx int) (string, error) {
        for key, v := range vMap {
            if v.key == idx {
                return key, nil
            }
        }

        return "", errors.New("Key does not exist")
    })

	var sb strings.Builder
	for key, v := range g.vertices {
		var sb2 strings.Builder
		edges := g.matrix[v.key]

		for i, v := range edges {
			if v == 1 {
                str, err := getKey(g.vertices, i)
                if err != nil {
                    str = "Not Found"
                }
				sb2.WriteString(fmt.Sprintf(" %s", str))
			}
		}

		sb.WriteString(fmt.Sprintf("(%s)->[%s]", key, sb2.String()))
	}
	return sb.String()
}
