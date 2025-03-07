//Write a function that applies a given function f, in order, to each element in the tree.

package main

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}


func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}

	f(root.Data)

	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}

	

}

func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	BTreeApplyInorder(root, fmt.Println)
}


func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
	return root
}