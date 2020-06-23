package queue

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T){
	q := NewQueue(5)
	if q.IsEmpty() == false{
		t.Error("Queue IsEmpty error")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	enqueueErr := q.Enqueue(6)
	if enqueueErr == nil{
		t.Error("Queue Enqueue error")
	}

	if q.IsFull() == false{
		t.Error("Queue IsFull error")
	}
	for q.IsEmpty()==false{
		tmp, err := q.Dequeue()
		if err != nil{
			t.Errorf("Queue Dequeue error: %s", err)
		}else{
			t.Logf("Queue Dequeue : %d", tmp.(int))
		}
		fmt.Println(tmp)
	}
}
