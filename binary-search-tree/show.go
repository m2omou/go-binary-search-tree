package main

import "fmt"

func showInOrder(root *Node) {
	if (root != nil) {
		showInOrder(root.left)
		fmt.Println(root.value)
		showInOrder(root.right)
	}
}

func showPreOrder(root *Node) {
	if (root != nil) {
		fmt.Println(root.value)
		showInOrder(root.left)
		showInOrder(root.right)
	}
}

func showPostOrder(root *Node) {
	if (root != nil) {
		fmt.Println(root.value)
		showInOrder(root.left)
		showInOrder(root.right)
	}
}
