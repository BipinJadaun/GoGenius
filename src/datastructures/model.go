package datastructures

type ListNode struct {
	Data int
	Next *ListNode
}

type DDLNode struct {
	Data int
	Next *DDLNode
	Prev *DDLNode
}

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}
