package alphatree

import "fmt"

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

// AlphaTree represents the binary tree itself
type AlphaTree struct {
	Root *TreeNode
}

// Insert inserts a new string into the binary tree in alphabetical order
func (tree *AlphaTree) Insert(value string) {
	if tree.Root == nil {
		tree.Root = &TreeNode{Value: value}
	} else {
		tree.Root.Insert(value)
	}
}

// Insert inserts a new string into the binary tree
func (n *TreeNode) Insert(value string) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &TreeNode{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else if value > n.Value {
		if n.Right == nil {
			n.Right = &TreeNode{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
	// If value == n.Value, do nothing (duplicate check)
}

// InOrderTraversal traverses the tree in order and applies the provided function to each node's value
func (tree *AlphaTree) InOrderTraversal(f func(string)) {
	if tree.Root != nil {
		tree.Root.InOrderTraversal(f)
	}
}

// InOrderTraversal traverses the tree in order and applies the provided function to each node's value
func (n *TreeNode) InOrderTraversal(f func(string)) {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal(f)
	f(n.Value)
	n.Right.InOrderTraversal(f)
}

func (tree *AlphaTree) OutputTreeInAlphabeticalOrder() {
	tree.InOrderTraversal(func(value string) {
		fmt.Println(value)
	})
}

func NewAlphaTree() *AlphaTree {
	return &AlphaTree{}
}
