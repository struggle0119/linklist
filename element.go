package linklist

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	head   *Node
	length int
}
