package BinaryTree

import (
	"github.com/kxiaong/goCommonLib/stack"
)

// Define Tree Node
type TreeNode struct {
	Val interface{}
	Left *TreeNode
	Right *TreeNode
}


func PreOrderTraversal(root *TreeNode){
	// 先序遍历二叉树
	// 递归方法
	if root == nil{
		return
	}

	// 访问根节点
	// fmt.Println(root.Val)

	//分别递归地遍历左子树和右子树
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func PreOrderIterTraversal(root *TreeNode) []interface{}{
	/*
	先序遍历二叉树, 非递归方法

	1. 访问一个结点p
	2. 将这个结点p压入stack
	3. 向左子树移动 p = p.left, go to 1
	4. 如果p为空，从stack弹出栈顶元素,访问栈顶元素的右子树p = stack.pop().Right, goto 1
	5. stack.IsEmpty() == true and p == nil  -->  finish
	*/
	result := make([]interface{}, 0)
	if root == nil{
		return result
	}
	nodeStack := stack.NewStack()

	for root != nil || !nodeStack.IsEmpty(){
		for root != nil{
			// 先序遍历，保存结果
			result = append(result, root.Val)
			nodeStack.Push(root)
			root = root.Left
		}
		root = nodeStack.Pop().(*TreeNode).Right
	}

	return result
}

func  InorderTraversal(root *TreeNode){
	// 中序遍历
	// 递归方法
	if root == nil{
		return
	}

	// 递归地遍历左子树
	InorderTraversal(root.Left)

	// 访问根节点
	//fmt.Println(root.Val.(int))

	// 递归地遍历右子树
	InorderTraversal(root.Right)
}

func InorderIterTraversal(root *TreeNode) []interface{}{
	/*
	中序遍历，非递归方法

	1. 将当前结点p压入stack
	2. 向该结点的左子树移动, p = p.Left, goto 1
	3. 当p为空，stack弹出栈顶元素，p=stack.Pop()。访问当前结点p.Val
	4. 向右子树移动, p = p.Right, goto 1
	*/
	result := make([]interface{}, 0)
	if root == nil{
		return result
	}

	nodeStack := stack.NewStack()
	for root != nil || !nodeStack.IsEmpty(){
		for root != nil{
			nodeStack.Push(root)
			root = root.Left
		}

		root = nodeStack.Pop().(*TreeNode)
		result = append(result, root.Val)
		root = root.Right
	}

	return result
}

func PostOrderTraversal(root *TreeNode){
	// 后序遍历，递归方法
	if root == nil{
		return
	}

	// 分别递归地遍历左子树和右子树
	PostOrderTraversal(root.Left)
	PostOrderTraversal(root.Right)

	// 访问根节点
	//fmt.Println(root.Val.(int))
}

func PostOrderIterTraversal(root *TreeNode) []interface{}{
	/* 后序遍历，非递归方法
	*/
	result := make([]interface{}, 0)
	if root == nil{
		return result
	}

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
			result = append(result, lastVisit.Val)
		}else{
			root = peak.Right
		}
	}

	return result
}
