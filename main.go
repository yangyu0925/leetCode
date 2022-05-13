package main

import "fmt"

func main() {
	//arr := []int{5,4,8,11,NULL,13,4,7,2,NULL,NULL,5,1}
	//
	//tree := Ints2TreeNode(arr)
	//
	//fmt.Println(pathSum(tree, 22))

	fmt.Println(NULL)

}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   ints[0],
	}

	queue := make([]*TreeNode, 1, n * 2)
	queue[0] = root

	i := 1

	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++

	}

	return root
}

func Tree2ints(tn *TreeNode) []int {
	res := make([]int, 0, 1024)

	queue := []*TreeNode{tn}

	for len(queue) > 0  {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, NULL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i - 1] == NULL {
		i--
	}
	return res[:i]

}

func Tree2Preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}
	res = append(res, Tree2Preorder(root.Left)...)
	res = append(res, Tree2Preorder(root.Right)...)
	return res
}

func Tree2Inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2Inorder(root.Left)
	res = append(res, root.Val)
	res = append(res, Tree2Inorder(root.Right)...)
	return res
}

func Tree2Postorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2Postorder(root.Left)
	res = append(res, Tree2Postorder(root.Right)...)
	res = append(res, root.Val)

	return res
}


func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	dfs(root,sum,[]int{},&ret)
	return ret
}

func dfs(root *TreeNode,sum int,arr []int,ret *[][]int){
	if root == nil{
		return
	}
	arr = append(arr,root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int,len(arr))
		copy(tmp,arr)
		*ret = append(*ret,tmp)
	}
	dfs(root.Left,sum - root.Val,arr,ret)
	dfs(root.Right,sum - root.Val,arr,ret)
	arr = arr[:len(arr)-1]
}