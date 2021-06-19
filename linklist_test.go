package linklist

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	node := New()
	node.Add(1)
	node.Add(2)
	node.Add(3)
	node.Push(0)
	node.Add(2)
	node.Push(0)
	node.RemoveValue(0)
	fmt.Println(node.Len())
	v := node.Head()
	for v != nil {
		fmt.Println(v.Data)
		v = v.Next
	}
}
