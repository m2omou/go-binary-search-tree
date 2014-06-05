package main

import "fmt"

type Node struct {
	left *Node
	right *Node
	value int
}

type Bst struct {
	root *Node;
	length int;
}


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

func link(root *Node, parent *Node, value int) {
	if (parent.value > value) {
		if (root.right != nil) {
			root.left.left = root.left
			parent.left = root.right
		} else if (root.left != nil) {
			root.left.right = root.right
			parent.left = root.left
		} else {
			parent.left = nil
		}
	} else {
		if (root.right != nil) {
			root.right.left = root.left 
			parent.right = root.right
		} else if (root.left != nil) {
			root.right.right = root.right 
			parent.right = root.left
		} else {
			parent.right = nil
		}
	}
}

func cut(root *Node, parent *Node, value int) (*Node) {
	if (root != nil) {
		if (value == root.value) {
			if (root == parent) {
				root.left = root.right
				root = root.left
				return root
			}
			if (parent.value > value) {
				if (root.right != nil) {
					root.left.left = root.left
					parent.left = root.right
				} else if (root.left != nil) {
					root.left.right = root.right
					parent.left = root.left
				} else {
					parent.left = nil
				}
			} else {
				if (root.right != nil) {
					root.right.left = root.left 
					parent.right = root.right
				} else if (root.left != nil) {
					root.right.right = root.right 
					parent.right = root.left
				} else {
					parent.right = nil
				}
			}
			return root
		} else if (value > root.value) {
			return cut(root.right, root, value)
		} else {
			return cut(root.left, root,  value)
		}
	}
	return root
}



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

func (tree *Bst) Remove(value int) (bool) {
	if (!search(tree.root, value)) {
		return false
	}
	tree.root = cut(tree.root, tree.root, value)
	return true;
}

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
	tree.length++
	tree.root.insert(&Node{nil, nil, value})
}

func main() {

	test := new(Bst)
	
	test.Insert(4)
	test.Insert(1)
	test.Insert(22)
	test.Insert(699)
	test.Insert(88)
	
	showInOrder(test.root)
	
	test.Remove(4)
	fmt.Println("#")
	showInOrder(test.root)
	
	//root := &Node{nil, nil, 1}
	//root.Insert(&Node{nil, nil, 4})
	//root.insert(&Node{nil, nil, 3})
	//root.insert(&Node{nil, nil, 2})
	//root.insert(&Node{nil, nil, 0})
	//root.insert(&Node{nil, nil, 42})
	//fmt.Println(search(&root, 42))
	//root.remove(1)

	//showInOrder(root)
}



