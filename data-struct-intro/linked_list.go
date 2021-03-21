package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
	prev *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// Append node as head
func (l *LinkedList) Append(n *Node) {
	l.length++

	sec := l.head
	l.head = n
	if sec == nil {
		l.tail = n
		return
	}
	l.head.prev = sec
	sec.next = l.head
}

// Prepend node as tail
func (l *LinkedList) Prepend(n *Node) {
	l.length++

	sec := l.tail
	l.tail = n
	if sec == nil {
		l.head = n
		return
	}
	l.tail.next = sec
	sec.prev = l.tail
}

// DeleteFirstByVal remove the fist node that matches to
// specified value
func (l *LinkedList) DeleteFirstByVal(val int) {
	n := l.tail
	if n == nil {
		return
	}

	for {
		if n.val == val {
			if n.next != nil {
				n.next.prev = n.prev
			}
			if n.prev != nil {
				n.prev.next = n.next
			}

			if l.tail.val == n.val {
				l.tail = n.next
			}
			if l.head.val == n.val {
				l.head = n.prev
			}
			return
		}

		n = n.next
		if n == nil {
			return
		}
	}
}

// Reverse the linked list
func (l *LinkedList) Reverse() {
	n := l.tail
	if n == nil {
		return
	}

	var prev *Node
	for n != nil {
		next := n.next
		n.next = prev
		prev = n
		n = next
	}
	l.head, l.tail = l.tail, l.head
}

func (l LinkedList) Print() {
	n := l.tail
	if n == nil {
		return
	}

	for {
		fmt.Print(n.val)
		n = n.next
		if n == nil {
			break
		}
		fmt.Print("=")
	}
	fmt.Print("\n")
}

func main() {
	sl := LinkedList{}

	n := Node{val: 0}
	n1 := Node{val: 1}
	n2 := Node{val: 2}
	n3 := Node{val: 3}
	n4 := Node{val: 2}
	n5 := Node{val: 6}

	sl.Append(&n)
	sl.Append(&n1)
	sl.Append(&n2)
	sl.Append(&n2)
	sl.Prepend(&n3)
	sl.Prepend(&n4)
	sl.Append(&n5)

	sl.Print()
	sl.Reverse()
	sl.Print()
}
