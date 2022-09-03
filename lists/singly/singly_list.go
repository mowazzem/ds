package singly

import "fmt"

type node[T comparable] struct {
	next  *node[T]
	value T
}

type List[T comparable] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// create and return a new List
func NewList[T comparable]() *List[T] {
	return &List[T]{}
}

// add to the first of the List
func (l *List[T]) Prepend(v T) {
	nwnode := &node[T]{value: v, next: l.head}
	l.head = nwnode
	l.length++
}

// add to the last of the List
func (l *List[T]) Append(v T) {
	nwnode := &node[T]{value: v}

	if l.head == nil {
		l.head = nwnode
		l.tail = l.head
	} else {
		l.tail.next = nwnode
		l.tail = l.tail.next
	}

	l.length++
}

// Deletes all the occurences
func (l *List[T]) Delete(v T) {
	curr := l.head
	var prev *node[T]
	for curr != nil {
		if prev == nil && curr.value == v {
			l.head = l.head.next
			curr = l.head
			l.length--
		} else if prev != nil && curr.value == v {
			prev.next = curr.next
			curr = curr.next
			l.length--
		} else {
			prev = curr
			curr = curr.next
		}
	}
}

// return index of matching value
// if not found return -1
func (l *List[T]) IndexOf(v T) int {
	idx := 0
	for it := l.head; it != nil; it = it.next {
		if it.value == v {
			return idx
		}
		idx++
	}
	return -1
}

// Delete node of given index
func (l *List[T]) DeleteOfIndex(idx int) {
	if l.length == 0 || idx < 0 {
		return
	}

	curr := l.head
	var prev *node[T]

	for i := 0; i != idx; i++ {
		prev = curr
		curr = curr.next
	}

	if prev == nil {
		l.head = l.head.next
		return
	}

	if curr == nil {
		prev.next = nil
		return
	}

	prev.next = curr.next
}

// Print all the values
func (l *List[T]) Print() {
	tmp := l.head
	for tmp != nil {
		fmt.Println(tmp.value)
		tmp = tmp.next
	}
	fmt.Println("end")
}

// Get all the List values
func (l *List[T]) GetAll() []T {
	vals := make([]T, 0, l.length)
	for it := l.head; it != nil; it = it.next {
		vals = append(vals, it.value)
	}
	return vals
}

// return totoal number of elements in the List
func (l List[T]) Len() int {
	return l.length
}
