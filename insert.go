package main

func (root *Node) insert(new_node *Node) {
	if (new_node.value > root.value) {
		if (root.right == nil) {
			root.right = new_node
		} else {
			root.right.insert(new_node)
		}
	} else if (new_node.value < root.value) {
		if (root.left == nil) {
			root.left = new_node
		} else {
			root.left.insert(new_node)
		}
	}
}

func (tree *Bst) Insert(value int) {
	if tree.root == nil {
		tree.root = &Node{nil, nil, value}
	}
	tree.size++
	tree.root.insert(&Node{nil, nil, value})
}
