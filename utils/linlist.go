package utils

import "fmt"

type Node struct {
	next  *Node
	value int
}

type List struct {
	head *Node
	tail *Node
	elm  int
	size int
}

func NewList(size int) List {
	return List{nil, nil, 0, size}
}

func (l *List) Push(a int) {
	el := Node{nil, a}
	if l.elm == 0 {
		l.head = &el
		l.tail = &el
		l.elm++
		return
	}

	l.tail.next = &el
	l.tail = &el
	l.elm++

	if l.elm > l.size {
		l.head = l.head.next
		l.elm--
	}
}

func (l *List) Pop() (int, error) {
	if l.elm == 0 {
		return 0, fmt.Errorf("empty list")
	}
	l.elm--
	a := l.head.value
	if l.elm == 0 {
		l.head = nil
	} else {
		l.head = l.head.next
	}
	return a, nil
}

func (l *List) Get(buf []int) {
	var max int
	if len(buf) > l.size {
		max = l.size
	} else {
		max = len(buf)
	}

	cur := l.head
	for i := 0; i < max; i++ {
		buf[i] = cur.value
		cur = cur.next
	}
}

func (l *List) Len() int {
	return l.elm
}
