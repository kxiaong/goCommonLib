package queue
/*
 循环队列的实现
 */

import "errors"

type Queue struct {
	queue []interface{}
	front int
	rear int
	Cap int
}

func NewQueue(cap int) *Queue{
	return &Queue{queue: make([]interface{}, cap), front:0, rear: 0, Cap: cap}
}

func (q *Queue) IsEmpty() bool{
	return q.rear == q.front
}

func (q *Queue) IsFull() bool{
	return (q.rear + 1)%q.Cap == q.front
}

func (q *Queue) Enqueue(value interface{}) error{
	if q.IsFull(){
		return errors.New("Queue is Full")
	}
	q.queue[q.rear] = value
	q.rear = (q.rear + 1) % q.Cap
	return nil
}

func (q *Queue) Dequeue() (val interface{}, err error){
	if q.IsEmpty(){
		return nil, errors.New("Queue is Empty")
	}
	tmp := q.queue[q.front]
	q.front	= (q.front + 1) % q.Cap
	return tmp, nil
}