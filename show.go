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

// BFS Traversal
// Print the tree level-order
// Traverse root/parent, left child, right child
func showLevelOrder(root *Node) {
	queue := []*Node{root}
	for (len(queue) != 0) {
		current := queue[0]
		fmt.Println(current.data)
		if (current.left != nil) {
			queue = append(queue, current.left)
		}	
		if (current.right != nil) {
			queue = append(queue, current.right)
		}
		queue = append(queue[:0], queue[1:]...)
	}
}
