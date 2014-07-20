package main

func search(root *Node, value int) (bool) {
	if (root != nil) {
		if (value == root.value) {
			return true
		} else if (value > root.value) {
			return search(root.right, value)
		} else {
			return search(root.left, value)
		}
	}
	return false
}

/* Check if a node with given value is in the tree */
func (bst *Bst) Exists(v int) bool {
	bst.size--
	return search(bst.root, v)
}
