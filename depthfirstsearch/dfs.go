package depthfirstsearch

type Node struct {
	Value int
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []int) []int {
	array = append(array, n.Value)
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}
	return array
}

func (n *Node) BreadthFirstSearch(array []int) []int {
	queue := []*Node{n}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		array = append(array, current.Value)
		for _, child := range n.Children {
			queue = append(queue, child)
		}
	}
}