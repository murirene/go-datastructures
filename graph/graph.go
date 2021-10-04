package graph

import (
	"errors"
	"fmt"
	"sort"
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
}

func (v Vertix) String() string {
	return fmt.Sprintf("(%s)", v.value)
}

type AdjacencyMatrixGraph struct {
	vertices map[string]*Vertix
	matrix   [][]int
}

func MakeAdjacencyMatrixGraph() AdjacencyMatrixGraph {
	vertices := make(map[string]*Vertix, 0)
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
	v := &Vertix{
		value:    key,
		key:      len(g.vertices),
		color:    WHITE,
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

func (g AdjacencyMatrixGraph) getVertixByIdx(idx int) (string, error) {
	for key, v := range g.vertices {
		if v.key == idx {
			return key, nil
		}
	}

	return "", errors.New("Key does not exist")
}

func getSortedKeys(vertices map[string]*Vertix) []string {
	keys := make([]string, 0)
	for _, v := range vertices {
		keys = append(keys, v.value)
	}
	sort.Strings(keys)

	return keys
}

func (g AdjacencyMatrixGraph) String() string {
	var sb strings.Builder

	keys := getSortedKeys(g.vertices)
	for j, k := range keys {
		var sb2 strings.Builder
		idx := g.vertices[k].key
		edges := g.matrix[idx]
		vSpace := ""
		if j > 0 {
			vSpace = " "
		}
		for i, v := range edges {
			// has an edge
			if v == 1 {
				space := ""
				if sb2.Len() > 0 {
					space = " "
				}

				str, err := g.getVertixByIdx(i)
				if err != nil {
					str = "Not Found"
				}
				sb2.WriteString(fmt.Sprintf("%s%s", space, str))
			}
		}

		sb.WriteString(fmt.Sprintf("%s(%s)->[%s]", vSpace, k, sb2.String()))
	}
	return sb.String()
}

func (g AdjacencyMatrixGraph) getEdges(node Vertix) []*Vertix {
	children := make([]*Vertix, 0)
	edges := g.matrix[node.key]
	for idx, edge := range edges {
		if edge == 1 {
			key, err := g.getVertixByIdx(idx)
			if err == nil {
				v := g.vertices[key]
				children = append(children, v)
			}
		}
	}

	return children
}

func (g *AdjacencyMatrixGraph) BfsTraversal(key string) (string, error) {
	var sb strings.Builder
	node, ok := g.vertices[key]

	if ok != true {
		return "", errors.New("Key not found")
	}

	stack := make([]*Vertix, 0)
	node.color = GREY
	stack = append(stack, node)

	for len(stack) > 0 {
		node = stack[0]
		stack = stack[1:]

		if node.color == GREY {
			node.color = BLACK
			edges := g.getEdges(*node)
			for _, edge := range edges {
				if edge.color == WHITE {
					edge.color = GREY
					stack = append(stack, edge)
				}
			}
			sb.WriteString(fmt.Sprintf("(%s)", node.value))
		}
	}
	return sb.String(), nil
}
