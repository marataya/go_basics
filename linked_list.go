package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Push(val T) {
	if l.next == nil {
		l.next = &List[T]{val: val}
	} else {
		l.next.Push(val)
	}
}

func (l *List[t]) Print() {
	current := l
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func main() {
	list := &List[int]{val: 1}
	list.Push(2)
	list.Push(3)
	list.Print()

}
