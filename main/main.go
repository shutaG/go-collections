package main

import (
	"fmt"
	"github.com/shutaG/go-collections/data_structure"
)

func main() {
	avl := data_structure.NewAvlTree()
	avl.Insert(1)
	avl.Insert(2)
	avl.Graph()
	avl.Insert(3)
	avl.Graph()
	avl.Insert(4)
	avl.Graph()
	avl.Insert(5)
	avl.Insert(0)
	avl.Graph()

	fmt.Printf("%+v\n", avl.PreOrder())
	fmt.Printf("%+v\n", avl.InOrder())
	fmt.Printf("%+v\n", avl.PostOrder())
	//node, depth := avl.GetRoot()
	//data_structure.PrintTree(node, depth, "_")

}
