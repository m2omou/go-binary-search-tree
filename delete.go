package main

/* Find the left most node under current tree */
func minValue(n *Node) int {
	if n.left == nil {
		return n.value
	}
	return minValue(n.right)
}

func link(parent *Node, n *Node) {
	if parent.left == n {
		if n.left != nil {
			parent.left = n.left
		} else {
			parent.left = n.right
		}
	} else if parent.right == n {
		if n.left != nil {
			parent.right = n.left
		} else {
			parent.right = n.right
		}
	}
}


func del(n *Node, parent *Node, v int) bool {
	switch {
	case n.value == v:
		if n.left != nil && n.right != nil {
			n.value = minValue(n.right)
			return del(n.right, n, n.value)
		}
		link(parent, n); return true
	case n.value > v:
		if n.left == nil {
			return false
		}
		return del(n.left, n, v)
	case n.value < v:
		if n.right == nil {
			return false
		}
		return del(n.right, n, v)
	}
	return false
}

func (bst *Bst) Delete(v int) bool {
	if !bst.Exists(v) || bst.root == nil {
		return false
	}

	if bst.root.value == v {
		tempRoot := &Node{nil, nil, 0}
		tempRoot.left = bst.root
		r := del(bst.root, tempRoot, v)
		bst.root = tempRoot.left
		return r
	}
	return del(bst.root.left, bst.root, v) || del(bst.root.right, bst.root, v)
}
