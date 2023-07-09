package main

import btree "github.com/arkdev9/algos/structures"

func main() {
	tree := btree.CreateNode(10)
	btree.InsertNode(30, tree)
	btree.InsertNode(20, tree)
	btree.InsertNode(60, tree)
	btree.InsertNode(150, tree)
	btree.InsertNode(15, tree)
	btree.InsertNode(38, tree)

	btree.PrintTree(tree)
}
