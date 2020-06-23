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
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{}{
	e := stack.list.Back()
	if e!= nil{
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{}{
	e := stack.list.Back()
	if e!= nil{
		return e.Value
	}
	return nil
}

func (stack *Stack) Len() int{
	return stack.list.Len()
}

func (stack *Stack) IsEmpty() bool{
	return stack.list.Len() == 0
}

func (stack *Stack) Init() *Stack{
	return NewStack()
}

func New() *Stack{return new(Stack).Init()}
