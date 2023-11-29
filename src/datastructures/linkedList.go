package datastructures

import "fmt"

type LinkedList struct {
}

var head *ListNode
var size int


func (list LinkedList) AddFirst(value int) {
	temp := ListNode{Data: value, Next: nil}
	if head == nil {
		head = &temp
	} else {
		temp.Next = head
		head = &temp
	}
	size++
	return
}

func (list LinkedList) AddLast(value int) {
	temp := ListNode{Data: value, Next: nil}
	if head == nil {
		head = &temp
	} else {
		cur := head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = &temp
	}
	size++
	return
}

func (list LinkedList) Print() {
	cur := head
	for cur != nil {
		fmt.Print(cur.Data)
		cur = cur.Next
	}
}
