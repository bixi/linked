/*
Provides an interface and class for making linked lists.
*/
package llist

import (
	"errors"
	"github.com/KTAtkinson/linked"
)

type List struct {
	head linked.Noder
	tail linked.Noder
	length uint
}

// Get the head node.
func (l *List) Head() linked.Noder {
	return l.head
}

// Get the tail node.
func (l *List) Tail() linked.Noder {
	return l.tail
}

// Get length.
func (l *List) Length() uint {
	return l.length
}

// Append adds the given node as the tail of the list.
func (l *List) Append(n linked.Noder) (err error) {
	if l.head == nil {
		l.head = n
	}
	if l.tail != nil {
		n.SetPrev(l.tail)
		l.tail.SetNext(n)
	}

	l.tail = n
	l.length++
	return err
}

// Push adds the given node as the head of the list.
func (l *List) Push(n linked.Noder) (err error) {
	if l.tail == nil {
		l.tail = n
	}
	if l.head != nil {
		n.SetNext(l.head)
		l.head.SetPrev(n)
	}

	l.head = n
	l.length++
	return err
}

// Pop returns the node at the head of the list.
func (l *List) Pop() (n linked.Noder, err error) {
	if l.head == nil {
		return n, errors.New("Can't pop from an empty list.")
	}
	newHead, isLast, _ := l.head.Next()
	oldHead := l.head

	if !isLast {
		newHead.SetPrev(n)
	}
	l.head = newHead
	l.length--
	return oldHead, err
}

// Remove node.
func (l *List) Remove(n linked.Noder) error {
	prev, isFirst, _ := n.Prev()
	next, isLast, _ := n.Next()
	if (isFirst) {
		if (n != l.head) {
			return errors.New("Invalid node.")
		}
		l.head = next
	} else {
		prev.SetNext(next)
	}
	if (isLast) {
		if (n != l.tail) {
			return errors.New("Invalid node.")
		}
		l.tail = prev
	} else {
		next.SetPrev(prev)
	}
	n.SetNext(nil)
	n.SetPrev(nil)
	l.length--
	return nil
}

// Reverse reverses the current linked list in place.
func (l *List) Reverse() (err error) {
	next := l.head
	l.head = nil
	l.tail = nil

	for next != nil {
		newNext, _, _ := next.Next()

		next.SetNext(nil)
		next.SetPrev(nil)
		l.Push(next)

		next = newNext
	}

	return err
}

// Swap with another list.
func (l *List) Swap(r *List) {
	temp := l.head
	l.head = r.head
	r.head = temp

	temp = l.tail
	l.tail = r.tail
	r.tail = temp

	tempLen := l.length
	l.length = r.length
	r.length = tempLen
}
