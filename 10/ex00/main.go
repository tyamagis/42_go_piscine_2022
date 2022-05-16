package main

import (
	"fmt"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	fmt.Println(root.Left.Data)
	fmt.Println(root.Data)
	fmt.Println(root.Right.Left.Data)
	fmt.Println(root.Right.Data)
	fmt.Println("=====")
	l := root.Left
	fmt.Println("Node", l.Data)
	// fmt.Println(l.Left.Data)
	// fmt.Println(l.Right.Data)
	fmt.Println("P", l.Parent.Data)
	fmt.Println("=====")
	r := root.Right
	fmt.Println("Node", r.Data)
	fmt.Println("L", r.Left.Data)
	// fmt.Println("R", r.Right.Data)
	fmt.Println("P", r.Parent.Data)
	fmt.Println("=====")
}
