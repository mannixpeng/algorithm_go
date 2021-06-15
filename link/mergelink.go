package link

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func (n *Node) String() {
	tmp := n
	for tmp.Next != nil {
		fmt.Print(tmp.Data, "->")
		tmp = tmp.Next
	}
	fmt.Println(tmp.Data)
}

func MergeLink(link1, link2 *Node) *Node{
	if link1 == nil {
		return link2
	}
	if link2 == nil {
		return link1
	}
	l1 := link1
	l2 := link2
	head := new(Node)
	prev := head

	for l1 != nil && l2 != nil {
		if l1.Data > l2.Data {
			tmp := l2.Next
			prev.Next = l2
			l2 = tmp
		} else {
			tmp := l1.Next
			prev.Next = l1
			l1 = tmp
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return head.Next
}
