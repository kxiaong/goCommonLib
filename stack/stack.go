package stack

import "container/list"

type Stack struct{
	list *list.List
}

func NewStack() *Stack{
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}){
	// push element to stack
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{}{
	// pop element from stack
	e := stack.list.Back()
	if e!= nil{
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{}{
	// 只检查栈顶元素，不弹出栈顶元素
	e := stack.list.Back()
	if e!= nil{
		return e.Value
	}
	return nil
}

func (stack *Stack) Len() int{
	// 当前栈大小（元素个数）
	return stack.list.Len()
}

func (stack *Stack) IsEmpty() bool{
	return stack.list.Len() == 0
}

func (stack *Stack) Init() *Stack{
	return NewStack()
}

