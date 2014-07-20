package main

import "fmt"

// Print the tree in-order
// Traverse the left sub-tree, root, right sub-tree
func showInOrder(root *Node) {
	if (root != nil) {
		showInOrder(root.left)
		fmt.Println(root.value)
		showInOrder(root.right)
	}
}

// Print the tree pre-order
// Traverse the root, left sub-tree, right sub-tree
func showPreOrder(root *Node) {
	if (root != nil) {
		fmt.Println(root.value)
		showInOrder(root.left)
		showInOrder(root.right)
	}
}

// Print the tree post-order
// Traverse left sub-tree, right sub-tree, root
func showPostOrder(root *Node) {
	if (root != nil) {
		fmt.Println(root.value)
		showInOrder(root.left)
		showInOrder(root.right)
	}
}
