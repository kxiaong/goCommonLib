package dictionaryString
// sort a string array by dictionary order by heap
// Usage:
// sh := &StringHeap{"some", "string", "here"}
// heap.Push(sh, "pushed string")
// heap.Push(sh, "push another string")
// for sh.Len()>0{
//     heap.Pop(sh).(string)   // pop string by dictionary order
// }

import (
	_ "container/heap"
)

type StringHeap []string

func (sh StringHeap) Len() int { return len(sh)}

func (sh StringHeap) Less(i, j int) bool{
	iBytes, jBytes := []byte(sh[i]), []byte(sh[j])
	iLen, jLen := len(iBytes), len(jBytes)

	if iLen == 0 && jLen == 0{
		// two null string
		return false
	}else if iLen == 0{
		// sh[i] is null string
		return true
	}else if jLen == 0{
		// sh[j] is null string
		return false
	}

	iptr, jptr := 0, 0
	for ;iptr<iLen && jptr<jLen;iptr,jptr=iptr+1,jptr+1{
		// To satisfy dictionary order, only 5 low-bits used
		if iBytes[iptr] & 0x1f == jBytes[jptr] & 0x1f {
			// if 5 low-bits equal, Upper-Case is less than lower-case
			return iBytes[iptr] < jBytes[iptr]
		} else{
			return iBytes[iptr]&0x1f < jBytes[jptr]&0x1f
		}
	}

	if iLen > jLen{
		// sh[i] > sh[j]
		return true
	}
	return false
}

func (sh StringHeap) Swap(i, j int){
	sh[i], sh[j] = sh[j], sh[i]
}

func (sh *StringHeap) Push(v interface{}){
	*sh = append(*sh, v.(string))
}

func (sh *StringHeap) Pop() interface{}{
	old := *sh
	n := len(old)
	x := old[n-1]
	*sh = old[0:n-1]
	return x
}
