package main

import (
	"Collection/stack"
	"fmt"
)

type TreeNode struct {
	Val interface{}
	Left *TreeNode
	Right *TreeNode
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}

	root.Left.Right.Left = &TreeNode{Val: 8}
	root.Left.Right.Right = &TreeNode{Val: 9}

	root.Right.Left.Right = &TreeNode{Val: 11}

	//PreOrderTraversal(root)
	//fmt.Println("********")
	//PreOrderIterTraversal(root)

	//InorderTraversal(root)
	//fmt.Println("********")
	//InorderIterTraversal(root)

	PostOrderTraversal(root)
	fmt.Println("********")
	PostOrderIterTraversal(root)
}


func PreOrderTraversal(root *TreeNode){
	// 先序遍历二叉树
	// 递归方法
	if root == nil{
		return
	}
	fmt.Println(root.Val)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func PreOrderIterTraversal(root *TreeNode){
	// 先序遍历二叉树
	// 非递归方法

	/*
		1. 访问一个结点p
		2. 将这个结点p压入stack
		3. 向左子树移动 p = p.left, go to 1
		4. 如果p为空，从stack弹出栈顶元素,访问栈顶元素的右子树p = stack.pop().Right, goto 1
		5. stack.IsEmpty() == true and p == nil  -->  finish
	*/

	if root == nil{
		return
	}
	result := make([]int, 0)
	nodeStack := stack.NewStack()

	for root != nil || !nodeStack.IsEmpty(){
		for root != nil{
			//先序遍历，保存结果
			//result.PushBack(root.Val)
			result = append(result, root.Val.(int))
			nodeStack.Push(root)
			root = root.Left
		}
		root = nodeStack.Pop().(*TreeNode).Right
	}

	for _, v := range result{
		fmt.Println(v)
	}
}

func  InorderTraversal(root *TreeNode){
	// 中序遍历
	// 递归方法
	if root == nil{
		return
	}
	InorderTraversal(root.Left)
	fmt.Println(root.Val.(int))
	InorderTraversal(root.Right)
}

func InorderIterTraversal(root *TreeNode){
	// 中序遍历，非递归方法
	/*
		1. 将当前结点p压入stack
		2. 向该结点的左子树移动, p = p.Left, goto 1
		3. 当p为空，stack弹出栈顶元素，p=stack.Pop()。访问当前结点p.Val
		4. 向右子树移动, p = p.Right, goto 1
	 */
	if root == nil{
		return
	}

	result := make([]int, 0)
	nodeStack := stack.NewStack()

	for root != nil || !nodeStack.IsEmpty(){
		for root != nil{
			nodeStack.Push(root)
			root = root.Left
		}

		root = nodeStack.Pop().(*TreeNode)
		result = append(result, root.Val.(int))
		root = root.Right
	}

	for _, v := range result{
		fmt.Println(v)
	}
}

func PostOrderTraversal(root *TreeNode){
	// 后序遍历，递归方法
	if root == nil{
		return
	}

	PostOrderTraversal(root.Left)
	PostOrderTraversal(root.Right)
	fmt.Println(root.Val.(int))
}

func PostOrderIterTraversal(root *TreeNode){
	// 后序遍历，非递归方法
	/*
		1.
	 */
	if root == nil{
		return
	}

	result := make([]int, 0)
	nodes := stack.NewStack()
	var lastVisit *TreeNode

	for root!=nil || !nodes.IsEmpty(){
		for root != nil{
			nodes.Push(root)
			root = root.Left
		}
		peak := nodes.Peak().(*TreeNode)
		if peak.Right == nil || peak.Right == lastVisit{
			lastVisit = nodes.Pop().(*TreeNode)
			result = append(result, lastVisit.Val.(int))
		}else{
			root = peak.Right
		}
	}

	fmt.Println(result)
}

