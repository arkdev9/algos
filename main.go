package main

import btree "github.com/arkdev9/algos/models"

func main() {
	var tree = btree.CreateNode(10)
	btree.InsertNode(20, tree)
}
