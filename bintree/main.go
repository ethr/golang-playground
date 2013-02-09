package main

import "fmt"

// Defines methods that a binary tree should have
type Tree interface {
        Add(i int)
        Contains(i int) bool
}

// A binary tree
type BinaryTree struct {
        root *Node
}

// A node in a binary tree
type Node struct {
        parent *Node
        left *Node
        right *Node
        value int
}

func NewNode(parent *Node, value int) *Node {
        node := new(Node)
        node.parent = parent
        node.value = value
        return node
}

func (node *Node) Add(value int) {
        if value == node.value {
                return
        } else if value > node.value {
                if node.right == nil {
                        node.right = NewNode(node, value)
                } else {
                        node.right.Add(value)
                }
        } else {
                if node.left == nil {
                        node.left = NewNode(node, value)
                } else {
                        node.left.Add(value)
                }
        }
}

func (node *Node) GetValue() int {
        return node.value
}

func NewBinaryTree() *BinaryTree {
        tree := new(BinaryTree)
        return tree
}

func (tree *BinaryTree) Add(value int) {
        if tree.root == nil {
                tree.root = NewNode(nil, value)
        } else {
                tree.root.Add(value)
        }
}

func (tree *BinaryTree) Contains(value int) bool {
        var curr *Node = tree.root
        for {
                if curr == nil {
                        return false
                } else if curr.GetValue() == value {
                        return true
                } else {
                        if value > curr.GetValue() {
                                curr = curr.right
                        } else {
                                curr = curr.left
                        }
                }
        }
        return false
}

func main() {
        var tree Tree = NewBinaryTree()
        tree.Add(1)
        tree.Add(2)
        tree.Add(3)
        tree.Add(4)
        tree.Add(0)
        fmt.Println(tree.Contains(4))
        fmt.Println(tree.Contains(20))
        fmt.Println(tree.Contains(1))
        fmt.Println(tree.Contains(0))
}
