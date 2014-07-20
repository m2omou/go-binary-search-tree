package main

func main() {

	test := NewBst()

	test.Insert(4)
	test.Insert(1)
	test.Insert(22)
	test.Insert(699)
	test.Insert(88)
	showInOrder(test.root)
	test.Delete(4)
	showInOrder(test.root)
}
