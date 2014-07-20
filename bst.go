package main

type Node struct {
	left *Node
	right *Node
	value int
	// Add whatever you want here
	//...
	//...
}

type Bst struct {
	root *Node;
	size int;
}

/* Get size of the tree */
func (bst *Bst) Size() int {
	return bst.size
}

/* Get size of the tree */
func (bst *Bst) Root() *Node {
	return bst.root
}

// Construtor
func NewBst() *Bst {
	bst := new(Bst)
	bst.size = 0;
	return bst
}
