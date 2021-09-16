package tree

import (
	"fmt"
	"strings"
)

type node struct {
	value  int
	parent *node
	left   *node
	right  *node
}

func (n *node) String() string {
	str := "nil"
	if n != nil {
		str = fmt.Sprintf("(%d)", n.value)
	}
	return str
}

type BinarySearchTree struct {
	root *node
	size int
}

func MakeBinarySearchTree(value int) BinarySearchTree {
	b := BinarySearchTree{
		root: &node{
			parent: nil,
			left:   nil,
			right:  nil,
			value:  value,
		},
		size: 1,
	}
	return b
}

func (b *BinarySearchTree) Add(value int) {
	if b == nil {
		b = &BinarySearchTree{
			root: &node{
				value:  value,
				parent: nil,
				right:  nil,
				left:   nil,
			},
			size: 1,
		}
		return
	}

	parent := b.root
	foundParent := false
	for foundParent != true {
		if parent.value < value && parent.right == nil {
			foundParent = true
		} else if parent.value > value && parent.left == nil {
			foundParent = true
		}

		if foundParent == false && parent.value < value {
			parent = parent.right
		} else if foundParent == false {
			parent = parent.left
		}
	}

	child := node{
		parent: parent,
		right:  nil,
		left:   nil,
		value:  value,
	}

	if parent.value > value {
		parent.left = &child
	} else {
		parent.right = &child
	}
	b.size += 1
}

func (b BinarySearchTree) GetSize() int {
	return b.size
}

func (b BinarySearchTree) BfsString() string {
	var sb strings.Builder
	buffer := make([]*node, 0)
	buffer = append(buffer, b.root)
	for len(buffer) > 0 {
		currentNode := buffer[0]
		if currentNode.left != nil {
			buffer = append(buffer, currentNode.left)
		}
		if currentNode.right != nil {
			buffer = append(buffer, currentNode.right)
		}
		sb.WriteString(fmt.Sprintf("%v", currentNode))
		buffer = buffer[1:]
	}
	return sb.String()
}

func dfsString(root *node, sb *strings.Builder) {
	if root == nil {
		return
	}

	dfsString(root.left, sb)
	sb.WriteString(fmt.Sprintf("%v", root))
	dfsString(root.right, sb)
}

func (b BinarySearchTree) String() string {
	var sb strings.Builder
	dfsString(b.root, &sb)
	return sb.String()
}

func bfsHelper(root *node, value int) *node {
	if root.value == value {
		return root
	}

	if root.value < value && root.right != nil {
		return bfsHelper(root.right, value)
	} else if root.value > value && root.left != nil {
		return bfsHelper(root.left, value)
	}

	return nil
}

func findNextInOrder(n *node) *node {
	if n.left == nil {
		return n
	}
	return findNextInOrder(n.left)
}

func (b BinarySearchTree) Bfs(value int) *node {
	return bfsHelper(b.root, value)
}

func (b *BinarySearchTree) Remove(value int) {
	n := b.Bfs(value)
	if b.size == 1 && n == nil {
		b.root = nil
		b.size = 0
		return
	}

	parent := n.parent
	if n.right != nil && n.left != nil {
		nextNode := findNextInOrder(n.right)
		if parent.value < nextNode.value {
			parent.right = nextNode
			nextNode.parent = parent
			nextNode.left = n.left
		} else {
			parent.left = nextNode
			nextNode.parent = parent
			nextNode.left = n.left
		}
	} else if n.right != nil && parent.value < n.right.value {
		parent.right = n.right
		n.right.parent = parent
	} else if n.right != nil && parent.value > n.right.value {
		parent.left = n.right
		n.right.parent = parent
	} else if n.left != nil && parent.value > n.left.value {
		parent.left = n.left
		n.left.parent = parent
	} else if n.left != nil && parent.value < n.left.value {
		parent.right = n.left
		n.left.parent = parent
	} else {
		if parent.right == n {
			parent.right = nil
		} else {
			parent.left = nil
		}
	}

	b.size -= 1
}
