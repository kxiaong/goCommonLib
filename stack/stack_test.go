package stack

import "testing"

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s.IsEmpty() == false{
		t.Error("Empty stack error")
	}

	s.Push(1)
	s.Push(2)
	val := s.Peak().(int)
	if val != 2{
		t.Error("Stack peak int error")
	}

	val = s.Pop().(int)
	if val != 2{
		t.Error("Stack pop error")
	}

	s.Push('a')
	rval := s.Peak().(rune)
	if rval != 'a'{
		t.Error("Stack peak rune error")
	}
	rval = s.Pop().(rune)
	if rval != 'a'{
		t.Error("Stack pop rune error")
	}

	val = s.Pop().(int)
	if val != 1{
		t.Error("Stack pop error")
	}

	if s.IsEmpty() == false{
		t.Error("Empty stack error")
	}
}