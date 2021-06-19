package linklist

func New() *List {
	return &List{head: nil, length: 0}
}

func newEmptyNode(data interface{}) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

// chain is empty
func (l *List) isEmpty() bool {
	return l.head == nil
}

// chain length
func (l *List) Len() int {
	return l.length
}

// add element at first
func (l *List) Add(data interface{}) *Node {
	n := newEmptyNode(data)
	l.length++
	if l.isEmpty() {
		l.head = n
	} else {
		n.Next = l.head
		l.head = n
	}

	return l.head
}

// add element at tail
func (l *List) Push(data interface{}) {
	n := newEmptyNode(data)
	if l.isEmpty() {
		l.Add(data)
	} else {
		l.length++
		cur := l.head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = n
	}
}

// insert element at index fill data
func (l *List) Insert(index int, data interface{}) {
	if index <= 0 {
		l.Add(data)
	} else if index > l.Len() {
		l.Push(data)
	} else {
		n := newEmptyNode(data)
		l.length++
		pre := l.head
		var i int
		for i < (index - 1) {
			pre = pre.Next
			i++
		}
		n.Next = pre.Next
		pre.Next = n
	}
}

// pop up first element
func (l *List) Pop() (data interface{}) {
	data = l.head.Data
	l.head = l.head.Next
	if !l.isEmpty() {
		l.length--
	}
	return
}

// remove at index
func (l *List) Remove(index int) {
	if index <= 0 {
		l.head = l.head.Next
		if !l.isEmpty() {
			l.length--
		}
	} else if index >= l.Len() {
		// do nothing out of length
		return
	} else {
		var i int
		head := l.head
		for i != (index-1) && head.Next != nil {
			i++
			head = head.Next
		}
		head.Next = head.Next.Next
		l.length--
	}
}

// remove data all contain data
func (l *List) RemoveValue(data interface{}) {
	head := l.head
	if data == head.Data {
		l.head = l.head.Next
		if !l.isEmpty() {
			l.length--
		}
	}
	for head.Next != nil {
		if head.Next.Data == data {
			head.Next = head.Next.Next
			l.length--
		} else {
			head = head.Next
		}
	}
}

// chain contains element or not
func (l *List) Contains(data interface{}) bool {
	head := l.head
	for head != nil {
		if head.Data == data {
			return true
		}
		head = head.Next
	}

	return false
}

// get the chain head
func (l *List) Head() *Node {
	return l.head
}
