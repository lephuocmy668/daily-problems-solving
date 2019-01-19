package linkedlist

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item the type of the linked list
type Item generic.Type

// Node is a single node that composes the linked list
type Node struct {
	Value Item
	Next  *Node
}

// LinkedList is the list of item
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
	Lock   sync.RWMutex
}

// NewLinkedList is constructor function that inits new linked list
func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (lk *LinkedList) Prepend(t Item) {
	lk.Lock.Lock()
	defer lk.Lock.Unlock()

	newNode := Node{
		Next:  lk.Head,
		Value: t,
	}

	if lk.Head != nil {
		newNode.Next = lk.Head
	}
	lk.Head = &newNode

	if lk.Tail == nil {
		lk.Tail = &newNode
	}

	lk.Length++
}

// Append ...
func (lk *LinkedList) Append(t Item) {
	lk.Lock.Lock()
	defer lk.Lock.Unlock()

	newNode := Node{
		Next:  nil,
		Value: t,
	}

	if lk.Tail != nil {
		lk.Tail.Next = &newNode
	}
	lk.Tail = &newNode

	if lk.Head == nil {
		lk.Head = &newNode
	}

	lk.Length++
}

// DeleteByValue ...
func (lk *LinkedList) DeleteByValue(value Item) {
	lk.Lock.Lock()
	defer lk.Lock.Unlock()

	if lk.Head == nil {
		return
	}

	node := &Node{
		Next:  lk.Head,
		Value: nil,
	}

	remove := lk.Head
	if lk.Head.Value == value {
		lk.Head = remove.Next
	}

	for {
		if node == nil || node.Next == nil {
			return
		}
		if node.Next.Value == value {
			remove = node.Next
			node.Next = remove.Next
			lk.Length--
		}
		node = node.Next
	}
}

// FindByValue ...
func (lk *LinkedList) FindByValue(value string) *Node {
	lk.Lock.Lock()
	defer lk.Lock.Unlock()

	node := lk.Head

	for {
		if node == nil {
			return nil
		}
		if node.Value == value {
			return node
		}
		node = node.Next
	}
}

// DeleteTail ...
func (lk *LinkedList) DeleteTail() {
	lk.Lock.Lock()
	defer lk.Lock.Unlock()

	if lk.Head == nil {
		return
	}

	node := lk.Head
	for {
		if node.Next.Next == nil {
			lk.Tail = node
			node.Next = nil
			lk.Length--
			return
		}

		node = node.Next
	}
}

// DeleteHead ...
func (lk *LinkedList) DeleteHead() {
	lk.Lock.Lock()
	defer lk.Lock.Unlock()

	if lk.Head == nil {
		return
	}

	if lk.Head.Next == nil {
		lk.Head = nil
		lk.Tail = nil
		lk.Length--
		return
	}

	if lk.Head.Next.Next == nil {
		lk.Head = lk.Head.Next
		lk.Tail = lk.Head
		lk.Length--
		return
	}

	lk.Head = lk.Head.Next
}

// FromSlice ...
func (lk *LinkedList) FromSlice(values []Item) {
	lk.Lock.Lock()
	defer lk.Lock.Unlock()

	for i := 0; i < len(values); i++ {
		lk.Append(values[i])
	}
}

// ToSlice ...
func (lk *LinkedList) ToSlice() []Item {
	lk.Lock.Lock()
	defer lk.Lock.Unlock()

	values := []Item{}

	node := lk.Head
	for {
		if node == nil {
			return values
		}

		values = append(values, node.Value)
		node = node.Next
	}
}

// Reverse ...
func (lk *LinkedList) Reverse() LinkedList {
	lk.Lock.Lock()
	defer lk.Lock.Unlock()

	return LinkedList{}
}
