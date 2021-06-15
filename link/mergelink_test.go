package link

import (
	"math/rand"
	"testing"
	"time"
)

func TestMergeLink(t *testing.T) {
	link1 := initLink(20)
	time.Sleep(2 * time.Second)
	link1.String()
	link2 := initLink(20)
	link2.String()
	link := MergeLink(link1, link2)
	link.String()
}

func initLink(num int) *Node {
	var head *Node
	tmp := head
	rand.Seed(time.Now().UnixNano())
	for i:= 0; i < num ; i++ {
		n := rand.Intn(2 * num)
		node := new(Node)
		node.Data = n
		if tmp == nil {
			head = node
			tmp = node
		} else {
			tmp.Next = node
			tmp = tmp.Next
		}
	}
	return head
}
