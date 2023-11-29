package datastructures

import (
	"fmt"
)

type DoublyLinkedList struct {
	head *DDLNode
	size int
}

func (list DoublyLinkedList) AddFirst(value int) {
	temp := DDLNode{Data: value, Next: nil, Prev: nil}
	if list.head == nil {
		list.head = &temp
	} else {
		temp.Next = list.head
		list.head.Prev = &temp
		list.head = &temp
	}
	size++
	return
}

func (list DoublyLinkedList) AddLast(value int) {
	temp := DDLNode{Data: value, Next: nil, Prev: nil}
	if list.head == nil {
		list.head = &temp
	} else {
		cur := list.head
		for cur != nil {
			cur = cur.Next
		}
		cur.Next = &temp
	}
	size++
	return
}

func (list DoublyLinkedList) Print() {
	cur := head
	for cur != nil {
		fmt.Print(cur.Data)
		cur = cur.Next
	}
}
