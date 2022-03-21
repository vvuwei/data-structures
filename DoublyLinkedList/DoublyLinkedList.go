package DoublyLinkedList

import "fmt"

type DoublyLinkedListNode struct {
	Value    int
	Previous *DoublyLinkedListNode
	Next     *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	Head *DoublyLinkedListNode
	Tail *DoublyLinkedListNode
}

func (l *DoublyLinkedList) Append(value int) {
	node := &DoublyLinkedListNode{Value: value}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Previous = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}
}

func (l *DoublyLinkedList) Prepend(value int) {
	node := &DoublyLinkedListNode{Value: value}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Previous = node
		l.Head = node
	}
}

func (l *DoublyLinkedList) Delete(value int) bool {
	node := l.Head

	if node == nil {
		return false
	}

	if node.Value == value {
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Head = node.Next
			node.Previous = nil
		}

		return true
	}

	node = node.Next

	for node != nil {
		if node.Value == value {
			node.Previous.Next = node.Next
			node.Next.Previous = node.Previous
			return true
		}

		node = node.Next
	}

	return false
}

func (l *DoublyLinkedList) Exist(value int) bool {
	node := l.Head

	for node != nil {
		if node.Value == value {
			return true
		}

		node = node.Next
	}

	return false
}

func (l *DoublyLinkedList) Length() int {
	length := 0
	node := l.Head

	for node != nil {
		length++
		node = node.Next
	}

	return length
}

func (l *DoublyLinkedList) String() string {
	if l.Head == nil {
		return "[]"
	}

	node := l.Head

	str := "[ "
	for node != nil {
		str += fmt.Sprintf("%v ", node.Value)
		node = node.Next
	}
	str += "]"

	return str
}