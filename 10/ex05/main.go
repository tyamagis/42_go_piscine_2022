package main

import (
	"fmt"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "3")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	root.Left.Right.Data = "8"
	fmt.Println(piscine.BTreeIsBinary(root))
}
