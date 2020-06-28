package dictionaryString

import (
	"container/heap"
	"testing"
)

func NewDictString(t *testing.T){
	sh := &StringHeap{"hello", "world", "Hello", "dict", "all", "my", "go"}
	heap.Init(sh)
	heap.Push(sh, "test")
	heap.Push(sh, "testing")
	heap.Push(sh, "W")
	heap.Push(sh, "f")
	for sh.Len() > 0{
		t.Logf(heap.Pop(sh).(string))
	}
}
